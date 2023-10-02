package creating

import (
	"context"
	"errors"

	"github.com/diegodhdev/hexagonal-go-api/final/kit/command"
)

const CourseCommandType command.Type = "command.creating.course"

// CourseCommand is the command dispatched to create a new course.
type CourseCommand struct {
	id       string
	name     string
	duration string
}

// RequestCommand is the command dispatched to create a new course.
type RequestCommand struct {
	name     string
	duration string
}

// NewCourseCommand creates a new CourseCommand.
func NewCourseCommand(id, name, duration string) CourseCommand {
	return CourseCommand{
		id:       id,
		name:     name,
		duration: duration,
	}
}

func (c CourseCommand) Type() command.Type {
	return CourseCommandType
}

// CourseCommandHandler is the command handler
// responsible for creating courses.
type CourseCommandHandler struct {
	service CourseService
}

// NewCourseCommandHandler initializes a new CourseCommandHandler.
func NewCourseCommandHandler(service CourseService) CourseCommandHandler {
	return CourseCommandHandler{
		service: service,
	}
}

// Handle implements the command.Handler interface.
func (h CourseCommandHandler) Handle(data command.DataResponse, ctx context.Context, cmd command.Command) (command.DataResponse, error) {
	createCourseCmd, ok := cmd.(CourseCommand)
	if !ok {
		return data, errors.New("unexpected command")
	}

	return h.service.CreateCourse(
		data,
		ctx,
		createCourseCmd.id,
		createCourseCmd.name,
		createCourseCmd.duration,
	)
}
