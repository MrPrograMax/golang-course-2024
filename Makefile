BINARY_NAME=myapp

build:
	go build -o ${BINARY_NAME} .

compile:
	go mod tidy