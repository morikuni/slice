.PHONY: mod
mod:
	go mod download

.PHONY: test
test:
	go test -v -rave -count 1 ./...

.PHONY: bench
bench:
	go test -bench . -benchmem -v -count 1 ./...
