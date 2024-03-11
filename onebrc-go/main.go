package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

type result struct {
	city  string
	min   float64
	max   float64
	mean  float64
	acc   float64
	count int
}

func compute(scanner *bufio.Scanner) (map[string]result, error) {
	m := make(map[string]result, 100)
	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), ";")
		city := vals[0]
		temp := must(strconv.ParseFloat(vals[1], 64))
		if res, has := m[city]; has {
			m[city] = result{city: city, min: temp, max: temp, acc: 1, count: 1}
		} else {
			res.count++
			res.acc += temp
			res.min = math.Min(res.min, temp)
			res.max = math.Max(res.max, temp)
			res.mean = res.acc / float64(res.count)
			res.min = math.Min(res.min, temp)
			m[city] = res
		}
	}
	return m, nil
}

func main() {
	filePath := "./data/weather_stations.csv"
	f := must(os.Open(filePath))
	defer f.Close()
	scanner := bufio.NewScanner(bufio.NewReader(f))
	m := must(compute(scanner))
	fmt.Println(m)
}
