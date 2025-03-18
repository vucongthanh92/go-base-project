package product

import (
	"context"

	httpcommon "github.com/vucongthanh92/go-base-project/helper/http_common"
	"github.com/vucongthanh92/go-base-project/internal/domain/entities"
	"github.com/vucongthanh92/go-base-project/internal/domain/interfaces"
	"github.com/vucongthanh92/go-base-project/internal/domain/models"
	"github.com/vucongthanh92/go-base-utils/tracing"
)

type ProductImpl struct {
	productReadRepo interfaces.ProductQueryRepoI
}

func NewProductService(productReadRepo interfaces.ProductQueryRepoI) ProductService {
	return &ProductImpl{
		productReadRepo: productReadRepo,
	}
}

func (s *ProductImpl) CreateProduct(ctx context.Context) error {
	return nil
}

func (s *ProductImpl) GetProductsByFilter(ctx context.Context, req models.ProductListFilter) (
	response []entities.Product, totalRows int64, errRes httpcommon.ErrorDTO) {

	ctx, span := tracing.StartSpanFromContext(ctx, "GetProductsByFilter")
	defer func() {
		span.End()
	}()

	response, totalRows, errRes = s.productReadRepo.GetProductByFilter(ctx, req)
	return response, totalRows, errRes
}

func (s *ProductImpl) GetProductByID(ctx context.Context) error {
	return nil
}

func (s *ProductImpl) UpdateProductByID(ctx context.Context) error {
	return nil
}

func (s *ProductImpl) DeleteProductByID(ctx context.Context) error {
	return nil
}
