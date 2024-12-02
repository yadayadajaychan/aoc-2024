// AOC Day 1 Part 1
package main

import (
	"fmt"
	"io"
	"strconv"
	"slices"
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

	slices.Sort(left)
	slices.Sort(right)

	var distance int
	for i, _ := range left {
		dist := left[i] - right[i]
		if dist < 0 {
			dist = -1 * dist
		}
		distance += dist
	}

	fmt.Println(distance)
}
