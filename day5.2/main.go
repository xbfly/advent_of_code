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
	seedToSoil            = map[Scheme]bool{}
	soilToFertilizer      = map[Scheme]bool{}
	fertilizerToWater     = map[Scheme]bool{}
	waterToLight          = map[Scheme]bool{}
	lightToTemperature    = map[Scheme]bool{}
	temperatureToHumidity = map[Scheme]bool{}
	humidityToLocation    = map[Scheme]bool{}
)

type Scheme struct {
	dst int
	src int
	rng int
}

func main() {
	file, _ := os.Open("day5input.txt")
	scanner := bufio.NewScanner(file)

	mapMap := map[string]map[Scheme]bool{
		"seed-to-soil map:":            seedToSoil,
		"soil-to-fertilizer map:":      soilToFertilizer,
		"fertilizer-to-water map:":     fertilizerToWater,
		"water-to-light map:":          waterToLight,
		"light-to-temperature map:":    lightToTemperature,
		"temperature-to-humidity map:": temperatureToHumidity,
		"humidity-to-location map:":    humidityToLocation,
	}
	var m map[Scheme]bool
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
	lowestLocation := int(math.Inf(1))

	for i := 0; i < len(seeds); i += 2 {
		for j := 0; j < seeds[i+1]; j++ {
			seed := seeds[i] + j
			soil = check(seed, seedToSoil)
			fertilizer = check(soil, soilToFertilizer)
			water = check(fertilizer, fertilizerToWater)
			light = check(water, waterToLight)
			temperature = check(light, lightToTemperature)
			humidity = check(temperature, temperatureToHumidity)
			location = check(humidity, humidityToLocation)

			if location < lowestLocation {
				lowestLocation = location
			}
		}

	}

	fmt.Println("lowest location: ", lowestLocation)
}

func check(i int, m map[Scheme]bool) int {
	for r := range m {
		if i >= r.src && i < r.src+r.rng {
			d := i - r.src
			return r.dst + d
		}
	}

	return i
}

func calculateMap(n string, m map[Scheme]bool) {
	s := strings.Split(n, " ")
	dst, _ := strconv.Atoi(s[0])
	src, _ := strconv.Atoi(s[1])
	rng, _ := strconv.Atoi(s[2])

	sch := Scheme{
		dst: dst,
		src: src,
		rng: rng,
	}

	m[sch] = true

}
