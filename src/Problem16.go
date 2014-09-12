package main


import (
	"fmt"
	"math/big"
	"strconv"
)

func main () {
	x := big.NewInt(0)
	rst := x.Exp(big.NewInt(2), big.NewInt(1000), nil).String()

	sum := 0;
	for i := 0; i < len(rst); i ++ {
		digit,_ := strconv.Atoi(rst[i:i+1])
		sum += digit
	}
	fmt.Println(sum)
}
