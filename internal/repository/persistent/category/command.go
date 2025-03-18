package category

import (
	"context"

	"github.com/vucongthanh92/go-base-project/database"
	"github.com/vucongthanh92/go-base-utils/tracing"
	"gorm.io/gorm"

	"github.com/vucongthanh92/go-base-project/internal/domain/entities"
	"github.com/vucongthanh92/go-base-project/internal/domain/interfaces"
)

type categoryCommandRepository struct {
	writeDB *gorm.DB
}

func NewCategoryCommandRepository(writeDB *database.GormWriteDb) interfaces.CategoryCommandRepoI {
	return &categoryCommandRepository{
		writeDB: *writeDB,
	}
}

func (repo *categoryCommandRepository) InsertCategory(ctx context.Context, entity entities.Category) (err error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "InsertCategory")
	defer span.End()

	err = repo.writeDB.WithContext(ctx).
		Create(&entity).Error

	return nil
}
