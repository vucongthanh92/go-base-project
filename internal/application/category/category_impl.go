package category

import (
	"context"

	"github.com/vucongthanh92/go-base-project/internal/domain/interfaces"
)

type CategoryImpl struct {
	categoryReadRepo interfaces.CategoryQueryRepoI
}

func NewOtherService(categoryReadRepo interfaces.ProductQueryRepoI) CategoryService {
	return &CategoryImpl{
		categoryReadRepo: categoryReadRepo,
	}
}

func (s *CategoryImpl) CreateCategory(ctx context.Context) error {
	return nil
}

func (s *CategoryImpl) GetCategoryByID(ctx context.Context) error {
	return nil
}

func (s *CategoryImpl) UpdateCategoryByID(ctx context.Context) error {
	return nil
}
