package main

import (
	"fmt"
	"math"
)

func main() {
	//resolve(13195)
	solve(600851475143)
}

// 素因数分解
func solve(x int64) {

	// init
	nokori := x
	prime := (int64)(2)
	m := make(map[int64]int)

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

	// view
	for key,value := range m {
		fmt.Printf("key =%f, value=%f\n", key, value)
	}
}

// 指定した数値の次のprimeを求める
func getNextPrime(prime int64) int64 {
	for x := (prime+1); ; x++ {
		if isPrime(x) == true {
			return x;
		}
	}
}

// エラトテネスのふるい (降順で判定)
func resolve(x int64) {
	sqrt := math.Sqrt((float64)(x));
	fmt.Println(sqrt)

	primes := getPrimes((int)(sqrt))
	fmt.Println(primes)

	for i := x; i > 2; i-- {

		isPrime := true
		for _, prime := range primes {
			if i%prime == 0 {
				isPrime = false
			}
		}

		if isPrime == true {
			fmt.Println(i)
			break;
		}
	}
}

// 特定の数値までの素数群を取得
func getPrimes(x int) []int64 {
	primes := []int64{2, 3, 5, 7, 11, 13}
	for i := 15; i < x; i++ {

		// 偶数は素数ではない
		if i%2 == 0 {
			continue;
		}

		count := 0
		for j := 1; j < i; j++ {
			if i%j == 0 {
				count++
			}

			if count > 1 {
				continue;
			}
		}
		if count < 2 {
			primes = append(primes, (int64)(i))
			fmt.Println(i)
		}
	}
	return primes
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
