package app

import (
	"log/slog"

	"github.com/wizzardich/autoreplace-reincarnated/domain"
)

type DeleteTemplateUsecase struct {
	logger *slog.Logger
	repo   domain.TemplateRepository
}

func NewDeleteTemplateUsecase(logger *slog.Logger, repo domain.TemplateRepository) *DeleteTemplateUsecase {
	return &DeleteTemplateUsecase{
		logger: logger,
		repo:   repo,
	}
}

func (u *DeleteTemplateUsecase) Execute(id string) error {
	err := u.repo.DeleteTemplate(id)

	if err == domain.ErrTemplateNotFound {
		return domain.ErrTemplateNotFound
	}

	if err != nil {
		u.logger.Error("could not delete template",
			slog.String("id", id),
			slog.String("error", err.Error()),
		)
		return err
	}

	return nil
}
