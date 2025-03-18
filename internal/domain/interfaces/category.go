package interfaces

import (
	"context"

	"github.com/vucongthanh92/go-base-project/internal/domain/entities"
)

type CategoryQueryRepoI interface {
}

type CategoryCommandRepoI interface {
	InsertCategory(ctx context.Context, entity entities.Category) (err error)
}
