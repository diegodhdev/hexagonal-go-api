package command

import (
	"context"
)

type DataResponse struct {
	data string
}

func NewDataResponse(value string) DataResponse {

	return DataResponse{
		data: value,
	}
}

// Bus defines the expected behaviour from a command bus.
type Bus interface {
	// Dispatch is the method used to dispatch new commands.
	Dispatch(DataResponse, context.Context, Command) (DataResponse, error)
	// Register is the method used to register a new command handler.
	Register(Type, Handler)
}

//go:generate mockery --case=snake --outpkg=commandmocks --output=commandmocks --name=Bus

// Type represents an application command type.
type Type string

// Command represents an application command.
type Command interface {
	Type() Type
}

// Handler defines the expected behaviour from a command handler.
type Handler interface {
	Handle(DataResponse, context.Context, Command) (DataResponse, error)
}
