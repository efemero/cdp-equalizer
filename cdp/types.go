package cdp

import (
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
