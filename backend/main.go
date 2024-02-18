package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/wizzardich/autoreplace-reincarnated/app"
	"github.com/wizzardich/autoreplace-reincarnated/infra"
	inframongo "github.com/wizzardich/autoreplace-reincarnated/infra/mongo"
)

func main() {
	mux := http.NewServeMux()

	loggerOpts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, loggerOpts))
	mongoDriver := inframongo.NewMongoDriver(logger, "localhost")
	mongoTemplateRepository := infra.NewTemplateMongoRepository(logger, mongoDriver)

	controller := infra.NewController(
		logger,
		app.NewFindTemplateByIDUsecase(logger, mongoTemplateRepository),
		app.NewRetrieveTemplatesUsecase(logger, mongoTemplateRepository),
	)

	mux.HandleFunc("GET /templates/", controller.ListTemplates)
	mux.HandleFunc("GET /templates/{id}", controller.GetTemplateByID)

	logger.Info("server starting at :8090")
	http.ListenAndServe("localhost:8090", mux)
}
