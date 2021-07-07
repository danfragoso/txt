all: run

run:
	@go run cmd/*.go

test:
	@go test -v .

build:
	@go build -o txt cmd/*.go


