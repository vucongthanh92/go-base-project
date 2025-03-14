package category

import (
	"github.com/jmoiron/sqlx"
	"github.com/vucongthanh92/go-base-project/database"

	"github.com/vucongthanh92/go-base-project/internal/domain/interfaces"
)

type categoryCommandRepository struct {
	readDb *sqlx.DB
}

func NewCategoryCommandRepository(readDb *database.ReadDb) interfaces.CategoryCommandRepoI {
	return &categoryQueryRepository{
		readDb: *readDb,
	}
}
