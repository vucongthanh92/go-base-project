package category

import (
	"context"

	"github.com/vucongthanh92/go-base-project/internal/domain/entities"
	"github.com/vucongthanh92/go-base-project/internal/domain/interfaces"
	"github.com/vucongthanh92/go-base-project/internal/domain/models"
)

type CategoryImpl struct {
	categoryReadRepo  interfaces.CategoryQueryRepoI
	categoryWriteRepo interfaces.CategoryCommandRepoI
}

func NewCategoryService(
	categoryReadRepo interfaces.ProductQueryRepoI,
	categoryWriteRepo interfaces.CategoryCommandRepoI,
) CategoryService {
	return &CategoryImpl{
		categoryReadRepo:  categoryReadRepo,
		categoryWriteRepo: categoryWriteRepo,
	}
}

func (s *CategoryImpl) CreateCategory(ctx context.Context, req models.CreateCategoryReq) error {

	categoryEntity := entities.Category{
		Name: req.Name,
	}

	err := s.categoryWriteRepo.InsertCategory(ctx, categoryEntity)
	if err != nil {
		return err
	}

	return nil
}

func (s *CategoryImpl) GetCategoryByID(ctx context.Context) error {
	return nil
}

func (s *CategoryImpl) UpdateCategoryByID(ctx context.Context) error {
	return nil
}
