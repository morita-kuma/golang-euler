package main

import (
	"fmt"
	"math"
	"util"
)

func main() {
	sum := float64(1)
	m := make(map[int64]int64)


	// 最大時数を抽出
	for i := int64(2); i <= 20; i++ {
		for key,value := range util.PrimeDecomposition(i) {

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

	fmt.Printf("%f", sum)


}
