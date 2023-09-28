package get

import (
	"context"
	"fmt"
	mooc "github.com/diegodhdev/hexagonal-go-api/final/internal"
	"github.com/diegodhdev/hexagonal-go-api/final/kit/event"
)

// ApiRequestService is the default ApiRequestService interface
// implementation returned by creating.NewCourseService.
type ApiRequestService struct {
	apiRequestRepository mooc.ApiRequestRepository
	eventBus             event.Bus
}

// NewApiRequestService returns the default Service interface implementation.
func NewApiRequestService(apiRequestRepository mooc.ApiRequestRepository, eventBus event.Bus) ApiRequestService {
	return ApiRequestService{
		apiRequestRepository: apiRequestRepository,
		eventBus:             eventBus,
	}
}

// GetApiRequest implements the get.GetApiRequest interface.
func (s ApiRequestService) GetApiRequest(ctx context.Context, id string, tag string) error {
	apiRequest, err := mooc.NewApiRequest(id)
	if err != nil {
		return err
	}

	if err := s.apiRequestRepository.Save(ctx, apiRequest); err != nil {
		return err
	}
	fmt.Println(tag)
	fmt.Println(apiRequest.PullEvents())

	return s.eventBus.Publish(ctx, apiRequest.PullEvents())
}
