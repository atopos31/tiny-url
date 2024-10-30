pro: init
	@go build -o ./bin/tiny-url .
	@nohup ./bin/tiny-url > /dev/null 2>&1 &

docker:
	@docker compose up --build -d

init:
	@go mod tidy
	@go fmt ./...
	