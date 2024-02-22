package app

import (
	"log/slog"

	"github.com/wizzardich/autoreplace-reincarnated/domain"
)

type UpdateTemplateUsecase struct {
	logger *slog.Logger
	repo   domain.TemplateRepository
}

func NewUpdateTemplateUsecase(logger *slog.Logger, repo domain.TemplateRepository) *UpdateTemplateUsecase {
	return &UpdateTemplateUsecase{
		logger: logger,
		repo:   repo,
	}
}

func (u *UpdateTemplateUsecase) Execute(id string, updateTemplate domain.Template) error {
	template, err := u.repo.FindByID(id)

	if err != nil {
		u.logger.Error("could not find template",
			slog.String("error", err.Error()),
			slog.String("template", template.Name),
		)
		return err
	}

	err = u.repo.UpdateTemplate(id, updateTemplate)

	if err != nil {
		u.logger.Error("could not update template",
			slog.String("error", err.Error()),
			slog.String("template", template.Name),
		)
		return err
	}

	return nil
}
