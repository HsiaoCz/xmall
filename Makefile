run: build
	@./bin/xmall

build:
	@go build -o bin/xmall main.go

test:
	@go test -v ./...

tidy:
	@go mod tidy
