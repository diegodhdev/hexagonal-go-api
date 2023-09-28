package bootstrap

import (
	"database/sql"
	"github.com/diegodhdev/hexagonal-go-api/final/internal/api_requests/get"
	"github.com/diegodhdev/hexagonal-go-api/final/internal/platform/bus/inmemory"
	"github.com/diegodhdev/hexagonal-go-api/final/internal/platform/storage/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func apiRequestsBootsrapping(db *sql.DB, cfg config, eventBus *inmemory.EventBus, commandBus *inmemory.CommandBus) {
	apiRequestRepository := mysql.NewApiRequestRepository(db, cfg.DbTimeout)
	creatingApiRequestService := get.NewApiRequestService(apiRequestRepository, eventBus)
	createApiRequestCommandHandler := get.NewApiRequestCommandHandler(creatingApiRequestService)
	commandBus.Register(get.ApiRequestCommandType, createApiRequestCommandHandler)
}
