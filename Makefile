include .env
export

BINARY_NAME=ftx-bot

build:
	go build -o ${BINARY_NAME} main.go

run: build
	./${BINARY_NAME}

watch-dev:
	reflex \
		-s -r '\.go$$|Makefile|.env' \
		-- make run
