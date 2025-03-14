package startup

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/vucongthanh92/go-base-utils/tracing"

	"github.com/vucongthanh92/go-base-project/config"
	"github.com/vucongthanh92/go-base-project/database"

	"github.com/vucongthanh92/go-base-project/helper/healthcheck"
	"github.com/vucongthanh92/go-base-project/internal"
	"github.com/vucongthanh92/go-base-project/internal/api"
	"github.com/vucongthanh92/go-base-project/redis"
	"go.uber.org/zap"

	"github.com/gammazero/workerpool"
	"github.com/vucongthanh92/go-base-utils/command"
	"github.com/vucongthanh92/go-base-utils/logger"

	"github.com/vucongthanh92/go-base-utils/localization"
)

func runServer(
	ctx context.Context,
	container *api.ApiContainer,
	readDb database.ReadDb,
	writeDb database.WriteDb,
) {
	wp := workerpool.New(5)

	// Run healthcheck
	wp.Submit(healthcheck.RunHealthCheck(ctx, cfg, readDb, writeDb))

	// Run Grpc server
	wp.Submit(container.GrpcServer.Run)

	// Run Http server
	wp.Submit(container.HttpServer.Run)

	// Run CronJob server
	wp.Submit(container.CronServer.Run)

	wp.StopWait()
}

func registerDependencies(ctx context.Context) (*api.ApiContainer, database.ReadDb, database.WriteDb) {
	redisClient := redis.Open(cfg.Redis)

	// Open database connection
	readDb, writeDb := database.Open(cfg.Database)

	return internal.InitializeContainer(
		cfg,
		&readDb,
		&writeDb,
		redisClient,
	), readDb, writeDb
}

var cfg *config.AppConfig

func Execute() {

	// Init AppConfig
	cfg = &config.AppConfig{}

	// Init commands
	command.UseCommands(
		command.WithStartCommand(start, cfg),
	)
}

func start() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	tracing.UseOpenTelemetry(tracing.Config(*cfg.Tracing))

	// Register dependencies
	container, readDb, writeDb := registerDependencies(ctx)
	// Init resources for localization
	err := localization.InitResources(cfg.Http.Resources)

	if err != nil {
		logger.Fatal("Fail to init resources", zap.Error(err))
	}

	// Init tracing
	// Init validation

	// validation.UseValidation(container.ValidationEngine.GetValidations()...)

	// Run server
	runServer(ctx, container, readDb, writeDb)
}
