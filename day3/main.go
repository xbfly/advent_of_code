package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Schematic [][]rune

type Position struct {
	row, col int
}

type SymbolMap map[Position]bool

func main() {
	schematic := readschematic("day3input.txt")
	symbols := findSymbols(schematic)
	fmt.Println(calculateGearRatios(schematic, symbols))
}

func findSymbols(schematic Schematic) SymbolMap {
	gearMap := make(SymbolMap)
	for row := 0; row < len(schematic); row++ {
		for col := 0; col < len(schematic[row]); col++ {
			if schematic[row][col] == '*' {
				gearMap[Position{row, col}] = true
			}
		}
	}
	return gearMap
}

func calculateGearRatios(schematic Schematic, gearMap SymbolMap) int {
	totalRatio := 0
	for gearPos := range gearMap {
		adjacentNumbers := extraAdjacentNumbers(schematic, gearPos)
		if len(adjacentNumbers) == 2 {
			ratio := adjacentNumbers[0] * adjacentNumbers[1]
			totalRatio += ratio
		}
	}
	return totalRatio
}

func extraAdjacentNumbers(schematic Schematic, gearPos Position) []int {
	numbers := []int{}
	visited := make(map[Position]bool)

	row, col := gearPos.row, gearPos.col
	for r := row - 1; r <= row+1; r++ {
		for c := col - 1; c <= col+1; c++ {
			if r >= 0 && r < len(schematic) && c >= 0 && c < len(schematic[r]) &&
				unicode.IsDigit(schematic[r][c]) && !visited[Position{r, c}] {

				startCol := c
				for startCol >= 0 && unicode.IsDigit(schematic[r][startCol]) {
					visited[Position{r, startCol}] = true
					startCol--
				}
				startCol++

				endCol := c
				for endCol < len(schematic[r]) && unicode.IsDigit(schematic[r][endCol]) {
					visited[Position{r, endCol}] = true
					endCol++
				}
				endCol--

				numStr := schematic[r][startCol : endCol+1]
				number, _ := strconv.Atoi(string(numStr))
				numbers = append(numbers, number)
			}
		}
	}
	return numbers
}

func readschematic(filename string) Schematic {
	file, _ := os.Open(filename)
	defer file.Close()

	var schematic Schematic
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		schematic = append(schematic, []rune(scanner.Text()))
	}

	return schematic
}

func sumPartNumbers(schematic Schematic, symbolMap SymbolMap) int {
	sum := 0
	for row := 0; row < len(schematic); row++ {
		for col := 0; col < len(schematic[row]); col++ {
			if unicode.IsDigit(schematic[row][col]) {
				startCol := col
				for col < len(schematic[row]) && unicode.IsDigit(schematic[row][col]) {
					col++
				}
				col--

				numStr := string(schematic[row][startCol : col+1])
				number, _ := strconv.Atoi(numStr)

				isPartNumber := false
				for i := startCol; i <= col; i++ {
					if isAdjacentSymbol(symbolMap, row, i) {
						isPartNumber = true
						break
					}
				}

				if isPartNumber {
					sum += number
				}
			}
		}
	}
	return sum
}

func isAdjacentSymbol(symbolMap SymbolMap, row, col int) bool {
	for r := row - 1; r <= row+1; r++ {
		for c := col - 1; c <= col+1; c++ {
			if symbolMap[Position{r, c}] {
				return true
			}
		}
	}
	return false
}
