BIN_FILE_NAME=blurhash_cli
BUILD_DIR=build

all: build test

build:
	go build -o ${BUILD_DIR}/${BIN_FILE_NAME}

test:
	go test ../...
	./test_script.sh

clean:
	rm -rf ${BUILD_DIR}
