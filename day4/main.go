package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, _ := os.Open("day4input.txt")
	scanner := bufio.NewScanner(file)

	cards := map[int]int{}

	row := 1
	for scanner.Scan() {
		x := strings.Split(scanner.Text(), ":")
		y := strings.Split(x[1], "|")

		cards[row]++
		cardCount := calculatePoints(y)

		for b := 1; b <= cardCount; b++ {
			cards[row+b] = cards[row+b] + cards[row]
		}

		row++
	}

	sum := 0
	for _, v := range cards {
		sum += v
	}
	fmt.Println(sum)

}

func calculatePoints(s []string) int {
	z1 := s[0]
	z1 = strings.TrimSpace(z1)
	winners := strings.Split(z1, " ")

	chicken := map[int]bool{}

	for _, w := range winners {
		if w == "" {
			continue
		}
		n, _ := strconv.Atoi(w)

		chicken[n] = true
	}

	count := 0

	z2 := s[1]
	z2 = strings.TrimSpace(z2)
	ours := strings.Split(z2, " ")

	for _, w := range ours {
		if w == "" {
			continue
		}
		n, _ := strconv.Atoi(w)

		if chicken[n] {
			count++
		}
	}
	return count
}
