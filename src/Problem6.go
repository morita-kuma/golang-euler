package main



import (
	"fmt"
	"math"
)

func main () {
	x := uint32(100);
	result := squareOfSum(x) - sumOfSquare(x)
	fmt.Println(result)
}

// 1^2 + 2^2 + 3^2 + 4^2 ...
func sumOfSquare (x uint32) uint32 {
	sum := uint32(0)
	for i := uint32(1); i <= x; i++ {
		sum += uint32(math.Pow(float64(i), 2))
	}
	return sum
}

// (1 + 2 + 3 + 4 ...)^2
func squareOfSum (x uint32) uint32 {
	sum := uint32(0)
	for i :=uint32(1); i<=x;i++ {
		sum += i
	}
	return (uint32)(math.Pow(float64(sum), 2))
}
