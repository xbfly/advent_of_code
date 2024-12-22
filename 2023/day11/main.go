package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Position struct {
	row int
	col int
}

func main() {

	file, _ := os.Open("day11input.txt")

	scanner := bufio.NewScanner(file)

	var universe [][]string
	for scanner.Scan() {
		u := strings.Split(scanner.Text(), "")
		universe = append(universe, u)

		// expand rows if empty
		allDots := true
		for _, r := range scanner.Text() {
			if string(r) != "." {
				allDots = false
				break
			}
		}

		if allDots {
			g := strings.Split(scanner.Text(), "")
			universe = append(universe, g)
		}

	}

	for c := 0; c < len(universe[0]); c++ {
		allDots := true
		for r := 0; r < len(universe); r++ {
			if universe[r][c] != "." {
				allDots = false
				break
			}
		}
		if allDots {
			for r := 0; r < len(universe); r++ {
				newSlice := make([]string, len(universe[r])+1)

				copy(newSlice, universe[r][:c])

				dot := "."
				newSlice[c] = dot
				copy(newSlice[c+1:], universe[r][c:])
				universe[r] = newSlice

			}
			c++
		}

	}

	galaxies := []Position{}

	for r := 0; r < len(universe); r++ {
		for c := 0; c < len(universe[0]); c++ {
			if universe[r][c] == "#" {
				p := Position{
					row: r,
					col: c,
				}
				galaxies = append(galaxies, p)
			}
		}
	}

	distance := 0

	for i := 0; i < len(galaxies)-1; i++ {
		for _, n := range galaxies[i+1:] {
			g1 := galaxies[i]
			g2 := n

			d1 := g1.row - g2.row
			if d1 < 0 {
				d1 = -d1
			}

			d2 := g1.col - g2.col
			if d2 < 0 {
				d2 = -d2
			}

			distance += d1 + d2
		}
	}

	fmt.Println("distance: ", distance)
}
