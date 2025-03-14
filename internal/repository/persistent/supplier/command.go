package supplier

import (
	"github.com/jmoiron/sqlx"
	"github.com/vucongthanh92/go-base-project/database"

	"github.com/vucongthanh92/go-base-project/internal/domain/interfaces"
)

type supplierCommandRepository struct {
	readDb *sqlx.DB
}

func NewSupplierCommandRepository(readDb *database.ReadDb) interfaces.SupplierCommandRepoI {
	return &supplierQueryRepository{
		readDb: *readDb,
	}
}
