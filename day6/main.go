package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, _ := os.Open("day6input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	time := []int{}
	dist := []int{}

	for scanner.Scan() {
		x := strings.Split(scanner.Text(), ":")
		// y := strings.TrimSpace(x[1])
		y := strings.Fields(x[1])
		dd := ""
		for _, w := range y {
			dd += w
		}

		zc, _ := strconv.Atoi(dd)

		if strings.Contains(x[0], "Time") {
			time = append(time, zc)
		}
		if strings.Contains(x[0], "Distance") {
			dist = append(dist, zc)
		}
	}

	b := []int{}

	fmt.Println("time: ", time)
	fmt.Println("distance: ", dist)

	for i := 0; i < len(time); i++ {
		m := buildTimeDistanceTable(time[i], dist[i])
		b = append(b, m)
	}

	fmt.Println("b: ", b)

	product := 1
	for _, v := range b {
		product *= v
	}

	fmt.Println(product)

}

func buildTimeDistanceTable(t, d int) int {

	// tdTable := map[int]int{}
	total := 0

	for i := 0; i <= t; i++ {
		time := t - i
		dist := time * i

		// if dist == (t-i+1)*(i+1) {
		// 	return total * 2
		// }

		if dist > d {
			total++
		}

	}
	return total
}
