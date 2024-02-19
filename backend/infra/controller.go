package infra

import (
	"log/slog"
	"net/http"

	"github.com/wizzardich/autoreplace-reincarnated/app"
	"github.com/wizzardich/autoreplace-reincarnated/domain"
	"github.com/wizzardich/autoreplace-reincarnated/dto"
	"github.com/wizzardich/autoreplace-reincarnated/util"
)

type Controller struct {
	logger                   *slog.Logger
	findTemplateByIDUsecase  *app.FindTemplateByIDUsecase
	retrieveTemplatesUsecase *app.RetrieveTemplatesUsecase
}

func NewController(logger *slog.Logger, findTemplateByIDUsecase *app.FindTemplateByIDUsecase, retrieveTemplatesUsecase *app.RetrieveTemplatesUsecase) *Controller {
	return &Controller{
		logger:                   logger,
		findTemplateByIDUsecase:  findTemplateByIDUsecase,
		retrieveTemplatesUsecase: retrieveTemplatesUsecase,
	}
}

func (c *Controller) ListTemplates(w http.ResponseWriter, r *http.Request) {
	c.logger.Debug("ListTemplates", util.FullRequest(r))
	templates, err := c.retrieveTemplatesUsecase.Execute()

	if err != nil {
		c.logger.Error("could not retrieve templates", err)
		http.Error(w, "could not retrieve templates", http.StatusInternalServerError)
		return
	}

	templateIDs := make([]dto.Template, len(templates))
	for i, t := range templates {
		replacements := make([]dto.Replacement, len(t.Replacements))
		for _, r := range t.Replacements {
			replacements = append(replacements, dto.Replacement{
				From: r.From,
				To:   r.To,
			})
		}

		templateIDs[i] = dto.Template{
			ID:           t.ID,
			Name:         t.Name,
			Replacements: replacements,
		}
	}

	w.WriteHeader(http.StatusOK)
	util.WriteJSON(w, templateIDs)
}

func (c *Controller) GetTemplateByID(w http.ResponseWriter, r *http.Request) {
	c.logger.Debug("GetTemplateByID", util.FullRequest(r))
	id := r.PathValue("id")

	if id == "" {
		c.logger.Error("invalid template ID")
		http.Error(w, "invalid template ID", http.StatusBadRequest)
		return
	}

	template, err := c.findTemplateByIDUsecase.Execute(id)

	if err == domain.ErrTemplateNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		c.logger.Error("could not retrieve template", err)
		http.Error(w, "could not retrieve template", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	util.WriteJSON(w, dto.TemplateToDTO(*template))
}
