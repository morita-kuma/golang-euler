package main

import (
	"fmt"
	"math/big"
	"strconv"
)


func factorial(n int64) *big.Int {
	if n < 0 {
		return big.NewInt(1)
	}
	if (n==0) {
		return big.NewInt(1)
	}
	bigN := big.NewInt(n)
	return bigN.Mul(bigN, factorial(n-1))
}

func main() {
	sum := 0
	result := factorial(100).String()
	for i := 0; i < len(result) - 1; i++ {
		x,_ := strconv.Atoi(result[i:i+1])
		sum = sum + x
	}
	fmt.Println(sum)
}
