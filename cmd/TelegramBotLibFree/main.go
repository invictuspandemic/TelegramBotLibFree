package main

import (
	"TelegramBotLibFree/internal/config"
	"TelegramBotLibFree/internal/lib/logger/handlers/slogpretty"
	"log"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

	err, cfg := config.LoadCFG()
	if err != nil {
		log.Fatal(err)
	}

	log := setupLogger(cfg.Logger)

	log.Info("Starting TelegramBotLibFree", slog.String("env", cfg.Env))
	log.Debug("Debug message are enabled")

	// TODO: run server

}

func setupLogger(loggerCFG config.Logger) *slog.Logger {
	var log *slog.Logger
	switch loggerCFG.LogType {
	case "pretty":
		log = setupPrettySlog()
	case "text":
		log = slog.New(
			slog.NewTextHandler(
				os.Stdout, &slog.HandlerOptions{
					Level: slog.Level(loggerCFG.LogLevel),
				},
			),
		)
	case "json":
		log = slog.New(
			slog.NewJSONHandler(
				os.Stdout, &slog.HandlerOptions{
					Level: slog.Level(loggerCFG.LogLevel),
				},
			),
		)
	default:
		log = slog.New(
			slog.NewJSONHandler(
				os.Stdout, &slog.HandlerOptions{
					Level: slog.LevelInfo,
				},
			),
		)
	}
	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
