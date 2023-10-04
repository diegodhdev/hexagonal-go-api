package mooc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/diegodhdev/hexagonal-go-api/requests/kit/event"
	"github.com/google/uuid"
)

var ErrInvalidApiRequestID = errors.New("invalid Api Request ID")

type Request struct {
	// json tag to de-serialize json body
	Url        string     `json:"url"`
	Credential Credential `json:"credential"`
}

type Credential struct {
	Token string `json:"token"`
}

type ApiRequestID struct {
	value string
}
type ApiRequestApi struct {
	value string
}
type ApiRequestMode struct {
	value string
}
type ApiRequestResponseType struct {
	value string
}
type ApiRequestRequest struct {
	value []Request
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

func NewApiRequestMode(value string) (ApiRequestMode, error) {

	return ApiRequestMode{
		value: value,
	}, nil
}

func (mode ApiRequestMode) String() string {
	return mode.value
}

func NewApiRequestApi(value string) (ApiRequestApi, error) {

	return ApiRequestApi{
		value: value,
	}, nil
}

func (api ApiRequestApi) String() string {
	return api.value
}

func NewApiRequestResponseType(value string) (ApiRequestResponseType, error) {

	return ApiRequestResponseType{
		value: value,
	}, nil
}

func (responseType ApiRequestResponseType) String() string {
	return responseType.value
}

func NewApiRequestRequest(value []Request) (ApiRequestRequest, error) {

	return ApiRequestRequest{
		value: value,
	}, nil
}

func (request ApiRequestRequest) String() string {

	data, _ := json.Marshal(request.value)

	return string(data)
}

// ApiRequest is the data structure that represents a api request.
type ApiRequest struct {
	id           ApiRequestID
	api          ApiRequestApi
	mode         ApiRequestMode
	responseType ApiRequestResponseType
	request      ApiRequestRequest

	events []event.Event
}

// ApiRequestRepository defines the expected behaviour from a course storage.
type ApiRequestRepository interface {
	Save(ctx context.Context, apiRequest ApiRequest) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=CourseRepository

// NewApiRequest creates an api request.
func NewApiRequest(id string, api string, mode string, responseType string, request []Request) (ApiRequest, error) {
	idVO, err := NewApiRequestID(id)
	if err != nil {
		return ApiRequest{}, err
	}

	apiVO, err := NewApiRequestApi(api)
	modeVO, err := NewApiRequestMode(mode)
	responseTypeVO, err := NewApiRequestResponseType(responseType)
	requestVO, err := NewApiRequestRequest(request)

	apiRequest := ApiRequest{
		id:           idVO,
		api:          apiVO,
		mode:         modeVO,
		responseType: responseTypeVO,
		request:      requestVO,
	}
	return apiRequest, nil
}

func (a ApiRequest) ID() ApiRequestID {
	return a.id
}
func (a ApiRequest) Api() ApiRequestApi {
	return a.api
}
func (a ApiRequest) Mode() ApiRequestMode {
	return a.mode
}
func (a ApiRequest) ResponseType() ApiRequestResponseType {
	return a.responseType
}
func (a ApiRequest) Request() ApiRequestRequest {
	return a.request
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
