package product

import (
	"context"

	"github.com/vucongthanh92/go-base-project/database"
	"github.com/vucongthanh92/go-base-utils/tracing"
	"gorm.io/gorm"

	httpcommon "github.com/vucongthanh92/go-base-project/helper/http_common"
	"github.com/vucongthanh92/go-base-project/internal/domain/entities"
	"github.com/vucongthanh92/go-base-project/internal/domain/interfaces"
	"github.com/vucongthanh92/go-base-project/internal/domain/models"
)

type productQueryRepository struct {
	readDb *gorm.DB
}

func NewProductQueryRepository(readDb *database.GormReadDb) interfaces.ProductQueryRepoI {
	return &productQueryRepository{
		readDb: *readDb,
	}
}

func (repo *productQueryRepository) GetProductByFilter(ctx context.Context, filter models.ProductListFilter) (
	response []entities.Product, totalRows int64, errRes httpcommon.ErrorDTO) {

	ctx, span := tracing.StartSpanFromContext(ctx, "GetProductByFilter")
	defer span.End()

	err := repo.readDb.WithContext(ctx).Model(&entities.Product{}).Select("*").Count(&totalRows).
		Limit(filter.Limit).Offset(filter.Offset).Find(&response).Error
	if err != nil {
		errRes.IsSystemError = true
		errRes.Error = err
		return response, totalRows, errRes
	}

	return response, totalRows, errRes
}
