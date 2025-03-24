package category

import (
	"context"
	"time"

	"github.com/vucongthanh92/go-base-project/database"
	"github.com/vucongthanh92/go-base-utils/tracing"
	"gorm.io/gorm"

	"github.com/vucongthanh92/go-base-project/helper/constants"
	errHandler "github.com/vucongthanh92/go-base-project/helper/error_handler"
	"github.com/vucongthanh92/go-base-project/internal/domain/entities"
	"github.com/vucongthanh92/go-base-project/internal/domain/interfaces"
	"github.com/vucongthanh92/go-base-project/internal/domain/models"
)

type categoryCommandRepository struct {
	writeDB *gorm.DB
}

func NewCategoryCommandRepository(writeDB *database.GormWriteDb) interfaces.CategoryCommandRepoI {
	return &categoryCommandRepository{
		writeDB: *writeDB,
	}
}

func (repo *categoryCommandRepository) InsertCategory(ctx context.Context, entity entities.Category) (
	entities.Category, *errHandler.ErrorBuilder) {

	ctx, span := tracing.StartSpanFromContext(ctx, "InsertCategory")
	defer span.End()

	err := repo.writeDB.WithContext(ctx).Model(entities.Category{}).
		Create(&entity).Error

	if err != nil {
		resErr := errHandler.InitErrorBuilder().SetIsSystemError(true).SetLogError(err).SetError(models.ErrorDTO{
			Message: err.Error(),
			Code:    constants.SystemError,
		})
		return entity, resErr
	}

	return entity, nil
}

func (repo *categoryCommandRepository) UpdateCategory(ctx context.Context, entity entities.Category) (
	entities.Category, *errHandler.ErrorBuilder) {
	ctx, span := tracing.StartSpanFromContext(ctx, "UpdateCategory")
	defer span.End()

	err := repo.writeDB.WithContext(ctx).Model(entities.Category{}).
		Where("id = ?", entity.ID).Where("updated_at = ?", entity.UpdatedAt).
		Update("name", entity.Name).Error

	if err != nil {
		resErr := errHandler.InitErrorBuilder().SetIsSystemError(true).SetLogError(err).SetError(models.ErrorDTO{
			Message: err.Error(),
			Code:    constants.SystemError,
		})
		return entity, resErr
	}

	return entity, nil
}

func (repo *categoryCommandRepository) SoftDeleteCategoryByID(ctx context.Context, id uint64, updatedAt time.Time) *errHandler.ErrorBuilder {
	ctx, span := tracing.StartSpanFromContext(ctx, "UpdateCategory")
	defer span.End()

	err := repo.writeDB.WithContext(ctx).Model(entities.Category{}).
		Where("id = ?", id).Where("updated_at = ?", updatedAt).
		UpdateColumn("deleted_at = ", time.Now()).Error

	if err != nil {
		resErr := errHandler.InitErrorBuilder().SetIsSystemError(true).SetLogError(err).SetError(models.ErrorDTO{
			Message: err.Error(),
			Code:    constants.SystemError,
		})
		return resErr
	}

	return nil
}
