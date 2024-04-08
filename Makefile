BINARY_NAME=xkcd

build:
	go build -o ${BINARY_NAME} ./cmd/xkcd/main.go

run:
	go run cmd/xkcd/main.go -o -n 10

compile:
	go mod tidy