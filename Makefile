.PHONY: test
test:
	go test -v -race -count 1 ./...

.PHONY: coverage
coverage:
	go test -v -race -count 1 -covermode=atomic -coverprofile=coverage.out ./...

.PHONY: bench
bench:
	go test -bench . -benchmem -v -count 1 ./...
