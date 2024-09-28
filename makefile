build: 
	@go build -o bin/h3

run: build
	@./bin/h3

test:
	@go test ./... -v