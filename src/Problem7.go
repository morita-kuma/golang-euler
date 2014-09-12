package main

import (
	"fmt"
)


// 10001st prime
func main () {
	count := int(1)
	currentPrime := int64(2)

	for {
		currentPrime = getNextPrime(currentPrime)
		count++

		if count == 10001 {
			fmt.Println(currentPrime)
			break;
		}
	}
}

// next prime
func getNextPrime(prime int64) int64 {
	for x := (prime+1); ; x++ {
		if isPrime(x) == true {
			return x;
		}
	}
}

// prime or not
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
