.PHONY: test
test:
	go test ./...

.PHONY: cover
cover:
	go test ./... -coverprofile=./build/coverage.out
	go tool cover -html=./build/coverage.out -o ./build/coverage.html


.PHONY: build
build:
	go build -o ./bin/account