package inmemory

import (
	"context"
	"fmt"

	"github.com/diegodhdev/hexagonal-go-api/final/kit/command"
)

// CommandBus is an in-memory implementation of the command.Bus.
type CommandBus struct {
	handlers map[command.Type]command.Handler
}

// NewCommandBus initializes a new instance of CommandBus.
func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[command.Type]command.Handler),
	}
}

// Dispatch implements the command.Bus interface.
func (b *CommandBus) Dispatch(data command.DataResponse, ctx context.Context, cmd command.Command) (command.DataResponse, error) {

	handler, ok := b.handlers[cmd.Type()]
	if !ok {
		return data, nil
	}

	return handler.Handle(data, ctx, cmd)
}

// Register implements the command.Bus interface.
func (b *CommandBus) Register(cmdType command.Type, handler command.Handler) {

	fmt.Println(cmdType)
	fmt.Println(handler)
	b.handlers[cmdType] = handler
}
