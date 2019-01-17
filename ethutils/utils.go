package ethutils

import (
	"errors"
	"math/big"
)

// FromWei transforms wei to a decimal value with `decimals` orders of magnitude (divide by 10^`decimals`)
func FromWei(wei *big.Int, decimals int) *big.Float {
	floatValue := new(big.Float).SetInt(wei)
	mul := big.NewInt(10)
	mul = mul.Exp(mul, big.NewInt(int64(decimals)), nil)
	floatMul := new(big.Float).SetInt(mul)
	result := new(big.Float).Quo(floatValue, floatMul)
	return result
}

// ToWei transforms decimal value to wei with `decimals` orders of magnitude (multiply by 10^`decimals`)
func ToWei(value *big.Float, decimals int) *big.Int {
	floatValue := new(big.Float).Copy(value)
	mul := big.NewInt(10)
	mul = mul.Exp(mul, big.NewInt(int64(decimals)), nil)
	floatMul := new(big.Float).SetInt(mul)
	floatValue = floatValue.Mul(floatValue, floatMul)
	wei, _ := floatValue.Int(new(big.Int))
	return wei
}

// SliceToBytes32 transforms a []byte to a [32]byte
func SliceToBytes32(bs []byte) (ba [32]byte, err error) {
	n := len(bs)
	if n > 32 {
		return ba, errors.New("byte slice too long (max 32)")
	}
	for i, b := range bs {
		ba[32+i-n] = b
	}
	return ba, nil
}
