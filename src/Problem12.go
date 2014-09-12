package main

import (
	"fmt"
	"util"
)

// x = n * (x+1) / 2 において xの約数が最初に500を超えるものを探す
// xを素因数分解(x = a^p+ b^q + c^r)した時の約数の個数は (p+1)(q+1)(r+1)で表せることを利用
func main() {
	util.Measure(resolver())
}

func resolver () util.Resolver {
	fc := func () {
		for i:=1;;i++ {
			num := triagleNumbers(i)
			divisors := int64(1)
			for _, value := range util.PrimeDecomposition(int64(num)) {
				divisors = divisors * (value + 1)
			}
			if divisors > 500 {
				fmt.Printf("i = %d, num = %d, divisors = %d\n", i,num, divisors)
				break;
			}
		}
	}
	return fc
}

// 三角数を取得
func triagleNumbers (x int) int {
	return x * (x+1) / 2
}
