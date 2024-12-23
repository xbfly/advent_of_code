package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	validCombinations int
	// Memoization cache for combinations
	combinationsCache = make(map[[2]int][][]int)
)

func main() {
	file, _ := os.Open("day12inputmod.txt")

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

		sampleNS := make([]string, len(ns))
		copy(sampleNS, ns)
		for addLines := 0; addLines < 4; addLines++ {
			ns = append(ns, "?")
			for _, al := range sampleNS {
				ns = append(ns, al)
			}
		}
		lines = append(lines, ns)

		n := strings.Split(x[1], ",")
		ms := []int{}
		for _, d := range n {
			y, _ := strconv.Atoi(d)
			ms = append(ms, y)
		}

		sampleMS := make([]int, len(ms))
		copy(sampleMS, ms)
		for addNums := 0; addNums < 4; addNums++ {
			for _, an := range sampleMS {
				ms = append(ms, an)
			}
		}
		nums = append(nums, ms)

	}
	// fmt.Println("lines; ", lines)
	// fmt.Println("nums: ", nums)

	for l := 0; l < len(lines); l++ {
		n, r, qi := getElements(lines[l], nums[l])

		key := [2]int{n, r}
		if _, ok := combinationsCache[key]; !ok {
			combinationsCache[key] = combinations(n, r) // Cache the combinations
		}
		allCombinations := combinationsCache[key]

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
