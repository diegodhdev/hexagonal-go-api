package fake_story_api

import (
	"context"
	"encoding/json"
	"errors"
	mooc "github.com/diegodhdev/hexagonal-go-api/requests/internal"
	filesystem2 "github.com/diegodhdev/hexagonal-go-api/requests/internal/platform/storage/filesystem"
	"github.com/diegodhdev/hexagonal-go-api/requests/internal/platform/storage/storagemocks"
	"github.com/diegodhdev/hexagonal-go-api/requests/kit/event"
	"github.com/diegodhdev/hexagonal-go-api/requests/kit/event/eventmocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_ApiRequestService_FakeStoryApiApiRequest_RepositoryError(t *testing.T) {

	apiRequestID := "6df370cc-5de1-11ee-b1f7-18dbf22dbbca"
	apiRequestApi := "test-api"
	apiRequestMode := "sync"
	apiRequestResponse := "direct"
	var jsonApiRequestRequests string = `[
		{
			"url": "https://fakestoreapi.com/products/2"
		},
		{
			"url": "https://fakestoreapi.com/products/3"
		}
	]`

	apiRequestRequests := []mooc.Request{}

	err := json.Unmarshal([]byte(jsonApiRequestRequests), &apiRequestRequests)
	if err != nil {
		return
	}

	apiRequestRepositoryMock := new(storagemocks.ApiRequestRepository)
	apiRequestRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("mooc.ApiRequest")).Return(errors.New("something unexpected happened"))

	customFilesystem := filesystem2.NewFilesystem("/home/diego/GolandProjects/hexagonal-go-api/requests/storage/tests/")

	eventBusMock := new(eventmocks.Bus)

	apiRequestService := NewApiRequestService(apiRequestRepositoryMock, eventBusMock, *customFilesystem)

	_, err = apiRequestService.FakeStoryApiApiRequest(context.Background(), apiRequestID, apiRequestApi, apiRequestMode, apiRequestResponse, apiRequestRequests)

	apiRequestRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_ApiRequestService_FakeStoryApiApiRequest_Succeed(t *testing.T) {

	apiRequestID := "6df370cc-5de1-11ee-b1f7-18dbf22dbbca"
	apiRequestApi := "test-api"
	apiRequestMode := "sync"
	apiRequestResponse := "direct"
	var jsonApiRequestRequests string = `[
		{
			"url": "https://fakestoreapi.com/products/2"
		},
		{
			"url": "https://fakestoreapi.com/products/3"
		}
	]`

	apiRequestRequests := []mooc.Request{}

	err := json.Unmarshal([]byte(jsonApiRequestRequests), &apiRequestRequests)
	if err != nil {
		return
	}
	customFilesystem := filesystem2.NewFilesystem("/home/diego/GolandProjects/hexagonal-go-api/requests/storage/tests/")

	apiRequestRepositoryMock := new(storagemocks.ApiRequestRepository)
	apiRequestRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("mooc.ApiRequest")).Return(nil)
	eventBusMock := new(eventmocks.Bus)
	eventBusMock.On("Publish", mock.Anything, mock.MatchedBy(func(events []event.Event) bool {
		return true
	})).Return(nil)

	apiRequestService := NewApiRequestService(apiRequestRepositoryMock, eventBusMock, *customFilesystem)

	_, err = apiRequestService.FakeStoryApiApiRequest(context.Background(), apiRequestID, apiRequestApi, apiRequestMode, apiRequestResponse, apiRequestRequests)

	apiRequestRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.NoError(t, err)
}
