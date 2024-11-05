package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, _ := os.Open("day1input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	number := []string{}

	for scanner.Scan() {
		number = append(number, scanner.Text())
	}

	newNums := []int{}

	ghost := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for _, n := range number {
		right := len(n) - 1
		var tmpLeft string
		var tmpRight string

		var tmp string

		lbreak := false
		rbreak := false

		for _, l := range n {
			if unicode.IsNumber(l) {
				tmp += string(l)
				lbreak = true
			} else {
				tmpLeft += string(l)
				for p := range ghost {
					if strings.Contains(tmpLeft, p) {
						tmp += ghost[p]
						lbreak = true
					}
				}
			}

			if lbreak {
				break
			}

		}

		for i := right; i >= 0; i-- {
			if unicode.IsNumber(rune(n[i])) {
				tmp += string(n[i])
				rbreak = true
			} else {
				tmpRight = string(n[i]) + tmpRight
				fmt.Println("tmpRight: ", tmpRight)
				for p := range ghost {
					if strings.Contains(tmpRight, p) {
						tmp += ghost[p]
						rbreak = true
					}
				}
			}
			if rbreak {
				break
			}
		}

		x, _ := strconv.Atoi(tmp)
		newNums = append(newNums, x)
	}

	sum := 0

	for _, v := range newNums {
		sum += v
	}
	fmt.Println(sum)
}
