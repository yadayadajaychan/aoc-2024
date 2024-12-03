// AOC Day 3 Part 2
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
	var do bool = true

	r := regexp.MustCompile(`(mul\([0-9]{1,3},[0-9]{1,3}\))|(do\(\))|(don't\(\))`)
	r2 := regexp.MustCompile(`[0-9]{1,3}`)
	for _, v := range r.FindAll(input, -1) {
		str := string(v)
		if str == "do()" {
			do = true
		} else if str == "don't()" {
			do = false
		} else if do == true {
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
	}

	fmt.Println(sum)
}
