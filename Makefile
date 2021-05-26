

.PHONY: build
build:
	go build

.PHONY: test
test: 
	go test ./...

.PHONY: run
run:
	xcaddy run 
