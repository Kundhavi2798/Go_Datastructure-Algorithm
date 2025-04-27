package main

import (
	"fmt"
	"math/big"
)

func main() {
	n := []*big.Int{
		big.NewInt(2112),
		big.NewInt(1124),
		new(big.Int).SetString("363588395960667043875487", 10),
	}

	six := big.NewInt(6)
	zero := big.NewInt(0)

	for _, val := range n {
		remainder := new(big.Int).Mod(val, six)
		if remainder.Cmp(zero) == 0 {
			fmt.Println(true)
		} else {
			fmt.Println(false)
		}
	}
}
