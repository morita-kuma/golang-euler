package main

import (
	"math"
	"fmt"
)

// a + b + c = 1000 なピタゴラス数を探す
func main () {

	m := 2
	n := 1

	for {

		if m > n && (m - n) % 2 == 1 {
			a,b,c := pythagorean(m , n)
			if a + b + c == 1000 {
				fmt.Printf("a = %d, b = %d, c = %d, result = %d", a,b,c, (a*b*c))
				break;
			}
		}

		if m - n == 1 {
			m ++
			n = 1
		} else {
			n ++
		}

	}
}

// Pythagorean Triplet (a^2 + b^2 = c^2)
// Theorem (a, b, c) = (m^2 - n^2, 2mn, m^2 + n^2)
//         * m - n = odd number
func pythagorean(m int, n int) (a int, b int, c int) {
	x := math.Pow((float64)(m), 2) - math.Pow((float64)(n), 2)
	y := 2 * m * n
	z := math.Pow((float64)(m), 2) + math.Pow((float64)(n), 2)
	return int(x), int(y), int(z)
}
