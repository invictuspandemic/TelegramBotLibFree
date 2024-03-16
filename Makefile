build:
	@go build cmd/TelegramBotLibFree/main.go -o bin/TelegramBotLibFree

run: build
	@./bin/TelegramBotLibFree

test:
	@go test -v ./...
