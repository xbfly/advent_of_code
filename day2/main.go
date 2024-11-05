package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("day2input.txt")

	scanner := bufio.NewScanner(file)

	powers := []int{}

	for scanner.Scan() {
		x := strings.Split(scanner.Text(), ":")

		z := strings.Split(x[1], ";")
		powers = append(powers, ghost(z))
	}

	sum := 0
	for _, t := range powers {
		sum += t
	}

	fmt.Println("sum: ", sum)
}

func ghost(s []string) int {
	b := 0
	g := 0
	r := 0

	for _, w := range s {
		u := strings.Split(w, ",")
		fmt.Println("U: ", u)

		for _, v := range u {
			vv := strings.TrimLeft(v, " ")
			vp := strings.Split(vv, " ")

			fmt.Println("VV: ", vv)
			fmt.Println("VV[0]: ", vp[0])
			num, _ := strconv.Atoi(vp[0])
			fmt.Println("G: ", g)
			if vp[1] == "blue" && num > b {
				b = num
			}
			if vp[1] == "green" && num > g {
				g = num
			}
			if vp[1] == "red" && num > r {
				r = num
			}

		}
	}

	return r * b * g

}
