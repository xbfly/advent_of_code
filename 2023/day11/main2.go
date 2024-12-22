package main

import (
	"bufio"
	"fmt"
	"math"
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

	emptyRows := map[int]bool{}
	var universe [][]string

	row := 0
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
			emptyRows[row] = true
		}
		row++

	}

	emptyCols := map[int]bool{}
	for c := 0; c < len(universe[0]); c++ {
		allDots := true
		for r := 0; r < len(universe); r++ {
			if universe[r][c] != "." {
				allDots = false
				break
			}
		}
		if allDots {
			emptyCols[c] = true
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

			countEmptyRow := 0

			minRow := int(math.Min(float64(g1.row), float64(g2.row)))
			maxRow := int(math.Max(float64(g1.row), float64(g2.row)))

			for h := minRow + 1; h < maxRow; h++ {
				if _, ok := emptyRows[h]; ok {
					countEmptyRow++
				}
			}

			d1 := g1.row - g2.row
			if d1 < 0 {
				d1 = -d1
			}

			d1 = d1 - countEmptyRow + (countEmptyRow * 1000000)

			countEmptyCol := 0

			minCol := int(math.Min(float64(g1.col), float64(g2.col)))
			maxCol := int(math.Max(float64(g1.col), float64(g2.col)))

			for k := minCol + 1; k < maxCol; k++ {
				if _, ok := emptyCols[k]; ok {
					countEmptyCol++
				}
			}

			d2 := g1.col - g2.col
			if d2 < 0 {
				d2 = -d2
			}

			d1 = d1 - countEmptyCol + (countEmptyCol * 1000000)

			distance += d1 + d2
		}
	}

	fmt.Println("distance: ", distance)
}
