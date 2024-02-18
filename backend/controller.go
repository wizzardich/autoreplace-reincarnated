package main

import (
	"log/slog"
	"net/http"
)

type Controller struct {
	logger *slog.Logger
}

func NewController(logger *slog.Logger) *Controller {
	return &Controller{
		logger: logger,
	}
}

func (c *Controller) ListTemplates(w http.ResponseWriter, r *http.Request) {
	c.logger.Debug("ListTemplates", FullRequest(r))
	// ...
}

func (c *Controller) GetTemplateByName(w http.ResponseWriter, r *http.Request) {
	c.logger.Debug("GetTemplateByName", FullRequest(r))
	// ...
}
