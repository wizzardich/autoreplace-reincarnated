package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	loggerOpts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, loggerOpts))
	controller := NewController(logger)

	mux.HandleFunc("GET /templates/", controller.ListTemplates)
	mux.HandleFunc("GET /templates/{name}", controller.GetTemplateByName)

	logger.Info("server starting at :8090")
	http.ListenAndServe("localhost:8090", mux)
}
