GO = go
OUTPUT = gort

.PHONY:
	all

all: build run

build:
	$(GO) build -o $(OUTPUT) .

run:
	./$(OUTPUT)
