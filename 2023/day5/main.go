package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	seeds                 = []int{}
	seedToSoil            = map[int]int{}
	soilToFertilizer      = map[int]int{}
	fertilizerToWater     = map[int]int{}
	waterToLight          = map[int]int{}
	lightToTemperature    = map[int]int{}
	temperatureToHumidity = map[int]int{}
	humidityToLocation    = map[int]int{}
)

func main() {
	file, _ := os.Open("day5input.txt")
	scanner := bufio.NewScanner(file)

	mapMap := map[string]map[int]int{
		"seed-to-soil map:":            seedToSoil,
		"soil-to-fertilizer map:":      soilToFertilizer,
		"fertilizer-to-water map:":     fertilizerToWater,
		"water-to-light map:":          waterToLight,
		"light-to-temperature map:":    lightToTemperature,
		"temperature-to-humidity map:": temperatureToHumidity,
		"humidity-to-location map:":    humidityToLocation,
	}
	var m map[int]int
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "seeds:") {
			x := strings.Split(scanner.Text(), ":")
			y := strings.TrimSpace(x[1])
			z := strings.Split(y, " ")

			for _, s := range z {
				a, _ := strconv.Atoi(s)
				seeds = append(seeds, a)
			}
			continue
		}
		if _, ok := mapMap[scanner.Text()]; ok {
			m = mapMap[scanner.Text()]
			continue
		}
		if scanner.Text() != "" {
			calculateMap(scanner.Text(), m)
		}

	}

	var soil, fertilizer, water, light, temperature, humidity, location int
	locations := []int{}

	for _, seed := range seeds {
		if check(seed, seedToSoil) {
			soil = seedToSoil[seed]
		} else {
			soil = seed
		}

		if check(soil, soilToFertilizer) {
			fertilizer = soilToFertilizer[soil]
		} else {
			fertilizer = soil
		}

		if check(fertilizer, fertilizerToWater) {
			water = fertilizerToWater[fertilizer]
		} else {
			water = fertilizer
		}

		if check(water, waterToLight) {
			light = waterToLight[water]
		} else {
			light = water
		}

		if check(light, lightToTemperature) {
			temperature = lightToTemperature[light]
		} else {
			temperature = light
		}

		if check(temperature, temperatureToHumidity) {
			humidity = temperatureToHumidity[temperature]
		} else {
			humidity = temperature
		}

		if check(humidity, humidityToLocation) {
			location = humidityToLocation[temperature]
		} else {
			location = humidity
		}

		locations = append(locations, location)
	}

	fmt.Println("locations; ", locations)

	lowestLocation := int(math.Inf(1))
	for _, l := range locations {
		if l < lowestLocation {
			lowestLocation = l
		}
	}

	fmt.Println("lowest location: ", lowestLocation)
}

func check(i int, m map[int]int) bool {
	if _, ok := m[i]; ok {
		return true
	}
	return false

}

func calculateMap(n string, m map[int]int) {
	s := strings.Split(n, " ")
	dst, _ := strconv.Atoi(s[0])
	src, _ := strconv.Atoi(s[1])
	ran, _ := strconv.Atoi(s[2])

	for i := src; i < src+ran; i++ {
		m[i] = dst
		dst++
	}
}
