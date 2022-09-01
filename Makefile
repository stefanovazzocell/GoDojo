.PHONY: bench
bench:
	@echo "Running tests..."
	go test -run=Test. -race ./bench
	@echo "Running benchmarks..."
	go test -run=Bench -bench=. ./bench