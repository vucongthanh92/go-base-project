//go:build wireinject
// +build wireinject

package internal

import (
	categoryService "github.com/vucongthanh92/go-base-project/internal/application/category"
	"github.com/vucongthanh92/go-base-project/internal/application/cronjob"
	productService "github.com/vucongthanh92/go-base-project/internal/application/product"
	supplierService "github.com/vucongthanh92/go-base-project/internal/application/supplier"

	categoryRepo "github.com/vucongthanh92/go-base-project/internal/repository/persistent/category"
	productRepo "github.com/vucongthanh92/go-base-project/internal/repository/persistent/product"
	supplierRepo "github.com/vucongthanh92/go-base-project/internal/repository/persistent/supplier"

	grpcserver "github.com/vucongthanh92/go-base-project/internal/api/grpc"
	v1 "github.com/vucongthanh92/go-base-project/internal/api/http/v1"

	"github.com/vucongthanh92/go-base-project/config"
	"github.com/vucongthanh92/go-base-project/database"
	"github.com/vucongthanh92/go-base-project/internal/api"
	"github.com/vucongthanh92/go-base-project/internal/api/cron"
	"github.com/vucongthanh92/go-base-project/internal/api/http"
	"github.com/vucongthanh92/go-base-project/redis"

	"github.com/google/wire"
)

var container = wire.NewSet(
	api.NewApiContainer,
)

var apiSet = wire.NewSet(
	cron.NewServer,
	grpcserver.NewServer,
	http.NewServer,
)

var handlerSet = wire.NewSet(
	v1.NewProductHandler,
	v1.NewCategoryHandler,
	v1.NewSupplierHandler,
)

var serviceSet = wire.NewSet(
	cronjob.NewCronJobService,
	productService.NewProductService,
	categoryService.NewOtherService,
	supplierService.NewSupplierService,
)

var repoSet = wire.NewSet(
	productRepo.NewProductCommandRepository,
	productRepo.NewProductQueryRepository,
	categoryRepo.NewCategoryCommandRepository,
	categoryRepo.NewCategoryQueryRepository,
	supplierRepo.NewSupplierCommandRepository,
	supplierRepo.NewSupplierQueryRepository,
)

func InitializeContainer(
	appCfg *config.AppConfig,
	readDb *database.ReadDb,
	writeDb *database.WriteDb,
	redisClient redis.Client,
) *api.ApiContainer {
	wire.Build(repoSet, serviceSet, handlerSet, apiSet, container)
	return &api.ApiContainer{}
}
