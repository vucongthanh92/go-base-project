package category

import "context"

type CategoryService interface {
	CreateCategory(ctx context.Context) error
	GetCategoryByID(ctx context.Context) error
	UpdateCategoryByID(ctx context.Context) error
}
