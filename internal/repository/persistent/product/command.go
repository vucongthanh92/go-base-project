package product

import (
	"github.com/jmoiron/sqlx"
	"github.com/vucongthanh92/go-base-project/database"

	"github.com/vucongthanh92/go-base-project/internal/domain/interfaces"
)

type productCommandRepository struct {
	writeDb *sqlx.DB
}

func NewProductCommandRepository(writeDb *database.ReadDb) interfaces.ProductCommandRepoI {
	return &productCommandRepository{
		writeDb: *writeDb,
	}
}
