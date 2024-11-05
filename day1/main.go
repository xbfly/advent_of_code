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
		var nums string

		var tmpLeft string
		for _, l := range n {
			lbreak := false
			if unicode.IsNumber(l) {
				nums += string(l)
				lbreak = true
			} else {
				tmpLeft += string(l)
				for p := range ghost {
					if strings.Contains(tmpLeft, p) {
						nums += ghost[p]
						lbreak = true
					}
				}
			}

			if lbreak {
				break
			}

		}

		var tmpRight string
		right := len(n) - 1
		for i := right; i >= 0; i-- {
			rbreak := false
			if unicode.IsNumber(rune(n[i])) {
				nums += string(n[i])
				rbreak = true
			} else {
				tmpRight = string(n[i]) + tmpRight
				for p := range ghost {
					if strings.Contains(tmpRight, p) {
						nums += ghost[p]
						rbreak = true
					}
				}
			}
			if rbreak {
				break
			}
		}

		x, _ := strconv.Atoi(nums)
		newNums = append(newNums, x)
	}

	sum := 0

	for _, v := range newNums {
		sum += v
	}
	fmt.Println(sum)
}
