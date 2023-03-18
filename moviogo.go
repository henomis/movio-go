package moviogo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/henomis/restclientgo"

	"github.com/henomis/movio-go/pkg/request"
	"github.com/henomis/movio-go/pkg/response"
)

const (
	MovioEndpointV1 = "https://api.movio.la/v1"
)

var (
	ErrInvalidRequest = fmt.Errorf("invalid request")
)

type MovioClient struct {
	httpClient *restclientgo.RestClient
}

func New(endpoint string, apiToken string) *MovioClient {

	restClient := restclientgo.New(endpoint)
	restClient.SetRequestModifier(func(req *http.Request) *http.Request {
		req.Header.Set("X-Api-Key", apiToken)
		return req
	})

	return &MovioClient{
		httpClient: restClient,
	}
}

func (m *MovioClient) SetHTTPClient(client *http.Client) {
	m.httpClient.SetHTTPClient(client)
}

func (i *MovioClient) ListAvatars(ctx context.Context, listAvatarsRequest *request.ListAvatars) (*response.ListAvatars, error) {

	listAvatarsResponse := &response.ListAvatars{}
	err := i.httpClient.Get(ctx, listAvatarsRequest, listAvatarsResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrInvalidRequest, err)
	}

	return listAvatarsResponse, nil
}

func (i *MovioClient) GenerateVideo(ctx context.Context, generateVideoRequest *request.GenerateVideo) (*response.GenerateVideo, error) {

	generateVideoResponse := &response.GenerateVideo{}
	err := i.httpClient.Post(ctx, generateVideoRequest, generateVideoResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrInvalidRequest, err)
	}

	return generateVideoResponse, nil
}

func (i *MovioClient) GetVideoStatus(ctx context.Context, videoStatusRequest *request.GetVideoStatus) (*response.GetVideoStatus, error) {

	videoStatusResponse := &response.GetVideoStatus{}
	err := i.httpClient.Get(ctx, videoStatusRequest, videoStatusResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrInvalidRequest, err)
	}

	return videoStatusResponse, nil
}

func (i *MovioClient) GetTemplate(ctx context.Context, getTemplateRequest *request.GetTemplate) (*response.GetTemplate, error) {

	getTemplateResponse := &response.GetTemplate{}
	err := i.httpClient.Get(ctx, getTemplateRequest, getTemplateResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrInvalidRequest, err)
	}

	return getTemplateResponse, nil
}

func (i *MovioClient) ListVoices(ctx context.Context, listVoicesRequest *request.ListVoices) (*response.ListVoices, error) {

	listVoicesResponse := &response.ListVoices{}
	err := i.httpClient.Get(ctx, listVoicesRequest, listVoicesResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrInvalidRequest, err)
	}

	return listVoicesResponse, nil
}

func (i *MovioClient) GenerateTemplate(ctx context.Context, generateTemplateRequest *request.GenerateTemplate) (*response.GenerateTemplate, error) {

	generateTemplateResponse := &response.GenerateTemplate{}
	err := i.httpClient.Post(ctx, generateTemplateRequest, generateTemplateResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrInvalidRequest, err)
	}

	return generateTemplateResponse, nil
}

func (i *MovioClient) GetUploadURL(ctx context.Context, getUploadURLRequest *request.GetUploadURL) (*response.GetUploadURL, error) {

	getUploadURLResponse := &response.GetUploadURL{}
	err := i.httpClient.Get(ctx, getUploadURLRequest, getUploadURLResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrInvalidRequest, err)
	}

	return getUploadURLResponse, nil
}

func (i *MovioClient) CreateTalkingPhoto(ctx context.Context, createTalkingPhotoRequest *request.CreateTalkingPhoto) (*response.CreateTalkingPhoto, error) {

	createTalkingPhotoResponse := &response.CreateTalkingPhoto{}
	err := i.httpClient.Get(ctx, createTalkingPhotoRequest, createTalkingPhotoResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrInvalidRequest, err)
	}

	return createTalkingPhotoResponse, nil
}

func (i *MovioClient) ListTalkingPhotos(ctx context.Context, talkingPhotosListRequest *request.ListTalkingPhotos) (*response.ListTalkingPhotos, error) {

	listTalkingPhotosResponse := &response.ListTalkingPhotos{}
	err := i.httpClient.Get(ctx, talkingPhotosListRequest, listTalkingPhotosResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrInvalidRequest, err)
	}

	return listTalkingPhotosResponse, nil
}
