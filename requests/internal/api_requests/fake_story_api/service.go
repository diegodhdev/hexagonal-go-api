package fake_story_api

import (
	"context"
	"encoding/json"
	"fmt"
	mooc "github.com/diegodhdev/hexagonal-go-api/requests/internal"
	"github.com/diegodhdev/hexagonal-go-api/requests/internal/platform/storage/filesystem"
	"github.com/diegodhdev/hexagonal-go-api/requests/kit/event"
	"io"
	"net/http"
	"strconv"
	"sync"
)

const folderResponseFiles = "fake_story_api"

var wg sync.WaitGroup

type ApiRequestService struct {
	apiRequestRepository mooc.ApiRequestRepository
	eventBus             event.Bus
	filesystem           filesystem.Filesystem
}

func NewApiRequestService(apiRequestRepository mooc.ApiRequestRepository, eventBus event.Bus, customFilesystem filesystem.Filesystem) ApiRequestService {
	return ApiRequestService{
		apiRequestRepository: apiRequestRepository,
		eventBus:             eventBus,
		filesystem:           customFilesystem,
	}
}

func (s ApiRequestService) FakeStoryApiApiRequest(ctx context.Context, id string, api string, mode string, response string, requests []mooc.Request) (any, error) {
	var pr ProductFakeStoryApiDataResponse
	apiRequest, err := mooc.NewApiRequest(id, api, mode, response, requests)
	if err != nil {
		return data, err
	}

	err = s.apiRequestRepository.Save(ctx, apiRequest)

	if err := s.apiRequestRepository.Save(ctx, apiRequest); err != nil {
		return nil, err
	}

	channel := make(chan FakeStoryApi)
	var i int = 0
	for _, request := range requests {
		i++
		if mode == "async" {
			wg.Add(1)
			go func() {
				err := requestProductFakeStoryApiAsync(s, request.Url, channel, id, i)
				if err != nil {
					fmt.Errorf("error executing async request function: %s\n", err.Error())
				}
			}()
		} else {

			var p, err = requestProductFakeStoryApiSync(s, request.Url, id, i)

			if err != nil {
				return pr, fmt.Errorf("error executing sync request function: %s\n", err.Error())
			}

			pr.Data = append(pr.Data, p)
		}
	}

	if mode == "async" {
		go func() {
			wg.Wait()
			close(channel)
		}()

		for elem := range channel {
			pr.Data = append(pr.Data, elem)
		}
	}

	return pr, s.eventBus.Publish(ctx, apiRequest.PullEvents())
}

type FakeStoryApi struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Rating      Rate    `json:"rating"`
}

type Rate struct {
	Rate  float32 `json:"rate"`
	Count int     `json:"count"`
}

type ProductFakeStoryApiDataResponse struct {
	Data []FakeStoryApi
}

func NewProductFakeStoryApiDataResponse() ProductFakeStoryApiDataResponse {
	return ProductFakeStoryApiDataResponse{}
}

// func requestProduct(url string, c chan FakeStoryApi, order int) {
func requestProductFakeStoryApiSync(s ApiRequestService, url string, id string, request_order int) (p FakeStoryApi, err error) {
	res, err := http.Get(url)
	if err != nil {
		return FakeStoryApi{}, fmt.Errorf("error making http request: %s\n", err)
	}

	reqBody, _ := io.ReadAll(res.Body)

	err = s.filesystem.Save(reqBody, folderResponseFiles+"/"+id+"/", strconv.FormatInt(int64(request_order), 10)+".json")
	if err != nil {
		return FakeStoryApi{}, fmt.Errorf("error saving file on filesystem: %s\n", err.Error())
	}

	err = json.Unmarshal(reqBody, &p)
	if err != nil {
		return FakeStoryApi{}, fmt.Errorf("error unmarshalling http request: %s\n", err.Error())
	}

	res.Body.Close()

	return p, nil
}

func requestProductFakeStoryApiAsync(s ApiRequestService, url string, c chan FakeStoryApi, id string, request_order int) error {
	defer wg.Done()
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error making http request: %s\n", err.Error())
	}

	var p FakeStoryApi

	reqBody, _ := io.ReadAll(res.Body)

	err = s.filesystem.Save(reqBody, folderResponseFiles+"/"+id+"/", strconv.FormatInt(int64(request_order), 10)+".json")
	if err != nil {
		return fmt.Errorf("error saving file on filesystem: %s\n", err.Error())
	}

	err = json.Unmarshal(reqBody, &p)
	if err != nil {
		return fmt.Errorf("error unmarshalling http request: %s\n", err.Error())
	}

	res.Body.Close()
	c <- p

	return nil
}
