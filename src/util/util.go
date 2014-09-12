package util

import (
	"time"
	"fmt"
	"math"
)

// 素因数分解
func PrimeDecomposition(x int64) map[int64]int64{

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
			prime = GetNextPrime(prime)
		}
		if IsPrime(nokori) {
			m[nokori] += 1
			break
		}
	}
	return m
}

// 指定した数値の次のprimeを求める
func GetNextPrime(prime int64) int64 {
	for x := (prime+1); ; x++ {
		if IsPrime(x) == true {
			return x;
		}
	}
}

// 素数であればtrue
func IsPrime(x int64) bool {
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

// エラトステネスのふるい (降順で判定)
func SieveOfEratosthenes(x int64) {

	sqrt := math.Sqrt((float64)(x));
	primes := getPrimes((int)(sqrt))
	sum := int64(0)
	for i := int64(2); i < x; i++ {
		isPrime := true

		// 偶数の場合は素数では無い
		if i > 2 && i % 2 == 0 {
			continue;
		}

		// 素数倍の場合は素数では無い
		for _, prime := range primes {
			if i != prime && i%prime == 0 {
				isPrime = false
				break;
			}
		}

		if isPrime == true {
			sum += i
		}
	}
	fmt.Println(sum)
}

// 特定の数値までの素数群を取得
func getPrimes(x int) []int64 {
	primes := []int64{}
	for i := 2; i <= x; i++ {

		// 偶数は素数ではない
		if i > 2 && i%2 == 0 {
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
		}
	}
	return primes
}

// 関数リテラルの方宣言
type Resolver func()

// 計測
func Measure (resolver Resolver) {
	start := time.Now()
	resolver()
	end := time.Now()
	fmt.Println(end.Sub(start).Seconds())
}
