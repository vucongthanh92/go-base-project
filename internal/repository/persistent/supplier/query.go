package supplier

import (
	"github.com/jmoiron/sqlx"
	"github.com/vucongthanh92/go-base-project/database"
	"github.com/vucongthanh92/go-base-project/redis"

	"github.com/vucongthanh92/go-base-project/internal/domain/interfaces"
)

type supplierQueryRepository struct {
	readDb *sqlx.DB
}

func NewSupplierQueryRepository(readDb *database.ReadDb, redisClient redis.Client) interfaces.SupplierQueryRepoI {
	return &supplierQueryRepository{
		readDb: *readDb,
	}
}
