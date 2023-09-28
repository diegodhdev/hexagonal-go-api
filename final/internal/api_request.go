package mooc

import (
	"context"
	"errors"
	"fmt"
	"github.com/diegodhdev/hexagonal-go-api/final/kit/event"
	"github.com/google/uuid"
)

var ErrInvalidApiRequestID = errors.New("invalid Api Request ID")

// ApiRequestID represents the api request unique identifier.
type ApiRequestID struct {
	value string
}

// NewApiRequestID instantiate the VO for ApiRequestID
func NewApiRequestID(value string) (ApiRequestID, error) {

	v, err := uuid.Parse(value)
	if err != nil {
		return ApiRequestID{}, fmt.Errorf("%w: %s", ErrInvalidApiRequestID, value)
	}

	return ApiRequestID{
		value: v.String(),
	}, nil
}

// String type converts the ApiRequestID into string.
func (id ApiRequestID) String() string {
	return id.value
}

// ApiRequest is the data structure that represents a api request.
type ApiRequest struct {
	id ApiRequestID

	events []event.Event
}

// ApiRequestRepository defines the expected behaviour from a course storage.
type ApiRequestRepository interface {
	Save(ctx context.Context, apiRequest ApiRequest) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=CourseRepository

// NewApiRequest creates an api request.
func NewApiRequest(id string) (ApiRequest, error) {
	idVO, err := NewApiRequestID(id)
	if err != nil {
		return ApiRequest{}, err
	}

	apiRequest := ApiRequest{
		id: idVO,
	}
	//course.Record(NewCourseCreatedEvent(idVO.String()))
	return apiRequest, nil
}

// ID returns the api request unique identifier.
func (a ApiRequest) ID() ApiRequestID {
	return a.id
}

// Record records a new domain event.
func (a *ApiRequest) Record(evt event.Event) {
	a.events = append(a.events, evt)
}

// PullEvents returns all the recorded domain events.
func (a ApiRequest) PullEvents() []event.Event {
	evt := a.events
	a.events = []event.Event{}

	return evt
}
