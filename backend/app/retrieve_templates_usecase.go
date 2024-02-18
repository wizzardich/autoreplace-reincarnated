package app

import (
	"log/slog"

	"github.com/wizzardich/autoreplace-reincarnated/domain"
)

type RetrieveTemplatesUsecase struct {
	logger *slog.Logger
	repo   domain.TemplateRepository
}

func NewRetrieveTemplatesUsecase(logger *slog.Logger, repo domain.TemplateRepository) *RetrieveTemplatesUsecase {
	return &RetrieveTemplatesUsecase{
		logger: logger,
		repo:   repo,
	}
}

func (u *RetrieveTemplatesUsecase) Execute() ([]domain.Template, error) {
	templates, err := u.repo.FindAll()

	if err != nil {
		u.logger.Error("could not retrieve templates")
		return nil, err
	}

	return templates, nil
}
