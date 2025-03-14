package product

import (
	"context"

	"github.com/vucongthanh92/go-base-project/internal/domain/interfaces"
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

func (s *ProductImpl) GetProductsByFilter(ctx context.Context) error {
	return nil
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
