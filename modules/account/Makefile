.PHONY: install-dev
install-dev:
	go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: swagger
swagger:
	swag init -g ./api/account.go

.PHONY: build
build: swagger
	go build -o ./bin/account

.PHONY: test
test: build
	go test ./...

.PHONY: cover
cover: build
	go test ./... -coverprofile=./build/coverage.out
	go tool cover -html=./build/coverage.out -o ./build/coverage.html

