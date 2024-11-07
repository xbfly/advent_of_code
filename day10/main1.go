package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	row, col int
}

var visited map[Position]bool
var scheme map[Position]string

func main() {
	visited = make(map[Position]bool)
	scheme = make(map[Position]string)

	file, _ := os.Open("day10input.txt")

	scanner := bufio.NewScanner(file)

	var startingPoint Position

	i := 0
	for scanner.Scan() {
		row := scanner.Text()

		for j := 0; j < len(row); j++ {
			p := Position{
				row: i,
				col: j,
			}
			scheme[p] = string(row[j])

			if string(row[j]) == "S" {
				startingPoint = p
			}
		}
		i++
	}

	visited[startingPoint] = true
	checkPosition(startingPoint)
	fmt.Println("farthest point away: ", len(visited)/2)

}

func checkPosition(p Position) {
	// north
	// n = |, 7, F

	north := Position{p.row - 1, p.col}
	if !visited[north] {
		if scheme[p] != "-" && scheme[p] != "F" && scheme[p] != "7" && (scheme[north] == "|" || scheme[north] == "7" || scheme[north] == "F") {
			visited[north] = true
			checkPosition(north)
		}
	}

	//south
	// s = |, L, J
	south := Position{p.row + 1, p.col}
	if !visited[south] {
		if scheme[p] != "-" && scheme[p] != "L" && scheme[p] != "J" && (scheme[south] == "|" || scheme[south] == "L" || scheme[south] == "J") {
			visited[south] = true
			checkPosition(south)
		}
	}

	//east
	// e = -, J, 7
	east := Position{p.row, p.col + 1}
	if !visited[east] {
		if scheme[p] != "|" && scheme[p] != "7" && scheme[p] != "J" && (scheme[east] == "-" || scheme[east] == "J" || scheme[east] == "7") {
			visited[east] = true
			checkPosition(east)
		}
	}

	//west
	// w = -, L, F
	west := Position{p.row, p.col - 1}
	if !visited[west] {
		if scheme[p] != "|" && scheme[p] != "L" && scheme[p] != "F" && (scheme[west] == "-" || scheme[west] == "L" || scheme[west] == "F") {
			visited[west] = true
			checkPosition(west)
		}
	}

}
