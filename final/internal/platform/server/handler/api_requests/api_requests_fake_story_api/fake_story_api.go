package api_requests_fake_story_api

import (
	"errors"
	"fmt"
	mooc "github.com/diegodhdev/hexagonal-go-api/final/internal"
	"github.com/diegodhdev/hexagonal-go-api/final/internal/api_requests/fake_story_api"
	"net/http"

	"github.com/diegodhdev/hexagonal-go-api/final/kit/command"
	"github.com/gin-gonic/gin"
)

var data command.DataResponse

// FakeStoryApiHandler returns an HTTP handler for courses creation.
func FakeStoryApiHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req fake_story_api.ApiRequestCommand

		if err := ctx.BindJSON(&req); err != nil {
			fmt.Println(err.Error())
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		fmt.Println(req)

		data, err := commandBus.Dispatch(data, ctx, fake_story_api.NewApiRequestCommand(
			req.ID,
			req.Api,
			req.Mode,
			req.Response,
			req.Requests,
		))

		fmt.Println("Command Bus returning")
		fmt.Println(data)
		fmt.Println(err)

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
