package main

import (
	"util"
)

// 2000000までの素数の総和を求める
// エラトステネスのふるいを使用
func main () {
	util.Measure(resolver())
}

func resolver () util.Resolver {
	return func () {
		x:=int64(2000000)
		util.SieveOfEratosthenes(x)
	}
}
