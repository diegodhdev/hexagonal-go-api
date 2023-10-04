package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	filesystem2 "github.com/diegodhdev/hexagonal-go-api/requests/internal/platform/storage/filesystem"
	"time"

	"github.com/diegodhdev/hexagonal-go-api/requests/internal/platform/bus/inmemory"
	"github.com/diegodhdev/hexagonal-go-api/requests/internal/platform/server"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
)

func Run() error {

	var cfg config
	err := envconfig.Process("MOOC", &cfg)
	if err != nil {
		return err
	}

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	var (
		commandBus       = inmemory.NewCommandBus()
		eventBus         = inmemory.NewEventBus()
		customFilesystem = filesystem2.NewFilesystem("/home/diego/GolandProjects/hexagonal-go-api/requests/storage/")
	)

	// Loading Courses Bootstrapping
	coursesBootstrapping(db, cfg, eventBus, commandBus)

	// Loading Api Requests Fake Story Api Bootstrapping
	apiRequestsFakeStoryApiBootsrapping(db, cfg, eventBus, commandBus, customFilesystem)

	ctx, srv := server.New(context.Background(), cfg.Host, cfg.Port, cfg.ShutdownTimeout, commandBus)
	return srv.Run(ctx)
}

type config struct {
	// Server configuration
	Host            string        `default:"localhost"`
	Port            uint          `default:"8080"`
	ShutdownTimeout time.Duration `default:"10s"`
	// Database configuration
	DbUser    string        `default:"codely"`
	DbPass    string        `default:"codely"`
	DbHost    string        `default:"localhost"`
	DbPort    uint          `default:"3308"`
	DbName    string        `default:"codely"`
	DbTimeout time.Duration `default:"5s"`
}