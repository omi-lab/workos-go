package auditlogs

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/omi-lab/workos-go/v4/pkg/models"
	"github.com/omi-lab/workos-go/v4/pkg/workos_errors"

	"github.com/omi-lab/workos-go/v4/internal/workos"
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

// Client represents a client that performs auditlogs requests to WorkOS API.
type Client struct {
	// The WorkOS API key. It can be found in
	// https://dashboard.workos.com/api-keys.
	APIKey string

	// The http.Client that is used to post Audit Log events to WorkOS. Defaults
	// to http.Client.
	HTTPClient *http.Client

	// The endpoint used to request WorkOS AuditLog events creation endpoint.
	// Defaults to https://api.workos.com/audit_logs/events.
	EventsEndpoint string

	// The endpoint used to request WorkOS AuditLog events creation endpoint.
	// Defaults to https://api.workos.com/audit_logs/exports.
	ExportsEndpoint string

	// The function used to encode in JSON. Defaults to json.Marshal.
	JSONEncode func(v interface{}) ([]byte, error)

	once sync.Once
}

// CreateEventOpts represents arguments to create an Audit Logs event.
type CreateEventOpts struct {
	// Organization identifier
	OrganizationID string `json:"organization_id" binding:"required"`

	// Event payload
	Event models.AuditLogEvent `json:"event" binding:"required"`

	// If no key is provided or the key is empty, the key will not be attached
	// to the request.
	IdempotencyKey string `json:"-"`
}

type CreateExportOpts struct {
	// Organization identifier
	OrganizationID string `json:"organization_id"`

	// ISO-8601 start datetime the date range filter
	RangeStart string `json:"range_start"`

	// ISO-8601 start datetime the date range filter
	RangeEnd string `json:"range_end"`

	// Optional list of actions to filter
	Actions []string `json:"actions,omitempty"`

	// Deprecated - use `ActorNames` instead
	Actors []string `json:"actors,omitempty"`

	// Optional list of actor names to filter by
	ActorNames []string `json:"actor_names,omitempty"`

	// Optional list of actor ids to filter by
	ActorIds []string `json:"actor_ids,omitempty"`

	// Optional list of targets to filter
	Targets []string `json:"targets,omitempty"`
}

type GetExportOpts struct {
	ExportID string `json:"export_id" binding:"required"`
}

func (c *Client) init() {
	if c.HTTPClient == nil {
		c.HTTPClient = &http.Client{Timeout: 10 * time.Second}
	}

	if c.EventsEndpoint == "" {
		c.EventsEndpoint = "https://api.workos.com/audit_logs/events"
	}

	if c.ExportsEndpoint == "" {
		c.ExportsEndpoint = "https://api.workos.com/audit_logs/exports"
	}

	if c.JSONEncode == nil {
		c.JSONEncode = json.Marshal
	}
}

// CreateEvent creates an Audit Log event.
func (c *Client) CreateEvent(ctx context.Context, e CreateEventOpts) error {
	c.once.Do(c.init)

	e.Event.OccurredAt = defaultTime(e.Event.OccurredAt)

	data, err := c.JSONEncode(e)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, c.EventsEndpoint, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("User-Agent", "workos-go/"+workos.Version)

	if e.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", e.IdempotencyKey)
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return workos_errors.TryGetHTTPError(res)
}

// CreateExport creates an export of Audit Log events. You can specify some filters.
func (c *Client) CreateExport(ctx context.Context, e CreateExportOpts) (models.AuditLogExport, error) {
	c.once.Do(c.init)

	data, err := c.JSONEncode(e)
	if err != nil {
		return models.AuditLogExport{}, err
	}

	req, err := http.NewRequest(http.MethodPost, c.ExportsEndpoint, bytes.NewBuffer(data))
	if err != nil {
		return models.AuditLogExport{}, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("User-Agent", "workos-go/"+workos.Version)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return models.AuditLogExport{}, err
	}
	defer res.Body.Close()

	if err = workos_errors.TryGetHTTPError(res); err != nil {
		return models.AuditLogExport{}, err
	}

	var body models.AuditLogExport
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&body)
	return body, err
}

// GetExport retrieves an export of Audit Log events
func (c *Client) GetExport(ctx context.Context, e GetExportOpts) (models.AuditLogExport, error) {
	c.once.Do(c.init)

	req, err := http.NewRequest(http.MethodGet, c.ExportsEndpoint+"/"+e.ExportID, nil)
	if err != nil {
		return models.AuditLogExport{}, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("User-Agent", "workos-go/"+workos.Version)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return models.AuditLogExport{}, err
	}
	defer res.Body.Close()

	if err = workos_errors.TryGetHTTPError(res); err != nil {
		return models.AuditLogExport{}, err
	}

	var body models.AuditLogExport
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&body)
	return body, err
}

func defaultTime(t time.Time) time.Time {
	if t == (time.Time{}) {
		t = time.Now().UTC()
	}
	return t
}
