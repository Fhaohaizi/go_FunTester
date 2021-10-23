PROJECT="example"

all: build

default: build

.PHONY: build
build:
	go build
