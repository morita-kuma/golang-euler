package main

import (
	"bufio"
	"fmt"
	"os"
	"math/big"
)
func main() {

	var fp *os.File
	var err error
	fileName := "/Users/A12715/IdeaProjects/go-eule/src/problem13"
	fp, err = os.Open(fileName)

	if err != nil {
		panic(err)
	}
	x := big.NewInt(1)
	y := big.NewInt(2)
	fmt.Println(x.Add(x, y))

	sum := big.NewInt(0)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		x := big.NewInt(0)
		x.SetString(scanner.Text(),0)
		sum.Add(sum, x)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(sum)
}
