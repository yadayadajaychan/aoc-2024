// AOC Day 4 Part 2
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
			if crossword[i][j] == 'A' {
				count += search(&crossword, i, j)
			}
		}
	}

	fmt.Println(count)
}

func search(crossword *[]string, i, j int) int {
	type character struct {
		i  []int
		j  []int
	}

	coor := make(map[byte]character)
	coor['M'] = character{i: []int{}, j: []int{}}
	coor['S'] = character{i: []int{}, j: []int{}}

	c := *crossword
	var max_i int = len(c)
	var max_j int = len(c[0])

	for di := -1; di <= 1; di += 2 {
		for dj := -1; dj <= 1; dj += 2 {
			ti := i
			tj := j

			if (ti + di >= max_i) || (ti + di < 0) || (tj + dj >= max_j) || (tj + dj < 0) {
				break
			} else {
				ti += di
				tj += dj
			}

			char := c[ti][tj]
			if char == 'M' || char == 'S' {
				si := coor[char].i
				sj := coor[char].j

				si = append(si, di)
				sj = append(sj, dj)

				coor[char] = character{i: si, j: sj}
			}
			
		}
	}

	for _, v := range []byte{'M', 'S'} {
		if len(coor[v].i) != 2 {
			return 0
		}

		if !anyEqual(coor[v].i, coor[v].j) {
			return 0
		}
	}

	return 1
}

func anyEqual(i, j []int) bool {
	for k, _ := range i {
		if i[k] == j[k] {
			return true
		}
	}

	return false
}
