BINARY_NAME=myapp

build:
	go build -o ${BINARY_NAME} cmd/stopWords.go cmd/main.go

compile:
	go mod tidy