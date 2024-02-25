BENCHMARKS = .
BENCHMARK_COUNT = 1
bench:
	go test -run '^$$' -bench=$(BENCHMARKS) -benchmem -cpuprofile=cpuprofile.out -count=$(BENCHMARK_COUNT) ./...
.PHONY: bench

bench-cpu: bench
	go tool pprof -web cpuprofile.out
.PHONY: bench-cpu

bench-compare:
	git stash push *.go
	make bench BENCHMARKS=$(BENCHMARKS) BENCHMARK_COUNT=10 > old.out
	git stash pop
	make bench BENCHMARKS=$(BENCHMARKS) BENCHMARK_COUNT=10 > new.out
	go run golang.org/x/perf/cmd/benchstat@latest old.out new.out