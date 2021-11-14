package service

import (
	"fmt"
	"math/big"
)

func BigAdd() {
	newInt := big.NewInt(2)
	in := newInt
	i := big.NewInt(4)
	i.Add(newInt, in).Add(i, in)
	fmt.Printf(i.String())
}
