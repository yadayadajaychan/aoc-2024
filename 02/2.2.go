// AOC Day 2 Part 2
package main

import (
	"io"
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
	"slices"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var count int
	for {
		nums, err := readLine(scanner)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		if isSafe(nums) {
			count++
		} else {
			for i, _ := range nums {
				deleted := make([]int, len(nums))
				copy(deleted, nums)
				deleted = slices.Delete(deleted, i, i+1)
				if isSafe(deleted) {
					count++
					break
				}
			}
		}
	}

	fmt.Println(count)
}

func isSafe(nums []int) bool {
	var increasing bool
	var safe bool = true
	for i := 1; i < len(nums); i++ {
		if !diffSafe(nums[i], nums[i-1]) {
			safe = false
			break
		}
		if i == 1 {
			if nums[i] - nums[i-1] > 0 {
				increasing = true
			} else {
				increasing = false
			}
		} else {
			if !increasing && (nums[i] - nums[i-1] > 0) || increasing && (nums[i] - nums[i-1] < 0) {
				safe = false
				break
			}
		}
	}

	return safe
}

func diffSafe(a int, b int) bool {
	diff := a - b
	if diff < 0 {
		diff = diff * -1
	}

	if diff >= 1 && diff <= 3 {
		return true
	} else {
		return false
	}
}

func readLine(scanner *bufio.Scanner) ([]int, error) {
	if !scanner.Scan() {
		if err:= scanner.Err(); err != nil {
			return nil, err
		} else {
			return nil, io.EOF
		}
	}

	line := scanner.Text()
	nums := strings.Split(line, " ")

	var output []int
	for _, v := range nums {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		output = append(output, n)
	}
	return output, nil
}
