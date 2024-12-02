// AOC Day 1 Part 2
package main

import (
	"fmt"
	"io"
	"strconv"
)

func main() {
	var left, right []int
	var a, b string
	var l, r int
	for {
		_, err := fmt.Scanln(&a, &b)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		l, err = strconv.Atoi(a)
		if err != nil {
			panic(err)
		}

		r, err = strconv.Atoi(b)
		if err != nil {
			panic(err)
		}

		left = append(left, l)
		right = append(right, r)
	}

	var similarity int
	for _, v := range left {
		var count int
		for _, v2 := range right {
			if v2 == v {
				count += 1
			}
		}

		similarity += v * count
	}

	fmt.Println(similarity)
}
