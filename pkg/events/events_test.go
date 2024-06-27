package events

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/omi-lab/workos-go/v4/pkg/common"
	"github.com/omi-lab/workos-go/v4/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestEventsListEvents(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(ListEventsTestHandler))
	defer server.Close()

	DefaultClient = &Client{
		HTTPClient: server.Client(),
		Endpoint:   server.URL,
	}
	SetAPIKey("test")

	params := ListEventsOpts{
		Events: []string{"dsync.user.created"},
	}

	expectedResponse := ListEventsResponse{
		Data: []models.Event{
			{
				ID:    "event_abcd1234",
				Event: "dsync.user.created",
				Data:  json.RawMessage(`{"foo":"bar"}`),
			},
		},
		ListMetadata: common.ListMetadata{
			After: "",
		},
	}
	eventsResponse, err := ListEvents(context.Background(), params)

	require.NoError(t, err)
	require.Equal(t, expectedResponse, eventsResponse)
}
