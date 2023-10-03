package api_requests_get

import (
	"errors"
	"github.com/diegodhdev/hexagonal-go-api/final/internal/api_requests/get"
	"net/http"

	mooc "github.com/diegodhdev/hexagonal-go-api/final/internal"
	"github.com/diegodhdev/hexagonal-go-api/final/kit/command"
	"github.com/gin-gonic/gin"
)

type createRequest struct {
	ID string `json:"id" binding:"required"`
}

var data command.DataResponse

// GetHandler returns an HTTP handler for courses creation.
func GetHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req createRequest

		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		_, err := commandBus.Dispatch(ctx, get.NewApiRequestCommand(
			req.ID,
			"cacac",
		))

		if err != nil {
			switch {
			case errors.Is(err, mooc.ErrInvalidApiRequestID),
				errors.Is(err, mooc.ErrEmptyCourseName), errors.Is(err, mooc.ErrInvalidCourseID):
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		ctx.Status(http.StatusCreated)
	}
}
