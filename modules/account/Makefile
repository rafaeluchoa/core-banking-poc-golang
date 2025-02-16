.PHONY: test
test:
	go test ./...

.PHONY: cover
cover:
	go test ./... -coverprofile=./build/coverage.out
	go tool cover -func=./build/coverage.out -o ./build/coverage.txt


.PHONY: build
build:
	go build -o ./bin/account