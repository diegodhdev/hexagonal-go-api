package creating

import (
	"context"
	"errors"

	mooc "github.com/diegodhdev/hexagonal-go-api/08-02-build-and-deploy/internal"
	"github.com/diegodhdev/hexagonal-go-api/08-02-build-and-deploy/internal/increasing"
	"github.com/diegodhdev/hexagonal-go-api/08-02-build-and-deploy/kit/event"
)

type IncreaseCoursesCounterOnCourseCreated struct {
	increasingService increasing.CourseCounterService
}

func NewIncreaseCoursesCounterOnCourseCreated(increaserService increasing.CourseCounterService) IncreaseCoursesCounterOnCourseCreated {
	return IncreaseCoursesCounterOnCourseCreated{
		increasingService: increaserService,
	}
}

func (e IncreaseCoursesCounterOnCourseCreated) Handle(_ context.Context, evt event.Event) error {
	courseCreatedEvt, ok := evt.(mooc.CourseCreatedEvent)
	if !ok {
		return errors.New("unexpected event")
	}

	return e.increasingService.Increase(courseCreatedEvt.ID())
}
