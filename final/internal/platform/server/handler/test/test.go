package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CheckHandler returns an HTTP handler to perform health checks.
func TestHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.String(http.StatusOK, "everything is ok!")
	}
}
