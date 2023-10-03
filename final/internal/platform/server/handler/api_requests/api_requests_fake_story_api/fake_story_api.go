package api_requests_fake_story_api

import (
	"encoding/json"
	"errors"
	mooc "github.com/diegodhdev/hexagonal-go-api/final/internal"
	"github.com/diegodhdev/hexagonal-go-api/final/internal/api_requests/fake_story_api"
	"github.com/diegodhdev/hexagonal-go-api/final/kit/command"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

// FakeStoryApiHandler returns an HTTP handler for courses creation.
func FakeStoryApiHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req fake_story_api.ApiRequestCommand

		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		data, err := commandBus.Dispatch(ctx, fake_story_api.NewApiRequestCommand(
			req.ID,
			req.Api,
			req.Mode,
			req.Response,
			req.Requests,
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

		p := fake_story_api.NewProductFakeStoryApiDataResponse()

		if reflect.TypeOf(data) != reflect.TypeOf(p) {
			ctx.JSON(http.StatusBadRequest, "Response unexpected")
			return
		}

		convertedData, err := json.Marshal(data)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Error marshalling response")
			return
		}

		err = json.Unmarshal(convertedData, &p)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Error unmarshalling response")
			return
		}

		ctx.JSON(http.StatusOK, p.Data)
	}
}
