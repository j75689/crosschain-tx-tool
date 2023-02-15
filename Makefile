build:
ifeq ($(OS),Windows_NT)
	go build -o build/crossd.exe main.go
else
	go build -o build/crossd main.go
endif

install:
ifeq ($(OS),Windows_NT)
	go install main.go
else
	go install main.go
endif

.PHONY: build install
