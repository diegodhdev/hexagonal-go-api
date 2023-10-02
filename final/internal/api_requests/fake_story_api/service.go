package fake_story_api

import (
	"context"
	"encoding/json"
	"fmt"
	mooc "github.com/diegodhdev/hexagonal-go-api/final/internal"
	"github.com/diegodhdev/hexagonal-go-api/final/kit/command"
	"github.com/diegodhdev/hexagonal-go-api/final/kit/event"
	"io"
	"net/http"
	"os"
	"sync"
)

var wg sync.WaitGroup

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

// FakeStoryApiApiRequest implements the fake_story_api.FakeStoryApiApiRequest interface.
func (s ApiRequestService) FakeStoryApiApiRequest(data command.DataResponse, ctx context.Context, id string, api string, mode string, response string, requests []Request) (command.DataResponse, error) {
	var pr []ProductDummyJson
	apiRequest, err := mooc.NewApiRequest(id)
	if err != nil {
		return data, err
	}

	fmt.Println(data)
	//if err := s.apiRequestRepository.Save(ctx, apiRequest); err != nil {
	//	return err
	//}
	fmt.Println(apiRequest)
	fmt.Println(api)
	fmt.Println(mode)
	fmt.Println(response)
	fmt.Println(requests)

	channel := make(chan ProductDummyJson)

	var i int = 0
	for _, request := range requests {
		i++
		if mode == "async" {
			fmt.Println("async Line: ", i)
			wg.Add(1)
			go requestProductDummyJsonAsync(request.Url, channel, i)
		} else {
			fmt.Println("sync Line: ", i)
			pr = append(pr, requestProductDummyJsonSync(request.Url, i))
		}
	}

	if mode == "async" {
		go func() {
			wg.Wait()
			close(channel)
		}()

		for elem := range channel {
			pr = append(pr, elem)
		}

	}
	fmt.Println(pr)

	//return
	//return pr
	//ctx.JSON(http.StatusOK, pr)

	return command.NewDataResponse("Me cago en tus muertos"), s.eventBus.Publish(ctx, apiRequest.PullEvents())
}

type ProductDummyJson struct {
	Id          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Price       int      `json:"price"`
	Brand       string   `json:"brand"`
	Images      []string `json:"images"`
}

// func requestProduct(url string, c chan ProductDummyJson, order int) {
func requestProductDummyJsonSync(url string, request_order int) (p ProductDummyJson) {

	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	reqBody, _ := io.ReadAll(res.Body)

	//filesystem.Save(reqBody, "dummyjson/"+strconv.FormatInt(int64(request_order), 10)+".json")

	err = json.Unmarshal(reqBody, &p)
	if err != nil {
	}

	res.Body.Close()

	return p
}

func requestProductDummyJsonAsync(url string, c chan ProductDummyJson, request_order int) {
	defer wg.Done()
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	var p ProductDummyJson

	reqBody, _ := io.ReadAll(res.Body)

	//filesystem.Save(reqBody, "dummyjson/"+strconv.FormatInt(int64(request_order), 10)+".json")

	err = json.Unmarshal(reqBody, &p)
	if err != nil {
	}

	res.Body.Close()
	c <- p
}
