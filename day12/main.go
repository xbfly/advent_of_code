package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var validCombinations int

func main() {

	file, _ := os.Open("day12input.txt")

	scanner := bufio.NewScanner(file)

	lines := [][]string{}
	nums := [][]int{}
	for scanner.Scan() {
		b := scanner.Text()
		x := strings.Split(b, " ")

		ns := []string{}
		for _, c := range x[0] {
			ns = append(ns, string(c))
		}
		lines = append(lines, ns)

		n := strings.Split(x[1], ",")
		ms := []int{}
		for _, d := range n {
			y, _ := strconv.Atoi(d)
			ms = append(ms, y)
		}
		nums = append(nums, ms)

	}

	for l := 0; l < len(lines); l++ {
		n, r, qi := getElements(lines[l], nums[l])
		allCombinations := combinations(n, r)

		for _, a := range allCombinations {
			tmpLine := make([]string, len(lines[l]))
			copy(tmpLine, lines[l])
			for _, b := range a {
				z := qi[b]
				tmpLine[z] = "#"
			}

			for i, t := range tmpLine {
				if t == "." || t == "?" {
					tmpLine[i] = " "
				}
			}

			j := strings.Join(tmpLine, "")
			newTmpLine := strings.Fields(j)
			if len(newTmpLine) != len(nums[l]) {
				continue
			}
			valid := true
			for i, w := range newTmpLine {
				if len(w) != nums[l][i] {
					valid = false
				}

			}
			if valid {
				validCombinations++
			}
		}
	}

	fmt.Println("valid combinations: ", validCombinations)

}

func getElements(s []string, i []int) (n, r int, qIndex []int) {
	hashes := 0
	questions := 0
	qIndex = []int{}
	for i, r := range s {
		if string(r) == "#" {
			hashes++
		}
		if string(r) == "?" {
			qIndex = append(qIndex, i)
			questions++
		}
	}

	sum := 0
	for _, d := range i {
		sum += d
	}

	n = questions
	r = sum - hashes
	return n, r, qIndex
}

// combinations generates all combinations of choosing r elements from a set of n elements.
// It returns a slice of slices, where each inner slice represents a combination.
func combinations(n, r int) [][]int {
	var result [][]int
	s := make([]int, r)
	for i := range s {
		s[i] = i
	}

	var helper func(int, int)
	helper = func(i, next int) {
		if i == r {
			// Create a copy of the slice to avoid overwriting
			temp := make([]int, r)
			copy(temp, s)
			result = append(result, temp)
			return
		}
		for j := next; j < n; j++ {
			s[i] = j
			helper(i+1, j+1)
		}
	}
	helper(0, 0)
	return result
}
