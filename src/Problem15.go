package main

import (
	"fmt"
)

func main () {

	// 初期化
	x := []int64{1}

	// パスカル
	for i:=int(1); i <= 20; i++ {
		x = pascal(x)
		fmt.Println(x)
	}

	// ネガティブパスカル
	y := negativePascal(x)
	fmt.Println(y)


}

func pascal(x []int64) []int64 {
	if len(x) == 1 {
		return []int64{1,1}
	}
	index := 0
	newx := []int64{}
	newx = append(newx, 1)
	for {
		if index + 1 == len(x) {
			break;
		} else {
			newx = append(newx, x[index] + x[index + 1])
			index++
		}
	}
	newx = append(newx, 1)
	return newx
}

func negativePascal(x []int64) int64 {
	result := x
	for {
		index := 0
		tmp := []int64{}
		for {
			if index + 1 == len(result) {
				break
			} else {
				tmp = append(tmp, result[index] + result[index + 1])
				index++
			}
		}

		fmt.Println(tmp)

		if len(tmp) == 1 {
			return tmp[0]
		} else {
			result = tmp
		}
	}
	return result[0]
}
