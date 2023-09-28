package bootstrap

import (
	"database/sql"
	mooc "github.com/diegodhdev/hexagonal-go-api/final/internal"
	"github.com/diegodhdev/hexagonal-go-api/final/internal/courses/creating"
	"github.com/diegodhdev/hexagonal-go-api/final/internal/courses/increasing"
	"github.com/diegodhdev/hexagonal-go-api/final/internal/platform/bus/inmemory"
	"github.com/diegodhdev/hexagonal-go-api/final/internal/platform/storage/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func coursesBootstrapping(db *sql.DB, cfg config, eventBus *inmemory.EventBus, commandBus *inmemory.CommandBus) {
	courseRepository := mysql.NewCourseRepository(db, cfg.DbTimeout)
	creatingCourseService := creating.NewCourseService(courseRepository, eventBus)
	increasingCourseService := increasing.NewCourseCounterService()
	createCourseCommandHandler := creating.NewCourseCommandHandler(creatingCourseService)
	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)
	eventBus.Subscribe(
		mooc.CourseCreatedEventType,
		creating.NewIncreaseCoursesCounterOnCourseCreated(increasingCourseService),
	)
}
