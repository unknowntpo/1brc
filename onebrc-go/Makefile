
all: help

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==================================================================================== #
# TEST
# ==================================================================================== #

## test: execute tests
.PHONY: test
test:
	@go test -v ./...
## bench: execute benchmark
.PHONY: bench
bench:
	go test -v -short ./... -run=^$$ -bench=. 
