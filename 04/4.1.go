// AOC Day 4 Part 1
package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	var crossword []string

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				panic(err)
			} else {
				break
			}
		}

		line := scanner.Text()
		crossword = append(crossword, line)
	}

	var count int
	for i, row := range crossword {
		for j, _ := range row {
			if crossword[i][j] == 'X' {
				count += search(&crossword, i, j)
			}
		}
	}

	fmt.Println(count)
}

func search(crossword *[]string, i, j int) int {
	var count int

	c := *crossword
	var max_i int = len(c)
	var max_j int = len(c[0])

	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			ti := i
			tj := j
			for _, v := range []byte{'M', 'A', 'S'} {
				if (ti + di >= max_i) || (ti + di < 0) || (tj + dj >= max_j) || (tj + dj < 0) {
					break
				} else if di == 0 && dj == 0 {
					break
				} else {
					ti += di
					tj += dj
				}

				if c[ti][tj] != v {
					break
				} else if v == 'S' {
					//fmt.Println(i, j, "found")
					count++
				}
			}
		}
	}

	return count
}
