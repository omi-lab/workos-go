package passwordless

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/omi-lab/workos-go/v4/pkg/models"
	"github.com/omi-lab/workos-go/v4/pkg/workos_errors"

	"github.com/omi-lab/workos-go/v4/internal/workos"
)

// Client represents a client that performs Passwordless requests to the WorkOS API.
type Client struct {
	// The WorkOS API Key.
	// It can be found in https://dashboard.workos.com/api-keys.
	//
	// REQUIRED
	APIKey string

	// The http.Client that is used to send request to WorkOS.
	//
	// Defaults to http.Client.
	HTTPClient *http.Client

	// The endpoint to WorkOS API.
	//
	// Defaults to https://api.workos.com.
	Endpoint string

	// The function used to encode in JSON. Defaults to json.Marshal.
	JSONEncode func(v interface{}) ([]byte, error)

	once sync.Once
}

func (c *Client) init() {
	if c.HTTPClient == nil {
		c.HTTPClient = &http.Client{Timeout: 10 * time.Second}
	}

	if c.Endpoint == "" {
		c.Endpoint = "https://api.workos.com"
	}

	if c.JSONEncode == nil {
		c.JSONEncode = json.Marshal
	}
}

// CreateSessionOpts contains the options to create a Passowordless Session.
type CreateSessionOpts struct {
	// The email of the user to authenticate.
	//
	// REQUIRED
	Email string `json:"email"`

	// The type of Passwordless Session to create.
	//
	// REQUIRED
	Type models.PasswordlessSessionType `json:"type"`

	// Optional The unique identifier for a WorkOS Connection.
	Connection string `json:"connection"`

	// Optional string value used to set the location
	// that the user will be redirected to after authenticating
	RedirectURI string `json:"redirect_uri"`

	// Optional string value used to manage application state
	// between authorization transactions.
	State string `json:"state"`

	// Optional The number of seconds the Passwordless Session
	// should live before expiring.
	ExpiresIn int `json:"expires_in"`
}

// CreateSession creates a a PasswordlessSession.
func (c *Client) CreateSession(ctx context.Context, opts CreateSessionOpts) (models.PasswordlessSession, error) {
	c.once.Do(c.init)

	data, err := c.JSONEncode(opts)
	if err != nil {
		return models.PasswordlessSession{}, err
	}

	endpoint := fmt.Sprintf("%s/passwordless/sessions", c.Endpoint)
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(data))
	if err != nil {
		return models.PasswordlessSession{}, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("User-Agent", "workos-go/"+workos.Version)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return models.PasswordlessSession{}, err
	}
	defer res.Body.Close()

	if err = workos_errors.TryGetHTTPError(res); err != nil {
		return models.PasswordlessSession{}, err
	}

	var body models.PasswordlessSession
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&body)
	return body, err
}

// SendSessionOpts contains the options to send a Passwordless Session via email.
type SendSessionOpts struct {
	// Passwordless Session unique identifier.
	SessionID string
}

// SendSession sends a Passwordless Session via email
func (c *Client) SendSession(
	ctx context.Context,
	opts SendSessionOpts,
) error {
	c.once.Do(c.init)

	data, err := c.JSONEncode(opts)
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf(
		"%s/passwordless/sessions/%s/send",
		c.Endpoint,
		opts.SessionID,
	)
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("User-Agent", "workos-go/"+workos.Version)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return workos_errors.TryGetHTTPError(res)
}
