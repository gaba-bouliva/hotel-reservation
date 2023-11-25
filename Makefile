build:
	@go build -o bin/api

run: build
	@./bin/api
cp: 
	cp -r . /mnt/c/Users/Gaba/workspace/github.com/hotel-reservation/

test:
	@go test -v ./...
