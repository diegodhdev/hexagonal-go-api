package get

import (
	"context"
	mooc "github.com/diegodhdev/hexagonal-go-api/final/internal"
	"github.com/diegodhdev/hexagonal-go-api/final/kit/command"
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
func (s ApiRequestService) GetApiRequest(ctx context.Context, id string, tag string) (any, error) {
	apiRequest, err := mooc.NewApiRequest(id)
	if err != nil {
		return command.NewEmptyDataResponse(), err
	}

	if err := s.apiRequestRepository.Save(ctx, apiRequest); err != nil {
		return command.NewEmptyDataResponse(), err
	}

	return command.NewEmptyDataResponse(), s.eventBus.Publish(ctx, apiRequest.PullEvents())
}
