package mailersend

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const smsWebhookPath = "/sms-webhooks"

type SmsWebhookService service

// smsWebhookRoot - format of activity response
type smsWebhookRoot struct {
	Data  []smsWebhook `json:"data"`
	Links Links        `json:"links"`
	Meta  Meta         `json:"meta"`
}

// singleSmsNumberRoot - format of activity response
type singleSmsWebhookRoot struct {
	Data smsWebhook `json:"data"`
}

type smsWebhook struct {
	Id        string    `json:"id"`
	Url       string    `json:"url"`
	Events    []string  `json:"events"`
	Name      string    `json:"name"`
	Enabled   bool      `json:"enabled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	SmsNumber Number    `json:"sms_number"`
}

// CreateSmsWebhookOptions - modifies the behavior of *WebhookService.Create Method
type CreateSmsWebhookOptions struct {
	SmsNumberId string   `json:"sms_number_id"`
	Name        string   `json:"name"`
	URL         string   `json:"url"`
	Enabled     *bool    `json:"enabled,omitempty"`
	Events      []string `json:"events"`
}

// UpdateSmsWebhookOptions - modifies the behavior of SmsNumbersService.Update method
type UpdateSmsWebhookOptions struct {
	Id      string   `json:"-"`
	URL     string   `json:"url,omitempty"`
	Name    string   `json:"name,omitempty"`
	Events  []string `json:"events,omitempty"`
	Status  string   `json:"status,omitempty"`
	Enabled *bool    `json:"enabled,omitempty"`
}

// ListSmsWebhookOptions - modifies the behavior of SmsNumbersService.List method
type ListSmsWebhookOptions struct {
	SmsNumberId string `url:"sms_number_id,omitempty"`
	Page        int    `url:"page,omitempty"`
	Limit       int    `url:"limit,omitempty"`
}

func (s *SmsWebhookService) List(ctx context.Context, options *ListSmsWebhookOptions) (*smsWebhookRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, smsWebhookPath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(smsWebhookRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *SmsWebhookService) Get(ctx context.Context, smsWebhookId string) (*singleSmsWebhookRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", smsWebhookPath, smsWebhookId)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(singleSmsWebhookRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *SmsWebhookService) Create(ctx context.Context, options *CreateSmsWebhookOptions) (*singleSmsWebhookRoot, *Response, error) {
	req, err := s.client.newRequest(http.MethodPost, smsWebhookPath, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(singleSmsWebhookRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *SmsWebhookService) Update(ctx context.Context, options *UpdateSmsWebhookOptions) (*singleSmsWebhookRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", smsWebhookPath, options.Id)

	req, err := s.client.newRequest(http.MethodPut, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(singleSmsWebhookRoot)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *SmsWebhookService) Delete(ctx context.Context, smsWebhookId string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", smsWebhookPath, smsWebhookId)

	req, err := s.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}
