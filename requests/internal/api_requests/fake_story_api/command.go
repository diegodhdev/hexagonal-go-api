package fake_story_api

import (
	"context"
	"errors"
	mooc "github.com/diegodhdev/hexagonal-go-api/requests/internal"
	"github.com/diegodhdev/hexagonal-go-api/requests/kit/command"
)

const ApiRequestCommandType command.Type = "command.fake_story_api.apiRequest"

var data command.DataResponse

type ApiRequestCommand struct {
	// json tag to de-serialize json body
	ID       string         `json:"ID"`
	Api      string         `json:"api"`
	Mode     string         `json:"mode"`
	Response string         `json:"response"`
	Requests []mooc.Request `json:"requests"`
}

// RequestCommand is the command dispatched to create a new apiRequest.
type RequestCommand struct {
	id string
}

// NewApiRequestCommand creates a new ApiRequestCommand.
func NewApiRequestCommand(id string, api string, mode string, response string, requests []mooc.Request) ApiRequestCommand {
	return ApiRequestCommand{
		ID:       id,
		Api:      api,
		Mode:     mode,
		Response: response,
		Requests: requests,
	}
}

func (c ApiRequestCommand) Type() command.Type {
	return ApiRequestCommandType
}

// ApiRequestCommandHandler is the command handler
// responsible for creating courses.
type ApiRequestCommandHandler struct {
	service ApiRequestService
}

// NewApiRequestCommandHandler initializes a new ApiRequestCommandHandler.
func NewApiRequestCommandHandler(service ApiRequestService) ApiRequestCommandHandler {
	return ApiRequestCommandHandler{
		service: service,
	}
}

// Handle implements the command.Handler interface.
func (h ApiRequestCommandHandler) Handle(ctx context.Context, cmd command.Command) (any, error) {
	getApiRequestCmd, ok := cmd.(ApiRequestCommand)

	if !ok {
		return command.NewEmptyDataResponse(), errors.New("unexpected command")
	}

	return h.service.FakeStoryApiApiRequest(
		ctx,
		getApiRequestCmd.ID,
		getApiRequestCmd.Api,
		getApiRequestCmd.Mode,
		getApiRequestCmd.Response,
		getApiRequestCmd.Requests,
	)
}
