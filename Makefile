.PHONY: fmt clean build run

fmt:
	@go fmt ./...

clean:
	@rm -rf glox

build:clean
	@go build main.go

run: build
	@./glox