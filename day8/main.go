package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Path struct {
	Left, Right string
}

func main() {

	file, _ := os.Open("day8input.txt")

	scanner := bufio.NewScanner(file)

	nodes := make(map[string]Path)

	var directions string
	var sp []string

	for scanner.Scan() {
		if !strings.Contains(scanner.Text(), "=") && scanner.Text() != "" {
			directions = scanner.Text()
		}

		if strings.Contains(scanner.Text(), "=") {
			x := strings.Split(scanner.Text(), " = ")

			n := x[0]
			if strings.HasSuffix(n, "A") {
				sp = append(sp, n)
			}

			p := strings.Split(x[1], ", ")
			pL := strings.Trim(p[0], "(")
			pR := strings.Trim(p[1], ")")

			nodes[n] = Path{
				Left:  pL,
				Right: pR,
			}

		}

	}

	// Find the lengths of cycles for each starting node
	cycleLengths := make([]int, 0, len(sp))
	for _, startNode := range sp {
		count := 0
		currentNode := startNode

		for {

			for _, d := range directions {
				count++
				path := nodes[currentNode]
				if d == 'R' {
					currentNode = path.Right
				} else if d == 'L' {
					currentNode = path.Left
				}
				if strings.HasSuffix(currentNode, "Z") {
					cycleLengths = append(cycleLengths, count)
					break
				}

			}

			if strings.HasSuffix(currentNode, "Z") {
				break // Exit outer loop when "Z" is reached
			}
		}
	}

	fmt.Println("cycle lengths: ", cycleLengths)
	result := cycleLengths[0]
	for _, length := range cycleLengths[1:] {
		result = lcm(result, length)
	}

	fmt.Println("count:", result)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return (a * b) / gcd(a, b)
}
