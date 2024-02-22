package infra

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/wizzardich/autoreplace-reincarnated/domain"
	"github.com/wizzardich/autoreplace-reincarnated/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	inframongo "github.com/wizzardich/autoreplace-reincarnated/infra/mongo"
)

type TemplateMongoRepository struct {
	logger *slog.Logger
	driver *inframongo.MongoDriver
}

func NewTemplateMongoRepository(logger *slog.Logger, driver *inframongo.MongoDriver) *TemplateMongoRepository {
	return &TemplateMongoRepository{
		logger: logger,
		driver: driver,
	}
}

func (repo *TemplateMongoRepository) FindAll() ([]domain.Template, error) {
	var result []dto.Template

	querier := func(collection *mongo.Collection, ctx *context.Context) error {
		cursor, err := collection.Find(*ctx, bson.M{})
		if err != nil {
			repo.logger.Error("could not extract templates from the database")
			return err
		}

		err = cursor.All(*ctx, &result)

		if err != nil {
			repo.logger.Error("could not decode the templates")
		}

		return err
	}

	err := repo.driver.Process(querier)

	if err != nil {
		repo.logger.Error("could not process the query")
		return nil, err
	}

	templates := make([]domain.Template, len(result))
	for i, t := range result {
		templates[i] = dto.TemplateFromDTO(t)
	}

	return templates, nil
}

func (repo *TemplateMongoRepository) FindByName(name string) (domain.Template, error) {
	var result dto.Template

	querier := func(collection *mongo.Collection, ctx *context.Context) error {
		err := collection.FindOne(*ctx, bson.M{"name": name}).Decode(&result)

		if err == mongo.ErrNoDocuments {
			repo.logger.Error("template not found")
			return domain.ErrTemplateNotFound
		}

		if err != nil {
			repo.logger.Error("could not extract the template from the database")
		}

		return err
	}

	err := repo.driver.Process(querier)

	if err != nil {
		return domain.Template{}, err
	}

	return dto.TemplateFromDTO(result), nil
}

func (repo *TemplateMongoRepository) FindByID(id string) (domain.Template, error) {
	var result dto.Template

	querier := func(collection *mongo.Collection, ctx *context.Context) error {
		err := collection.FindOne(*ctx, bson.M{"_id": id}).Decode(&result)

		if err == mongo.ErrNoDocuments {
			repo.logger.Error("template not found")
			return domain.ErrTemplateNotFound
		}

		if err != nil {
			repo.logger.Error("could not extract the template from the database")
		}

		return err
	}

	err := repo.driver.Process(querier)

	if err != nil {
		return domain.Template{}, err
	}

	return dto.TemplateFromDTO(result), nil
}

func (repo *TemplateMongoRepository) StoreTemplate(template domain.Template) error {
	var dtoTemplate = dto.TemplateToDTO(template)

	querier := func(collection *mongo.Collection, ctx *context.Context) error {
		_, err := collection.InsertOne(*ctx, dtoTemplate)

		if err != nil {
			repo.logger.Error("could not insert the template from the database",
				slog.String("template", fmt.Sprintf("%v", dtoTemplate)),
				slog.String("error", err.Error()),
			)
		}

		return err
	}

	err := repo.driver.Process(querier)

	if err != nil {
		return err
	}

	return nil
}
