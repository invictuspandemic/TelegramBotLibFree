build:
	@go build cmd/telegram-bot-lib-free/main.go -o bin/TelegramBotLibFree

run: build
	@./bin/TelegramBotLibFree

test:
	@go test -v ./...
