package bootstrap

import (
	"github.com/diegodhdev/hexagonal-go-api/02-01-post-course-endpoint/internal/platform/server"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host = "localhost"
	port = 8080
)

func Run() error {
	srv := server.New(host, port)
	return srv.Run()
}
