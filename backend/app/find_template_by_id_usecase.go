package app

import (
	"log/slog"

	"github.com/wizzardich/autoreplace-reincarnated/domain"
)

type FindTemplateByIDUsecase struct {
	logger *slog.Logger
	repo   domain.TemplateRepository
}

func NewFindTemplateByIDUsecase(logger *slog.Logger, repo domain.TemplateRepository) *FindTemplateByIDUsecase {
	return &FindTemplateByIDUsecase{
		logger: logger,
		repo:   repo,
	}
}

func (u *FindTemplateByIDUsecase) Execute(ID string) (*domain.Template, error) {
	template, err := u.repo.FindByID(ID)

	if err != nil {
		u.logger.Error("could not retrieve template")
		return nil, err
	}

	return &template, nil
}
