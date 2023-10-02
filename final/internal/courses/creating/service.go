package creating

import (
	"context"
	"github.com/diegodhdev/hexagonal-go-api/final/kit/command"

	mooc "github.com/diegodhdev/hexagonal-go-api/final/internal"
	"github.com/diegodhdev/hexagonal-go-api/final/kit/event"
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
func (s CourseService) CreateCourse(data command.DataResponse, ctx context.Context, id, name, duration string) (command.DataResponse, error) {
	course, err := mooc.NewCourse(id, name, duration)
	if err != nil {
		return data, err
	}

	if err := s.courseRepository.Save(ctx, course); err != nil {
		return data, err
	}

	return data, s.eventBus.Publish(ctx, course.PullEvents())
}
