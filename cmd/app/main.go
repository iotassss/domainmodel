package main

import (
	"log/slog"
	"os"

	"github.com/iotassss/domainmodel/internal/infrastructure/router"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		slog.Error("failed to open log file", slog.Any("error", err))
		os.Exit(1)
	}
	defer file.Close()

	logger := slog.New(slog.NewJSONHandler(file, nil))
	slog.SetDefault(logger)

	r := router.NewRouter()

	if err := r.Run(":8080"); err != nil {
		slog.Error("server failed to start", slog.Any("error", err))
		os.Exit(1)
	}
}
