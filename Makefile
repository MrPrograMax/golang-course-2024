BINARY_NAME=main.out

build:
	go build -o ${BINARY_NAME} cmd/stopWords.go cmd/main.go

compile:
	go mod tidy