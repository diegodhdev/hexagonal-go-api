package bootstrap

import (
	"database/sql"
	"github.com/diegodhdev/hexagonal-go-api/final/internal/api_requests/fake_story_api"
	"github.com/diegodhdev/hexagonal-go-api/final/internal/platform/bus/inmemory"
	"github.com/diegodhdev/hexagonal-go-api/final/internal/platform/storage/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func apiRequestsFakeStoryApiBootsrapping(db *sql.DB, cfg config, eventBus *inmemory.EventBus, commandBus *inmemory.CommandBus) {
	apiRequestRepository := mysql.NewApiRequestRepository(db, cfg.DbTimeout)
	creatingApiRequestFakeStoryApiService := fake_story_api.NewApiRequestService(apiRequestRepository, eventBus)
	createApiRequestFakeStoryApiCommandHandler := fake_story_api.NewApiRequestCommandHandler(creatingApiRequestFakeStoryApiService)
	commandBus.Register(fake_story_api.ApiRequestCommandType, createApiRequestFakeStoryApiCommandHandler)
}
