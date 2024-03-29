package infra

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/wizzardich/autoreplace-reincarnated/domain"
	"github.com/wizzardich/autoreplace-reincarnated/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (repo *TemplateMongoRepository) FindByID(id domain.IDString) (domain.Template, error) {
	var result dto.Template

	querier := func(collection *mongo.Collection, ctx *context.Context) error {
		objectID, err := primitive.ObjectIDFromHex(id)
		filter := bson.M{"_id": objectID}
		if err != nil {
			repo.logger.Warn("could not convert the id to an ObjectID",
				slog.String("id", id),
				slog.String("error", err.Error()),
			)
			// not that I especially like the fallback here...
			filter = bson.M{"_id": id}
		}

		err = collection.FindOne(*ctx, filter).Decode(&result)

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

func (repo *TemplateMongoRepository) StoreTemplate(template domain.Template) (*domain.IDString, error) {
	var dtoTemplate = dto.TemplateToDTO(template)
	var insertedID domain.IDString

	querier := func(collection *mongo.Collection, ctx *context.Context) error {
		result, err := collection.InsertOne(*ctx, dtoTemplate)

		if err != nil {
			repo.logger.Error("could not insert the template from the database",
				slog.String("template", fmt.Sprintf("%v", dtoTemplate)),
				slog.String("error", err.Error()),
			)
			return err
		}

		insertedID = result.InsertedID.(primitive.ObjectID).Hex()

		return nil
	}

	err := repo.driver.Process(querier) // TODO: at least return the ID of the inserted template

	if err != nil {
		return nil, err
	}

	return &insertedID, nil
}

func (repo *TemplateMongoRepository) UpdateTemplate(id domain.IDString, template domain.Template) error {
	var dtoTemplate = dto.TemplateToDTO(template)

	querier := func(collection *mongo.Collection, ctx *context.Context) error {
		var objectID interface{} // this is horribly untypesafe, but this is all mongo's fault!
		// TODO: I will do a wrapper later
		objectID, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			repo.logger.Warn("could not convert the id to an ObjectID",
				slog.String("id", id),
				slog.String("error", err.Error()),
			)
			objectID = id
		}

		payload := bson.M{}
		if template.Name != "" {
			payload["name"] = template.Name
		}
		if len(template.Replacements) > 0 {
			payload["reps"] = dtoTemplate.Replacements
		}

		_, err = collection.UpdateByID(*ctx, objectID, bson.M{"$set": payload})

		if err != nil {
			repo.logger.Error("could not update the template from the database",
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

func (repo *TemplateMongoRepository) DeleteTemplate(id domain.IDString) error {
	querier := func(collection *mongo.Collection, ctx *context.Context) error {
		objectID, err := primitive.ObjectIDFromHex(id)

		filter := bson.M{"_id": objectID}
		if err != nil {
			repo.logger.Warn("could not convert the id to an ObjectID",
				slog.String("id", id),
				slog.String("error", err.Error()),
			)
			// not that I especially like the fallback here...
			filter = bson.M{"_id": id}
		}

		deletionResult, err := collection.DeleteOne(*ctx, filter)

		if deletionResult.DeletedCount == 0 {
			repo.logger.Error("template not found")
			return domain.ErrTemplateNotFound
		}

		if err != nil {
			repo.logger.Error("could not delete the template from the database",
				slog.String("id", id),
				slog.String("error", err.Error()),
			)
			return err
		}

		return nil
	}

	err := repo.driver.Process(querier)

	if err != nil {
		return err
	}

	return nil
}
