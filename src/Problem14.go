package main

import (
	"fmt"
)

// n → n/2 (n is even)
// n → 3n + 1 (n is odd)
// n * 2 ただし (n-1)/3の式で解が正の奇数の場合は分岐する)
func main() {
	/*	13 40 20 10 5 16 8 4 2 1

		1 2 4 8 16 32 64 128 256 512 1024
						(64)

				   (16)
					5  10 20  40  80  160  320
										   (160)
											53
								(40)
								 13  26   52	                             */

	current := int64(8)
	node := []int64{}
	for {
		current = current * 2
		if (current-1)%3 == 0 && ((current-1)/3) % 3 != 0 {
			current = (current-1)/3
			node = append(node, current)
			if current > 5000000 {
				fmt.Println(current)
				break;
			}
		}
	}


	for _, n := range node {
		fmt.Println(n)
	}

/*	x := 4367557
	for {

		if x % 2 == 0 {
			x = x / 2
		} else {
			x = 3 * x + 1
		}

		fmt.Println(x)
		if x < 1000000 {
			break
		}
	}*/
}
