build:
	@go build -o bin/TelegramBotLibFree

run: build
	@./bin/TelegramBotLibFree

test:
	@go test -v ./...
