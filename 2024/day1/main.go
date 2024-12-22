package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	part = flag.String("part", "0", "Which part")
)

func main() {
	file, _ := os.Open("day1input.txt")
	defer file.Close()

	flag.Parse()
	if *part == "1" {
		part1(file)
	}
	if *part == "2" {
		part2(file)
	}

}

func part1(file *os.File) {
	numLeft := []int{}
	numRight := []int{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		x := strings.Split(scanner.Text(), "   ")
		nL, _ := strconv.Atoi(x[0])
		numLeft = append(numLeft, nL)

		nR, _ := strconv.Atoi(x[1])
		numRight = append(numRight, nR)
	}

	sort.Ints(numLeft)
	sort.Ints(numRight)

	sum := 0

	for y := range numLeft {
		if numLeft[y] > numRight[y] {
			sum += numLeft[y] - numRight[y]
		} else {
			sum += numRight[y] - numLeft[y]
		}
	}

	fmt.Println(sum)

}

func part2(file *os.File) {
	numLeft := []int{}
	numRightMap := map[int]int{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		x := strings.Split(scanner.Text(), "   ")
		nL, _ := strconv.Atoi(x[0])
		numLeft = append(numLeft, nL)

		nR, _ := strconv.Atoi(x[1])
		numRightMap[nR]++
	}

	sum := 0

	for _, y := range numLeft {
		if l, ok := numRightMap[y]; ok {
			sum += y * l
		}
	}

	fmt.Println(sum)
}
