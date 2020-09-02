package mongodbatlas

import (
	"context"
	"fmt"
	"net/http"
)

const (
	integrationBasePath = "groups/%s/integrations"
)

// IntegrationsService is an interface for interfacing with the Third-Party Integrations
// endpoints of the MongoDB Atlas API.

// See more: https://docs.atlas.mongodb.com/reference/api/third-party-integration-settings/
type IntegrationsService interface {
	Create(context.Context, string, string, *ThirdPartyService) (*IntegrationResponse, *Response, error)
	Replace(context.Context, string, string, *ThirdPartyService) (*IntegrationResponse, *Response, error)
	Delete(context.Context, string, string) (*Response, error)
	Get(context.Context, string, string) (*ThirdPartyService, *Response, error)
	List(context.Context, string) (*IntegrationResponse, *Response, error)
}

// TeamsServiceOp handles communication with the Teams related methods of the
// MongoDB Atlas API
type IntegrationsServiceOp service

var _ IntegrationsService = &IntegrationsServiceOp{}

// IntegrationRequest contains parameters for different third-party services
type ThirdPartyService struct {
	Type        string `json:"type,omitempty"`
	LicenseKey  string `json:"licenseKey,omitempty"`
	AccountId   string `json:"accountId,omitempty"`
	WriteToken  string `json:"writeToken,omitempty"`
	ReadToken   string `json:"readToken,omitempty"`
	ApiKey      string `json:"apiKey,omitempty"`
	Region      string `json:"region,omitempty"`
	ServiceKey  string `json:"serviceKey,omitempty"`
	ApiToken    string `json:"apiToken,omitempty"`
	TeamName    string `json:"teamName,omitempty"`
	ChannelName string `json:"channelName,omitempty"`
	RoutingKey  string `json:"routingKey,omitempty"`
	FlowName    string `json:"flowName,omitempty"`
	OrgName     string `json:"orgName,omitempty"`
	Url         string `json:"url,omitempty"`
	Secret      string `json:"secret,omitempty"`
}

// IntegrationResponse contains the response from the endpoint
type IntegrationResponse struct {
	Links      []*Link              `json:"links"`
	Results    []*ThirdPartyService `json:"results"`
	TotalCount int                  `json:"totalCount"`
}

// Create adds a new third-party integration configuration.
//
// See more: https://docs.atlas.mongodb.com/reference/api/third-party-integration-settings-create/index.html
func (s *IntegrationsServiceOp) Create(ctx context.Context, projectID, integrationType string, body *ThirdPartyService) (*IntegrationResponse, *Response, error) {
	if projectID == "" {
		return nil, nil, NewArgError("projectID", "must be set")
	}

	if integrationType == "" {
		return nil, nil, NewArgError("integrationType", "must be set")
	}

	basePath := fmt.Sprintf(integrationBasePath, projectID)
	path := fmt.Sprintf("%s/%s", basePath, integrationType)

	req, err := s.Client.NewRequest(ctx, http.MethodPost, path, body)
	if err != nil {
		return nil, nil, err
	}

	root := new(IntegrationResponse)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	if l := root.Links; l != nil {
		resp.Links = l
	}

	return root, resp, nil
}

// Replace replaces the third-party integration configuration with a new configuration, or add a new configuration if there is no configuration.
//
// https://docs.atlas.mongodb.com/reference/api/third-party-integration-settings-update/
func (s *IntegrationsServiceOp) Replace(ctx context.Context, projectID, integrationType string, body *ThirdPartyService) (*IntegrationResponse, *Response, error) {
	if projectID == "" {
		return nil, nil, NewArgError("projectID", "must be set")
	}

	if integrationType == "" {
		return nil, nil, NewArgError("integrationType", "must be set")
	}

	basePath := fmt.Sprintf(integrationBasePath, projectID)
	path := fmt.Sprintf("%s/%s", basePath, integrationType)

	req, err := s.Client.NewRequest(ctx, http.MethodPut, path, body)
	if err != nil {
		return nil, nil, err
	}

	root := new(IntegrationResponse)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	if l := root.Links; l != nil {
		resp.Links = l
	}

	return root, resp, nil
}

// Delete removes the third-party integration configuration
//
// https://docs.atlas.mongodb.com/reference/api/third-party-integration-settings-delete/
func (s *IntegrationsServiceOp) Delete(ctx context.Context, projectID, integrationType string) (*Response, error) {
	if projectID == "" {
		return nil, NewArgError("projectID", "must be set")
	}

	if integrationType == "" {
		return nil, NewArgError("integrationType", "must be set")
	}

	basePath := fmt.Sprintf(integrationBasePath, projectID)
	path := fmt.Sprintf("%s/%s", basePath, integrationType)

	req, err := s.Client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(ctx, req, nil)
	return resp, nil
}

// Get retrieves a specific third-party integration configuration
//
// https://docs.atlas.mongodb.com/reference/api/third-party-integration-settings-get-one/
func (s *IntegrationsServiceOp) Get(ctx context.Context, projectID, integrationType string) (*ThirdPartyService, *Response, error) {
	if projectID == "" {
		return nil, nil, NewArgError("projectID", "must be set")
	}

	if integrationType == "" {
		return nil, nil, NewArgError("integrationType", "must be set")
	}

	basePath := fmt.Sprintf(integrationBasePath, projectID)
	path := fmt.Sprintf("%s/%s", basePath, integrationType)

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(ThirdPartyService)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, nil
}

// List retrieves all third-party integration configurations.
//
// See more: https://docs.atlas.mongodb.com/reference/api/third-party-integration-settings-get-all/
func (s *IntegrationsServiceOp) List(ctx context.Context, projectID string) (*IntegrationResponse, *Response, error) {
	if projectID == "" {
		return nil, nil, NewArgError("projectID", "must be set")
	}

	path := fmt.Sprintf(integrationBasePath, projectID)

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(IntegrationResponse)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	if l := root.Links; l != nil {
		resp.Links = l
	}

	return root, resp, nil
}
