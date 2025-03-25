package category

import (
	"context"
	"net/http"

	"github.com/vucongthanh92/go-base-project/database"
	"github.com/vucongthanh92/go-base-project/helper/constants"
	errHandler "github.com/vucongthanh92/go-base-project/helper/error_handler"
	"github.com/vucongthanh92/go-base-utils/tracing"
	"gorm.io/gorm"

	"github.com/vucongthanh92/go-base-project/internal/domain/entities"
	"github.com/vucongthanh92/go-base-project/internal/domain/interfaces"
	"github.com/vucongthanh92/go-base-project/internal/domain/models"
)

type categoryQueryRepository struct {
	readDb *gorm.DB
}

func NewCategoryQueryRepository(readDb *database.GormReadDb) interfaces.CategoryQueryRepoI {
	return &categoryQueryRepository{
		readDb: *readDb,
	}
}

func (repo *categoryQueryRepository) GetCategoryByID(ctx context.Context, id uint64) (res entities.Category, errRes *errHandler.ErrorBuilder) {
	ctx, span := tracing.StartSpanFromContext(ctx, "GetCategoryByID")
	defer span.End()

	err := repo.readDb.WithContext(ctx).Model(&entities.Category{}).
		Select("id, name, created_at, updated_at").
		Where("id = ?", id).Where("deleted_at is null").
		Take(&res).Error
	if err != nil {
		resErr := errHandler.InitErrorBuilder().
			SetIsSystemError(false).
			SetLogError(err).
			SetStatus(http.StatusNotFound).
			SetError(models.ErrorDTO{
				Message: "Category not found",
				Code:    constants.RECORD_NOT_EXIST,
			})
		return res, resErr
	}

	return res, errRes
}

func (repo *categoryQueryRepository) GetCategoryList(ctx context.Context) (res []entities.Category, errRes *errHandler.ErrorBuilder) {
	ctx, span := tracing.StartSpanFromContext(ctx, "GetCategoryList")
	defer span.End()

	err := repo.readDb.WithContext(ctx).Model(&entities.Category{}).
		Select("id, name, created_at, updated_at").
		Find(&res).Error
	if err != nil {
		resErr := errHandler.InitErrorBuilder().SetIsSystemError(true).SetLogError(err).SetError(models.ErrorDTO{
			Message: err.Error(),
			Code:    constants.SYSTEM_ERROR,
		})
		return res, resErr
	}

	return res, nil
}
