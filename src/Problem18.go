package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	var fp *os.File
	var err error
	fileName := "/Users/A12715/IdeaProjects/go-eule/src/problem18"
	fp, err = os.Open(fileName)

	if err != nil {
		panic(err)
	}

	buf := [][]int{}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		strArray := strings.Split(scanner.Text(), " ")
		buf = append(buf, stringsToInts(strArray))
	}


	tmp := []int{}
	for bufIndex,ints := range buf {
		if bufIndex == 0 {
			tmp = ints
			continue;
		}

		new := []int{}
		for intsIndex,val := range ints {
			if intsIndex == 0 {
				new = append(new, tmp[0] + val)
			} else if intsIndex == len(ints)-1 {
				new = append(new, tmp[len(tmp) - 1] + val)
			} else {
				if (tmp[intsIndex-1] > tmp[intsIndex]) {
					new = append(new, tmp[intsIndex-1] + val)
				} else {
					new = append(new, tmp[intsIndex] + val)
				}
			}
		}
		fmt.Println(tmp)
		tmp = new
	}
	fmt.Println(tmp)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func stringsToInts(str []string) []int {
	ints := []int{}
	for _, val := range str {
		intVal,_ := strconv.Atoi(val)
		ints = append(ints, intVal)
	}
	return ints
}
