package main

import (
	"fmt"
	"math"
)

func main() {
	sum := float64(1)
	m := make(map[int64]int64)


	// 最大時数を抽出
	for i := int64(2); i <= 20; i++ {
		for key,value := range solve(i) {

			if m[key] == 0 {
				m[key] = value
			} else {
				if m[key] < value {
					m[key] = value
				}
			}
		}
	}

	for key, value := range m {
		fmt.Printf("sum = %f, key = %f, value = %f\n",sum, key, value)
		sum = sum*math.Pow((float64)(key), (float64)(value))
	}

	fmt.Println(sum)

}

// 素因数分解
func solve(x int64) map[int64]int64{

	// init
	nokori := x
	prime := (int64)(2)
	m := make(map[int64]int64)

	// resolve
	for {
		if (nokori % prime == 0 ) {
			m[prime] += 1
			nokori = nokori / prime
		} else {
			prime = getNextPrime(prime)
		}

		if isPrime(nokori) {
			m[nokori] += 1
			break
		}
	}

	return m
}

// 指定した数値の次のprimeを求める
func getNextPrime(prime int64) int64 {
	for x := (prime+1); ; x++ {
		if isPrime(x) == true {
			return x;
		}
	}
}

func isPrime(x int64) bool {
	count := 0
	for j := int64(1); j < x; j++ {
		if x%j == 0 {
			count++
		}

		if count > 1 {
			return false;
		}
	}
	return true;
}
