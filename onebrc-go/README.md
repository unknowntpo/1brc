# 1brc-go: 1brc challenge in Go.

## Approach 1: Bufio with single goroutine file parsing + aggregation 

```
________________________________________________________
Executed in  180.43 secs    fish           external
   usr time  153.96 secs    0.06 millis  153.96 secs
   sys time   10.33 secs    1.17 millis   10.32 secs

## Approach 2A: bench FileChunkReader: Read with single goroutine

```
$ go test -v -short ./... -run=^$ -bench=.
goos: darwin
goarch: arm64
pkg: github.com/unknowntpo/1brc/onebrc-go
BenchmarkFileChunkReader
BenchmarkFileChunkReader/read_and_print_file_content:_./data/weather_stations.csv
BenchmarkFileChunkReader/read_and_print_file_content:_./data/weather_stations.csv-8             1000000000               0.0002890 ns/op
PASS
ok      github.com/unknowntpo/1brc/onebrc-go    0.356s
```

## Approach 2B: bench FileChunkReader: Read with multiple goroutine

```
go test -v -short ./... -run=^$ -bench=.
goos: darwin
goarch: arm64
pkg: github.com/unknowntpo/1brc/onebrc-go
BenchmarkFileChunkReader
BenchmarkFileChunkReader/read_and_print_file_content:_./data/weather_stations.csv
BenchmarkFileChunkReader/read_and_print_file_content:_./data/weather_stations.csv-8             1000000000               0.001613 ns/op
PASS
ok      github.com/unknowntpo/1brc/onebrc-go    0.297s
```
