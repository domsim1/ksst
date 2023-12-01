BINARY_NAME=ksst
KSST_PATH=cmd/ksst-gui/main.go
KSST_ENCODER_PATH=cmd/ksst-encoder/main.go



.PHONY: build build-encoder clean run run-encoder

build:
	go mod tidy
	go build -o bin/${BINARY_NAME} ${KSST_PATH}

build-encoder:
	go mod tidy
	go build -o bin/${BINARY_NAME}-encoder ${KSST_ENCODER_PATH}

clean:
	go clean
	rm -rf bin

run:
	go mod tidy
	go run ${KSST_PATH}

run-encoder:
	go mod tidy
	go run ${KSST_ENCODER_PATH}

