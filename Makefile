.PHONY: fmt clean build run

fmt:
	@go fmt ./...

clean:
	@rm -rf glox

build: clean
	@go build -o glox main.go

run: build
	@./glox