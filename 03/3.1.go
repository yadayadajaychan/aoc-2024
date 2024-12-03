// AOC Day 3 Part 1
package main

import (
	"fmt"
	"regexp"
	"os"
	"strconv"
)

func main () {
	input, err := os.ReadFile("/dev/stdin")
	if err != nil {
		panic(err)
	}

	var sum int

	r := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	r2 := regexp.MustCompile(`[0-9]{1,3}`)
	for _, v := range r.FindAll(input, -1) {
		str := string(v)
		nums := r2.FindAllString(str, 2)

		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}

		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}

		sum += num1 * num2
	}

	fmt.Println(sum)
}
