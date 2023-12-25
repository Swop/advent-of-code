# -*- mode: makefile -*-
CURRENT_YEAR = 2023

.PHONY: test
test:
	go test ./...

.PHONY: test-no-cache
test-no-cache:
	go test ./... -count=1

.PHONY: lint
lint:
	golangci-lint run

.PHONY: bench
bench:
	go test ./... -run=XXX -bench=. # XXX allows to filter out tests and only run benchmarks

.PHONY: run-all
run-all:
	go run cmd/run_all.go -path="." -year=$(CURRENT_YEAR)
