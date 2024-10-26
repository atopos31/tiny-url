pro: init
	@go build -o ./bin/tiny-url .
	@nohup ./bin/tiny-url > /dev/null 2>&1 &

init:
	@go mod tidy
	@go fmt ./...
	