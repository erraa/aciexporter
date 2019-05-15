clean-build:

install:
	go install ./...

build:
	go build ./...

test:
	go test ./... -v

run:
	go run main.go
