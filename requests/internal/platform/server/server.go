package server

import (
	"context"
	"fmt"
	"github.com/diegodhdev/hexagonal-go-api/requests/internal/platform/server/handler/api_requests/api_requests_fake_story_api"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/diegodhdev/hexagonal-go-api/requests/internal/platform/server/handler/courses"
	"github.com/diegodhdev/hexagonal-go-api/requests/internal/platform/server/handler/health"
	"github.com/diegodhdev/hexagonal-go-api/requests/internal/platform/server/middleware/logging"
	"github.com/diegodhdev/hexagonal-go-api/requests/internal/platform/server/middleware/recovery"
	"github.com/diegodhdev/hexagonal-go-api/requests/kit/command"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	shutdownTimeout time.Duration

	// deps
	commandBus command.Bus
}

func New(ctx context.Context, host string, port uint, shutdownTimeout time.Duration, commandBus command.Bus) (context.Context, Server) {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		shutdownTimeout: shutdownTimeout,

		commandBus: commandBus,
	}

	srv.registerRoutes()
	return serverContext(ctx), srv
}

func (s *Server) registerRoutes() {

	s.engine.Use(recovery.Middleware(), logging.Middleware())

	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/courses", courses.CreateHandler(s.commandBus))
	s.engine.POST("/requests/fake-story-api", api_requests_fake_story_api.FakeStoryApiHandler(s.commandBus))
}

func (s *Server) Run(ctx context.Context) error {
	log.Println("Server running on", s.httpAddr)

	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shut down", err)
		}
	}()

	<-ctx.Done()
	ctxShutDown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return srv.Shutdown(ctxShutDown)
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}