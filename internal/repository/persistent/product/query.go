package product

import (
	"github.com/jmoiron/sqlx"
	"github.com/vucongthanh92/go-base-project/database"
	"github.com/vucongthanh92/go-base-project/redis"

	"github.com/vucongthanh92/go-base-project/internal/domain/interfaces"
)

type productQueryRepository struct {
	readDb *sqlx.DB
}

func NewProductQueryRepository(readDb *database.ReadDb, redisClient redis.Client) interfaces.ProductQueryRepoI {
	return &productQueryRepository{
		readDb: *readDb,
	}
}
