package server

import (
	"fmt"
	"log"

	mooc "github.com/diegodhdev/hexagonal-go-api/03-02-repository-test/internal"
	"github.com/diegodhdev/hexagonal-go-api/03-02-repository-test/internal/platform/server/handler/courses"
	"github.com/diegodhdev/hexagonal-go-api/03-02-repository-test/internal/platform/server/handler/health"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	// deps
	courseRepository mooc.CourseRepository
}

func New(host string, port uint, courseRepository mooc.CourseRepository) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		courseRepository: courseRepository,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/courses", courses.CreateHandler(s.courseRepository))
}
