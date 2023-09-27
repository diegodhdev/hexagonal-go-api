package bootstrap

import (
	"database/sql"
	"fmt"

	"github.com/diegodhdev/hexagonal-go-api/04-02-application-service-test/internal/creating"
	"github.com/diegodhdev/hexagonal-go-api/04-02-application-service-test/internal/platform/server"
	"github.com/diegodhdev/hexagonal-go-api/04-02-application-service-test/internal/platform/storage/mysql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host = "localhost"
	port = 8080

	dbUser = "codely"
	dbPass = "codely"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "codely"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	courseRepository := mysql.NewCourseRepository(db)

	creatingCourseService := creating.NewCourseService(courseRepository)

	srv := server.New(host, port, creatingCourseService)
	return srv.Run()
}
