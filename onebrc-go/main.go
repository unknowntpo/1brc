package main

import (
	"bufio"
	"fmt"
	"os"
)

func must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

func main() {
	filePath := "./data/weather_stations.csv"
	f := must(os.Open(filePath))
	defer f.Close()
	scanner := bufio.NewScanner(bufio.NewReader(f))
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
