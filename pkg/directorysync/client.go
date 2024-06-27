package directorysync

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/omi-lab/workos-go/v4/pkg/models"
	"github.com/omi-lab/workos-go/v4/pkg/workos_errors"

	"github.com/omi-lab/workos-go/v4/internal/workos"
	"github.com/omi-lab/workos-go/v4/pkg/common"
)

// ResponseLimit is the default number of records to limit a response to.
const ResponseLimit = 10

// Order represents the order of records.
type Order string

// Constants that enumerate the available orders.
const (
	Asc  Order = "asc"
	Desc Order = "desc"
)

// Client represents a client that performs Directory Sync requests to the WorkOS API.
type Client struct {
	// The WorkOS API Key. It can be found in https://dashboard.workos.com/api-keys.
	APIKey string

	// The http.Client that is used to get Directory Sync records from WorkOS.
	// Defaults to http.Client.
	HTTPClient *http.Client

	// The endpoint to WorkOS API. Defaults to https://api.workos.com.
	Endpoint string

	once sync.Once
}

func (c *Client) init() {
	if c.HTTPClient == nil {
		c.HTTPClient = &http.Client{Timeout: 10 * time.Second}
	}

	if c.Endpoint == "" {
		c.Endpoint = "https://api.workos.com"
	}
}

// ListUsersOpts contains the options to request provisioned Directory Users.
type ListUsersOpts struct {
	// Directory unique identifier.
	Directory string `url:"directory,omitempty"`

	// Directory Group unique identifier.
	Group string `url:"group,omitempty"`

	// Maximum number of records to return.
	Limit int `url:"limit"`

	// The order in which to paginate records.
	Order Order `url:"order,omitempty"`

	// Pagination cursor to receive records before a provided User ID.
	Before string `url:"before,omitempty"`

	// Pagination cursor to receive records after a provided User ID.
	After string `url:"after,omitempty"`
}

// ListUsersResponse describes the response structure when requesting
// provisioned Directory Users.
type ListUsersResponse struct {
	// List of provisioned Users.
	Data []models.DirectoryUser `json:"data"`

	// Cursor pagination options.
	ListMetadata common.ListMetadata `json:"listMetadata"`
}

// ListUsers gets a list of provisioned Users for a Directory.
func (c *Client) ListUsers(
	ctx context.Context,
	opts ListUsersOpts,
) (ListUsersResponse, error) {
	c.once.Do(c.init)

	endpoint := fmt.Sprintf("%s/directory_users", c.Endpoint)
	req, err := http.NewRequest(
		http.MethodGet,
		endpoint,
		nil,
	)
	if err != nil {
		return ListUsersResponse{}, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "workos-go/"+workos.Version)
	if opts.Limit == 0 {
		opts.Limit = ResponseLimit
	}

	if opts.Order == "" {
		opts.Order = Desc
	}

	v, err := query.Values(opts)
	if err != nil {
		return ListUsersResponse{}, err
	}

	req.URL.RawQuery = v.Encode()
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return ListUsersResponse{}, err
	}
	defer res.Body.Close()

	if err = workos_errors.TryGetHTTPError(res); err != nil {
		return ListUsersResponse{}, err
	}

	var body ListUsersResponse
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&body)
	return body, err
}

// ListGroupsOpts contains the options to request provisioned Directory Groups.
type ListGroupsOpts struct {
	// Directory unique identifier.
	Directory string `url:"directory,omitempty"`

	// Directory unique identifier.
	User string `url:"user,omitempty"`

	// Maximum number of records to return.
	Limit int `url:"limit"`

	// The order in which to paginate records.
	Order Order `url:"order,omitempty"`

	// Pagination cursor to receive records before a provided Group ID.
	Before string `url:"before,omitempty"`

	// Pagination cursor to receive records after a provided Group ID.
	After string `url:"after,omitempty"`
}

// ListGroupsResponse describes the response structure when requesting
// provisioned Directory Groups.
type ListGroupsResponse struct {
	// List of provisioned Users.
	Data []models.DirectoryGroup `json:"data"`

	// Cursor pagination options.
	ListMetadata common.ListMetadata `json:"listMetadata"`
}

// ListGroups gets a list of provisioned Groups for a Directory Endpoint.
func (c *Client) ListGroups(
	ctx context.Context,
	opts ListGroupsOpts,
) (ListGroupsResponse, error) {
	c.once.Do(c.init)

	endpoint := fmt.Sprintf("%s/directory_groups", c.Endpoint)
	req, err := http.NewRequest(
		http.MethodGet,
		endpoint,
		nil,
	)
	if err != nil {
		return ListGroupsResponse{}, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "workos-go/"+workos.Version)

	if opts.Limit == 0 {
		opts.Limit = ResponseLimit
	}

	if opts.Order == "" {
		opts.Order = Desc
	}

	v, err := query.Values(opts)
	if err != nil {
		return ListGroupsResponse{}, err
	}

	req.URL.RawQuery = v.Encode()
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return ListGroupsResponse{}, err
	}
	defer res.Body.Close()

	if err = workos_errors.TryGetHTTPError(res); err != nil {
		return ListGroupsResponse{}, err
	}

	var body ListGroupsResponse
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&body)
	return body, err
}

// GetUserOpts contains the options to request details for a provisioned Directory User.
type GetUserOpts struct {
	// Directory User unique identifier.
	User string
}

// GetUser gets a provisioned User for a Directory Endpoint.
func (c *Client) GetUser(
	ctx context.Context,
	opts GetUserOpts,
) (models.DirectoryUser, error) {
	c.once.Do(c.init)

	endpoint := fmt.Sprintf(
		"%s/directory_users/%s",
		c.Endpoint,
		opts.User,
	)
	req, err := http.NewRequest(
		http.MethodGet,
		endpoint,
		nil,
	)
	if err != nil {
		return models.DirectoryUser{}, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "workos-go/"+workos.Version)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return models.DirectoryUser{}, err
	}
	defer res.Body.Close()

	if err = workos_errors.TryGetHTTPError(res); err != nil {
		return models.DirectoryUser{}, err
	}

	var body models.DirectoryUser
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&body)
	return body, err
}

// GetGroupOpts contains the options to request details for a provisioned Directory Group.
type GetGroupOpts struct {
	// Directory Group unique identifier.
	Group string
}

// GetGroup gets a provisioned Group for a Directory.
func (c *Client) GetGroup(
	ctx context.Context,
	opts GetGroupOpts,
) (models.DirectoryGroup, error) {
	c.once.Do(c.init)

	endpoint := fmt.Sprintf(
		"%s/directory_groups/%s",
		c.Endpoint,
		opts.Group,
	)
	req, err := http.NewRequest(
		http.MethodGet,
		endpoint,
		nil,
	)
	if err != nil {
		return models.DirectoryGroup{}, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "workos-go/"+workos.Version)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return models.DirectoryGroup{}, err
	}
	defer res.Body.Close()

	if err = workos_errors.TryGetHTTPError(res); err != nil {
		return models.DirectoryGroup{}, err
	}

	var body models.DirectoryGroup
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&body)
	return body, err
}

// ListDirectoriesOpts contains the options to request a Project's Directories.
type ListDirectoriesOpts struct {
	// Domain of a Directory. Can be empty.
	Domain string `url:"domain,omitempty"`

	// Searchable text for a Directory. Can be empty.
	Search string `url:"search,omitempty"`

	// Organization ID of a Directory. Can be empty.
	OrganizationID string `url:"organization_id,omitempty"`

	// Maximum number of records to return.
	Limit int `url:"limit"`

	// The order in which to paginate records.
	Order Order `url:"order,omitempty"`

	// Pagination cursor to receive records before a provided Directory ID.
	Before string `url:"before,omitempty"`

	// Pagination cursor to receive records after a provided Directory ID.
	After string `url:"after,omitempty"`
}

// ListDirectoriesResponse describes the response structure when requesting
// existing Directories.
type ListDirectoriesResponse struct {
	// List of Directories.
	Data []models.Directory `json:"data"`

	// Cursor pagination options.
	ListMetadata common.ListMetadata `json:"listMetadata"`
}

// ListDirectories gets details of existing Directories.
func (c *Client) ListDirectories(
	ctx context.Context,
	opts ListDirectoriesOpts,
) (ListDirectoriesResponse, error) {
	c.once.Do(c.init)

	endpoint := fmt.Sprintf("%s/directories", c.Endpoint)
	req, err := http.NewRequest(
		http.MethodGet,
		endpoint,
		nil,
	)
	if err != nil {
		return ListDirectoriesResponse{}, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "workos-go/"+workos.Version)
	if opts.Limit == 0 {
		opts.Limit = ResponseLimit
	}

	if opts.Order == "" {
		opts.Order = Desc
	}

	v, err := query.Values(opts)
	if err != nil {
		return ListDirectoriesResponse{}, err
	}

	req.URL.RawQuery = v.Encode()
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return ListDirectoriesResponse{}, err
	}
	defer res.Body.Close()

	if err = workos_errors.TryGetHTTPError(res); err != nil {
		return ListDirectoriesResponse{}, err
	}
	var body ListDirectoriesResponse
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&body)
	return body, err
}

// GetDirectoryOpts contains the options to request details for an Directory.
type GetDirectoryOpts struct {
	// Directory unique identifier.
	Directory string
}

// GetDirectory gets a Directory.
func (c *Client) GetDirectory(
	ctx context.Context,
	opts GetDirectoryOpts,
) (models.Directory, error) {
	c.once.Do(c.init)

	endpoint := fmt.Sprintf(
		"%s/directories/%s",
		c.Endpoint,
		opts.Directory,
	)
	req, err := http.NewRequest(
		http.MethodGet,
		endpoint,
		nil,
	)
	if err != nil {
		return models.Directory{}, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "workos-go/"+workos.Version)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return models.Directory{}, err
	}
	defer res.Body.Close()

	if err = workos_errors.TryGetHTTPError(res); err != nil {
		return models.Directory{}, err
	}

	var body models.Directory
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&body)
	return body, err
}

// DeleteDirectoryOpts contains the options to delete a Connection.
type DeleteDirectoryOpts struct {
	// Directory unique identifier.
	Directory string
}

// DeleteDirectory deletes a Connection.
func (c *Client) DeleteDirectory(
	ctx context.Context,
	opts DeleteDirectoryOpts,
) error {
	c.once.Do(c.init)

	endpoint := fmt.Sprintf(
		"%s/directories/%s",
		c.Endpoint,
		opts.Directory,
	)
	req, err := http.NewRequest(
		http.MethodDelete,
		endpoint,
		nil,
	)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "workos-go/"+workos.Version)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return workos_errors.TryGetHTTPError(res)
}
