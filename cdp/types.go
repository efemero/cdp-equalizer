package cdp

import (
	"errors"
	"log"
	"math/big"
)

// CDP represents a CDP
type CDP struct {
	ID      int64
	BytesID [32]byte
	DaiDebt *big.Float
	PethCol *big.Float
	EthCol  *big.Float
}

// GetRatio returns the collateralization ratio of the CDP at the actual price and Peth / Eth ratio
func (cdp *CDP) GetRatio(ethPrice *big.Float, pethRatio *big.Float) *big.Float {

	ecol := new(big.Float).Mul(cdp.PethCol, pethRatio)
	dcol := new(big.Float).Mul(ecol, ethPrice)
	if cdp.DaiDebt.Cmp(big.NewFloat(0.0)) == 0 {
		return nil
	}
	ratio := new(big.Float).Quo(dcol, cdp.DaiDebt)

	return ratio
}

// GetMaxEthToFree returns the maximum number of eth to free
func (cdp *CDP) GetMaxEthToFree(ethPrice *big.Float) *big.Float {
	if cdp.PethCol.Cmp(big.NewFloat(0.0)) == 0 {
		return nil // nothing available to free!
	}
	edebt := new(big.Float).Quo(cdp.DaiDebt, ethPrice)
	finalEcol := new(big.Float).Mul(edebt, big.NewFloat(1.5))
	ethToFree := new(big.Float).Sub(cdp.EthCol, finalEcol)

	return ethToFree
}

// GetEthToFree returns the number of eth to free to go to the target ratio
func (cdp *CDP) GetEthToFree(ethPrice, target *big.Float) *big.Float {
	if cdp.EthCol.Cmp(big.NewFloat(0.0)) == 0 {
		return nil // nothing available to free!
	}
	edebt := new(big.Float).Quo(cdp.DaiDebt, ethPrice)
	net := new(big.Float).Sub(cdp.EthCol, edebt)
	target2 := new(big.Float).Sub(target, big.NewFloat(1.0))
	finalDebt := new(big.Float).Quo(net, target2)
	finalEcol := new(big.Float).Mul(finalDebt, target)
	ethToFree := new(big.Float).Sub(cdp.EthCol, finalEcol)
	max := cdp.GetMaxEthToFree(ethPrice)

	limit := new(big.Float).Mul(max, big.NewFloat(0.95))
	if ethToFree.Cmp(limit) > 0 {
		ethToFree = limit
	}

	return ethToFree
}

// GetMaxPethToFree returns the maximum number of peth to free
func (cdp *CDP) GetMaxPethToFree(ethPrice *big.Float, pethRatio *big.Float) *big.Float {
	if cdp.PethCol.Cmp(big.NewFloat(0.0)) == 0 {
		return nil // nothing available to free!
	}
	edebt := new(big.Float).Quo(cdp.DaiDebt, ethPrice)
	pdebt := new(big.Float).Quo(edebt, pethRatio)
	finalPcol := new(big.Float).Mul(pdebt, big.NewFloat(1.5))
	pethToFree := new(big.Float).Sub(cdp.PethCol, finalPcol)

	return pethToFree
}

// GetPethToFree returns the number of peth to free to go to the target ratio
func (cdp *CDP) GetPethToFree(ethPrice, pethRatio, target *big.Float) *big.Float {
	if cdp.PethCol.Cmp(big.NewFloat(0.0)) == 0 {
		return nil // nothing available to free!
	}
	edebt := new(big.Float).Quo(cdp.DaiDebt, ethPrice)
	pdebt := new(big.Float).Quo(edebt, pethRatio)
	net := new(big.Float).Sub(cdp.PethCol, pdebt)
	target2 := new(big.Float).Sub(target, big.NewFloat(1.0))
	finalDebt := new(big.Float).Quo(net, target2)
	finalPcol := new(big.Float).Mul(finalDebt, target)
	pethToFree := new(big.Float).Sub(cdp.PethCol, finalPcol)
	max := cdp.GetMaxPethToFree(ethPrice, pethRatio)

	limit := new(big.Float).Mul(max, big.NewFloat(0.95))
	if pethToFree.Cmp(limit) > 0 {
		pethToFree = limit
	}

	return pethToFree
}

// GetMaxDaiToDraw returns the maximum number of DAI to draw
func (cdp *CDP) GetMaxDaiToDraw(ethPrice *big.Float, pethRatio *big.Float) *big.Float {
	ecol := new(big.Float).Mul(cdp.PethCol, pethRatio)
	dcol := new(big.Float).Mul(ecol, ethPrice)
	finalDaiDebt := new(big.Float).Quo(dcol, big.NewFloat(1.5))
	daiToDraw := new(big.Float).Sub(finalDaiDebt, cdp.DaiDebt)

	return daiToDraw
}

// GetDaiToDraw returns the number of DAI to draw to go to the target ratio
func (cdp *CDP) GetDaiToDraw(ethPrice, pethRatio, target *big.Float) *big.Float {
	ecol := new(big.Float).Mul(cdp.PethCol, pethRatio)
	dcol := new(big.Float).Mul(ecol, ethPrice)
	net := new(big.Float).Sub(dcol, cdp.DaiDebt)
	target2 := new(big.Float).Sub(target, big.NewFloat(1.0))
	finalDebt := new(big.Float).Quo(net, target2)
	daiToDraw := new(big.Float).Sub(finalDebt, cdp.DaiDebt)

	max := cdp.GetMaxDaiToDraw(ethPrice, pethRatio)

	limit := new(big.Float).Mul(max, big.NewFloat(0.95))
	if daiToDraw.Cmp(limit) > 0 {
		daiToDraw = limit
	}

	return daiToDraw
}

// GetChangePrices returns the prices (up and down) where this CDP must be equalized
func (cdp *CDP) GetChangePrices(ethPrice, minRatio, maxRatio, pethRatio *big.Float) (minPrice, maxPrice *big.Float) {
	currentRatio := cdp.GetRatio(ethPrice, pethRatio)
	minPrice = new(big.Float).Mul(ethPrice, new(big.Float).Quo(minRatio, currentRatio))
	maxPrice = new(big.Float).Mul(ethPrice, new(big.Float).Quo(maxRatio, currentRatio))
	return minPrice, maxPrice
}

// EqualizeCDP returns a new CDP equalized at targetRatio for a given price
func (cdp *CDP) EqualizeCDP(ethPrice, targetRatio, pethRatio *big.Float) (newCDP *CDP, err error) {
	newCDP = new(CDP)
	newCDP.ID = cdp.ID
	newCDP.BytesID = cdp.BytesID
	ethToFree := cdp.GetEthToFree(ethPrice, targetRatio)
	if ethToFree.Cmp(big.NewFloat(0.0)) > 0 {
		daiToWipe := new(big.Float).Mul(ethPrice, ethToFree)
		pethToFree := new(big.Float).Quo(ethToFree, pethRatio)
		newCDP.EthCol = new(big.Float).Sub(cdp.EthCol, ethToFree)
		newCDP.PethCol = new(big.Float).Sub(cdp.PethCol, pethToFree)
		newCDP.DaiDebt = new(big.Float).Sub(cdp.DaiDebt, daiToWipe)
	} else {
		daiToDraw := cdp.GetDaiToDraw(ethPrice, pethRatio, targetRatio)
		if daiToDraw.Cmp(big.NewFloat(0.0)) < 0 {
			return nil, errors.New("cannot equalize the CDP")
		}
		ethToLock := new(big.Float).Quo(daiToDraw, ethPrice)
		pethToLock := new(big.Float).Quo(ethToLock, pethRatio)
		newCDP.DaiDebt = new(big.Float).Add(cdp.DaiDebt, daiToDraw)
		newCDP.EthCol = new(big.Float).Add(cdp.EthCol, ethToLock)
		newCDP.PethCol = new(big.Float).Add(cdp.PethCol, pethToLock)
	}
	return newCDP, nil
}

// GetStatus returns the status of the CDP for this price
func (cdp *CDP) GetStatus(ethPrice, pethRatio, target *big.Float) (status *Status, err error) {
	status = new(Status)
	status.ID = cdp.ID
	bigZero := *big.NewFloat(0.0)

	daiDebt := Float(bigZero)
	if cdp.DaiDebt != nil {
		daiDebt = Float(*cdp.DaiDebt)
	}
	status.DaiDebt = &daiDebt

	ethCol := Float(bigZero)
	if cdp.EthCol != nil {
		ethCol = Float(*cdp.EthCol)
	}
	status.EthCol = &ethCol

	price := Float(bigZero)
	if ethPrice != nil {
		price = Float(*ethPrice)
	}
	status.Price = &price

	ratio := Float(bigZero)
	if cdp.GetRatio(ethPrice, pethRatio) != nil {
		ratio = Float(*(cdp.GetRatio(ethPrice, pethRatio)))
	}
	status.Ratio = &ratio

	dcol := new(big.Float).Mul(cdp.EthCol, ethPrice)
	net := new(big.Float).Sub(dcol, cdp.DaiDebt)
	daiNet := Float(bigZero)
	if net != nil {
		daiNet = Float(*net)
	}
	status.DaiNet = &daiNet

	enet := new(big.Float).Quo(net, ethPrice)
	ethNet := Float(bigZero)
	if enet != nil {
		ethNet = Float(*enet)
	}
	status.EthNet = &ethNet

	return status, nil
}

// Log print infos on the CDP in the logs
func (cdp *CDP) Log(ethPrice, pethRatio, target *big.Float) {
	ecol := new(big.Float).Mul(cdp.PethCol, pethRatio)
	dcol := new(big.Float).Mul(ecol, ethPrice)
	net := new(big.Float).Sub(dcol, cdp.DaiDebt)

	log.Printf("CDP #%d status:\n", cdp.ID)
	log.Printf("\tDebt       : %.2f DAI", cdp.DaiDebt)
	log.Printf("\tCol        : %.5f PETH", cdp.PethCol)
	log.Printf("\t           : %.5f ETH", ecol)
	log.Printf("\tETH price  : %.5f DAI", ethPrice)
	log.Printf("\tPETH / ETH : %.5f ", pethRatio)
	log.Printf("\tRatio      : %.2f %%", new(big.Float).Mul(cdp.GetRatio(ethPrice, pethRatio), big.NewFloat(100.0)))
	log.Printf("\tNet Value  : %.2f DAI", net)
	log.Printf("\t           : %.5f ETH", new(big.Float).Quo(net, ethPrice))
	log.Printf("\tDraw       : %.2f DAI", cdp.GetDaiToDraw(ethPrice, pethRatio, target))
	log.Printf("\tFree       : %.5f PETH", cdp.GetPethToFree(ethPrice, pethRatio, target))
	log.Printf("\t           : %.5f ETH", new(big.Float).Mul(cdp.GetPethToFree(ethPrice, pethRatio, target), pethRatio))
	log.Println()
}
