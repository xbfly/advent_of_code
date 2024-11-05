package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	fives      = []string{}
	fours      = []string{}
	fullHouses = []string{}
	threes     = []string{}
	twos       = []string{}
	ones       = []string{}
	highs      = []string{}
)

var trump = map[string]string{
	"A": "D",
	"K": "C",
	"Q": "B",
	"T": "A",
	"9": "9",
	"8": "8",
	"7": "7",
	"6": "6",
	"5": "5",
	"4": "4",
	"3": "3",
	"2": "2",
	"J": "1",
}

func main() {

	file, _ := os.Open("day7input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	major := map[string]int{}

	for scanner.Scan() {
		x := strings.Split(scanner.Text(), " ")
		hand := x[0]
		figureOut(hand)
		bet, _ := strconv.Atoi(x[1])
		major[hand] = bet
	}

	var combined []string
	combined = append(combined, order(highs)...)
	combined = append(combined, order(ones)...)
	combined = append(combined, order(twos)...)
	combined = append(combined, order(threes)...)
	combined = append(combined, order(fullHouses)...)
	combined = append(combined, order(fours)...)
	combined = append(combined, order(fives)...)

	product := 0
	for i := 0; i < len(combined); i++ {
		product += major[combined[i]] * (i + 1)
	}
	fmt.Println(product)

}

func order(m []string) []string {
	var ordered []int

	hexMap := make(map[int]string)

	for i := 0; i < len(m); i++ {
		hex := convertToHex(m[i])
		conInt, _ := strconv.ParseInt(hex, 16, 64)
		hexMap[int(conInt)] = m[i]
		ordered = append(ordered, int(conInt))
	}

	sort.Ints(ordered)

	orderedHands := []string{}

	for _, v := range ordered {
		orderedHands = append(orderedHands, hexMap[v])
	}

	return orderedHands
}

func convertToHex(s string) string {
	var hexVal string

	for i := 0; i < len(s); i++ {
		hv := trump[string(s[i])]
		hexVal += hv
	}

	return hexVal
}

func figureOut(hand string) {

	w := map[string]int{}

	for _, s := range hand {
		w[string(s)]++
	}

	highVal := 0
	var highCard string
	for a, h := range w {
		if h > highVal && a != "J" {
			highVal = h
			highCard = a
		}
	}

	w[highCard] += w["J"]
	delete(w, "J")

	three := 0
	two := 0

	for _, t := range w {
		switch t {
		case 5:
			fives = append(fives, hand)
			return
		case 4:
			fours = append(fours, hand)
			return
		case 3:
			three++
		case 2:
			two++
		}
	}

	if three > 0 && two > 0 {
		fullHouses = append(fullHouses, hand)
		return
	}
	if three > 0 {
		threes = append(threes, hand)
		return
	}
	if two == 2 {
		twos = append(twos, hand)
		return
	}
	if two == 1 {
		ones = append(ones, hand)
		return
	}
	highs = append(highs, hand)
}
