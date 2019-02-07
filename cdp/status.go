package cdp

import (
	"fmt"
	"math/big"
)

// Status represents the status of a CDP at a given price
type Status struct {
	ID      int64
	DaiDebt *Float
	EthCol  *Float
	Price   *Float
	Ratio   *Float
	DaiNet  *Float
	EthNet  *Float
}

// Float is for json purposes
type Float big.Float

// MarshalJSON transforms the Float in JSON number
func (f *Float) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("%.2f", (*big.Float)(f))
	return []byte(s), nil
}

func (f *Float) String() string {
	return fmt.Sprintf("%.2f", (*big.Float)(f))
}

func (s *Status) String() string {
	return fmt.Sprintf("%.2f: %.2f (%.2f)", (*big.Float)(s.Price), (*big.Float)(s.EthNet), (*big.Float)(s.DaiNet))
}
