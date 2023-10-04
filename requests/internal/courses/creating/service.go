package creating

import (
	"context"
	mooc "github.com/diegodhdev/hexagonal-go-api/requests/internal"
	"github.com/diegodhdev/hexagonal-go-api/requests/kit/command"
	"github.com/diegodhdev/hexagonal-go-api/requests/kit/event"
)

// CourseService is the default CourseService interface
// implementation returned by creating.NewCourseService.
type CourseService struct {
	courseRepository mooc.CourseRepository
	eventBus         event.Bus
}

// NewCourseService returns the default Service interface implementation.
func NewCourseService(courseRepository mooc.CourseRepository, eventBus event.Bus) CourseService {
	return CourseService{
		courseRepository: courseRepository,
		eventBus:         eventBus,
	}
}

// CreateCourse implements the creating.CourseService interface.
func (s CourseService) CreateCourse(ctx context.Context, id, name, duration string) (any, error) {
	course, err := mooc.NewCourse(id, name, duration)
	if err != nil {
		return command.NewEmptyDataResponse(), err
	}

	if err := s.courseRepository.Save(ctx, course); err != nil {
		return command.NewEmptyDataResponse(), err
	}

	return command.NewEmptyDataResponse(), s.eventBus.Publish(ctx, course.PullEvents())
}
