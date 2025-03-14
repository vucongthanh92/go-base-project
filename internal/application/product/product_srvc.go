package product

import "context"

type ProductService interface {
	CreateProduct(ctx context.Context) error
	GetProductsByFilter(ctx context.Context) error
	GetProductByID(ctx context.Context) error
	UpdateProductByID(ctx context.Context) error
	DeleteProductByID(ctx context.Context) error
}
