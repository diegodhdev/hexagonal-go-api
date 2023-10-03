package get

import (
	"context"
	"errors"
	"github.com/diegodhdev/hexagonal-go-api/final/kit/command"
)

const ApiRequestCommandType command.Type = "command.get.apiRequest"

// ApiRequestCommand is the command dispatched to create a new apiRequest.
type ApiRequestCommand struct {
	id  string
	tag string
}

// RequestCommand is the command dispatched to create a new apiRequest.
type RequestCommand struct {
	id string
}

// NewApiRequestCommand creates a new ApiRequestCommand.
func NewApiRequestCommand(id string, tag string) ApiRequestCommand {
	return ApiRequestCommand{
		id:  id,
		tag: tag,
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

	return h.service.GetApiRequest(
		ctx,
		getApiRequestCmd.id,
		getApiRequestCmd.tag,
	)
}
