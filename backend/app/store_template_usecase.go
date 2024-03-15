package app

import (
	"log/slog"

	"github.com/wizzardich/autoreplace-reincarnated/domain"
)

type StoreTemplateUsecase struct {
	logger *slog.Logger
	repo   domain.TemplateRepository
}

func NewStoreTemplateUsecase(logger *slog.Logger, repo domain.TemplateRepository) *StoreTemplateUsecase {
	return &StoreTemplateUsecase{
		logger: logger,
		repo:   repo,
	}
}

func (u *StoreTemplateUsecase) Execute(template domain.Template) (*domain.IDString, error) {
	_, err := u.repo.FindByName(template.Name)

	if err != domain.ErrTemplateNotFound {
		return nil, domain.ErrTemplateAlreadyExists
	}

	id, err := u.repo.StoreTemplate(template)

	if err != nil {
		u.logger.Error("could not save template")
		return nil, err
	}

	return id, nil
}
