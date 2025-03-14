package category

import (
	"github.com/jmoiron/sqlx"
	"github.com/vucongthanh92/go-base-project/database"
	"github.com/vucongthanh92/go-base-project/redis"

	"github.com/vucongthanh92/go-base-project/internal/domain/interfaces"
)

type categoryQueryRepository struct {
	readDb *sqlx.DB
}

func NewCategoryQueryRepository(readDb *database.ReadDb, redisClient redis.Client) interfaces.CategoryQueryRepoI {
	return &categoryQueryRepository{
		readDb: *readDb,
	}
}
