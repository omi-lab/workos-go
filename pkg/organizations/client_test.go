package organizations

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/omi-lab/workos-go/v4/pkg/common"
	"github.com/omi-lab/workos-go/v4/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestGetOrganization(t *testing.T) {
	tests := []struct {
		scenario string
		client   *Client
		options  GetOrganizationOpts
		expected models.Organization
		err      bool
	}{
		{
			scenario: "Request without API Key returns an error",
			client:   &Client{},
			err:      true,
		},
		{
			scenario: "Request returns an Organization",
			client: &Client{
				APIKey: "test",
			},
			options: GetOrganizationOpts{
				Organization: "organization_id",
			},
			expected: models.Organization{
				ID:                               "organization_id",
				Name:                             "Foo Corp",
				AllowProfilesOutsideOrganization: false,
				Domains: []models.OrganizationDomain{
					models.OrganizationDomain{
						ID:     "organization_domain_id",
						Domain: "foo-corp.com",
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(getOrganizationTestHandler))
			defer server.Close()

			client := test.client
			client.Endpoint = server.URL
			client.HTTPClient = server.Client()

			organization, err := client.GetOrganization(context.Background(), test.options)
			if test.err {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, test.expected, organization)
		})
	}
}

func getOrganizationTestHandler(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if auth != "Bearer test" {
		http.Error(w, "bad auth", http.StatusUnauthorized)
		return
	}

	body, err := json.Marshal(models.Organization{
		ID:                               "organization_id",
		Name:                             "Foo Corp",
		AllowProfilesOutsideOrganization: false,
		Domains: []models.OrganizationDomain{
			models.OrganizationDomain{
				ID:     "organization_domain_id",
				Domain: "foo-corp.com",
			},
		},
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func TestListOrganizations(t *testing.T) {
	tests := []struct {
		scenario string
		client   *Client
		options  ListOrganizationsOpts
		expected ListOrganizationsResponse
		err      bool
	}{
		{
			scenario: "Request without API Key returns an error",
			client:   &Client{},
			err:      true,
		},
		{
			scenario: "Request returns Organizations",
			client: &Client{
				APIKey: "test",
			},
			options: ListOrganizationsOpts{
				Domains: []string{"foo-corp.com"},
			},

			expected: ListOrganizationsResponse{
				Data: []models.Organization{
					models.Organization{
						ID:                               "organization_id",
						Name:                             "Foo Corp",
						AllowProfilesOutsideOrganization: false,
						Domains: []models.OrganizationDomain{
							models.OrganizationDomain{
								ID:     "organization_domain_id",
								Domain: "foo-corp.com",
							},
						},
					},
				},
				ListMetadata: common.ListMetadata{
					Before: "",
					After:  "",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(listOrganizationsTestHandler))
			defer server.Close()

			client := test.client
			client.Endpoint = server.URL
			client.HTTPClient = server.Client()

			organizations, err := client.ListOrganizations(context.Background(), test.options)
			if test.err {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, test.expected, organizations)
		})
	}
}

func listOrganizationsTestHandler(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if auth != "Bearer test" {
		http.Error(w, "bad auth", http.StatusUnauthorized)
		return
	}

	if userAgent := r.Header.Get("User-Agent"); !strings.Contains(userAgent, "workos-go/") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := json.Marshal(struct {
		ListOrganizationsResponse
	}{
		ListOrganizationsResponse: ListOrganizationsResponse{
			Data: []models.Organization{
				models.Organization{
					ID:                               "organization_id",
					Name:                             "Foo Corp",
					AllowProfilesOutsideOrganization: false,
					Domains: []models.OrganizationDomain{
						models.OrganizationDomain{
							ID:     "organization_domain_id",
							Domain: "foo-corp.com",
						},
					},
				},
			},
			ListMetadata: common.ListMetadata{
				Before: "",
				After:  "",
			},
		},
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func TestCreateOrganization(t *testing.T) {
	tests := []struct {
		scenario string
		client   *Client
		options  CreateOrganizationOpts
		expected models.Organization
		err      bool
	}{
		{
			scenario: "Request without API Key returns an error",
			client:   &Client{},
			err:      true,
		},
		{
			scenario: "Request returns Organization with Domains",
			client: &Client{
				APIKey: "test",
			},
			options: CreateOrganizationOpts{
				Name:    "Foo Corp",
				Domains: []string{"foo-corp.com"},
			},
			expected: models.Organization{
				ID:                               "organization_id",
				Name:                             "Foo Corp",
				AllowProfilesOutsideOrganization: false,
				Domains: []models.OrganizationDomain{
					models.OrganizationDomain{
						ID:     "organization_domain_id",
						Domain: "foo-corp.com",
					},
				},
			},
		},
		{
			scenario: "Request returns Organization with DomainData",
			client: &Client{
				APIKey: "test",
			},
			options: CreateOrganizationOpts{
				Name: "Foo Corp",
				DomainData: []models.OrganizationDomainData{
					models.OrganizationDomainData{
						Domain: "foo-corp.com",
						State:  models.OrganizationDomainDataStateVerified,
					},
				},
			},
			expected: models.Organization{
				ID:                               "organization_id",
				Name:                             "Foo Corp",
				AllowProfilesOutsideOrganization: false,
				Domains: []models.OrganizationDomain{
					models.OrganizationDomain{
						ID:     "organization_domain_id",
						Domain: "foo-corp.com",
					},
				},
			},
		},
		{
			scenario: "Request with duplicate Organization Domain returns error",
			client: &Client{
				APIKey: "test",
			},
			err: true,
			options: CreateOrganizationOpts{
				Name:    "Foo Corp",
				Domains: []string{"duplicate.com"},
			},
		},
		{
			scenario: "Idempotency Key with different event payloads returns error",
			client: &Client{
				APIKey: "test",
			},
			err: true,
			options: CreateOrganizationOpts{
				Name:           "New Corp",
				Domains:        []string{"foo-corp.com"},
				IdempotencyKey: "duplicate",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(createOrganizationTestHandler))
			defer server.Close()

			client := test.client
			client.Endpoint = server.URL
			client.HTTPClient = server.Client()

			organization, err := client.CreateOrganization(context.Background(), test.options)
			if test.err {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, test.expected, organization)
		})
	}
}

func createOrganizationTestHandler(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if auth != "Bearer test" {
		http.Error(w, "bad auth", http.StatusUnauthorized)
		return
	}

	var opts CreateOrganizationOpts
	json.NewDecoder(r.Body).Decode(&opts)
	for _, domain := range opts.Domains {
		if domain == "duplicate.com" {
			http.Error(w, "duplicate domain", http.StatusConflict)
			return
		}
	}

	if opts.IdempotencyKey == "duplicate" {
		for _, domain := range opts.Domains {
			if domain != "foo-corp.com" {
				http.Error(w, "duplicate idempotency key", http.StatusConflict)
				return
			}
		}
		if opts.Name != "Foo Corp" {
			http.Error(w, "duplicate idempotency key", http.StatusConflict)
			return
		}
	}
	if userAgent := r.Header.Get("User-Agent"); !strings.Contains(userAgent, "workos-go/") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := json.Marshal(
		models.Organization{
			ID:                               "organization_id",
			Name:                             "Foo Corp",
			AllowProfilesOutsideOrganization: false,
			Domains: []models.OrganizationDomain{
				models.OrganizationDomain{
					ID:     "organization_domain_id",
					Domain: "foo-corp.com",
				},
			},
		})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func TestUpdateOrganization(t *testing.T) {
	tests := []struct {
		scenario string
		client   *Client
		options  UpdateOrganizationOpts
		expected models.Organization
		err      bool
	}{
		{
			scenario: "Request without API Key returns an error",
			client:   &Client{},
			err:      true,
		},
		{
			scenario: "Request returns Organization with Domains",
			client: &Client{
				APIKey: "test",
			},
			options: UpdateOrganizationOpts{
				Organization: "organization_id",
				Name:         "Foo Corp",
				Domains:      []string{"foo-corp.com", "foo-corp.io"},
			},
			expected: models.Organization{
				ID:                               "organization_id",
				Name:                             "Foo Corp",
				AllowProfilesOutsideOrganization: false,
				Domains: []models.OrganizationDomain{
					models.OrganizationDomain{
						ID:     "organization_domain_id",
						Domain: "foo-corp.com",
					},
					models.OrganizationDomain{
						ID:     "organization_domain_id_2",
						Domain: "foo-corp.io",
					},
				},
			},
		},
		{
			scenario: "Request returns Organization with DomainData",
			client: &Client{
				APIKey: "test",
			},
			options: UpdateOrganizationOpts{
				Organization: "organization_id",
				Name:         "Foo Corp",
				DomainData: []models.OrganizationDomainData{
					models.OrganizationDomainData{
						Domain: "foo-corp.com",
						State:  models.OrganizationDomainDataStateVerified,
					},
					models.OrganizationDomainData{
						Domain: "foo-corp.io",
						State:  models.OrganizationDomainDataStateVerified,
					},
				},
			},
			expected: models.Organization{
				ID:                               "organization_id",
				Name:                             "Foo Corp",
				AllowProfilesOutsideOrganization: false,
				Domains: []models.OrganizationDomain{
					models.OrganizationDomain{
						ID:     "organization_domain_id",
						Domain: "foo-corp.com",
					},
					models.OrganizationDomain{
						ID:     "organization_domain_id_2",
						Domain: "foo-corp.io",
					},
				},
			},
		},
		{
			scenario: "Request with duplicate Organization Domain returns error",
			client: &Client{
				APIKey: "test",
			},
			err: true,
			options: UpdateOrganizationOpts{
				Organization: "organization_id",
				Name:         "Foo Corp",
				Domains:      []string{"duplicate.com"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(updateOrganizationTestHandler))
			defer server.Close()

			client := test.client
			client.Endpoint = server.URL
			client.HTTPClient = server.Client()

			organization, err := client.UpdateOrganization(context.Background(), test.options)
			if test.err {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, test.expected, organization)
		})
	}
}

func updateOrganizationTestHandler(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if auth != "Bearer test" {
		http.Error(w, "bad auth", http.StatusUnauthorized)
		return
	}

	var opts UpdateOrganizationOpts
	json.NewDecoder(r.Body).Decode(&opts)
	for _, domain := range opts.Domains {
		if domain == "duplicate.com" {
			http.Error(w, "duplicate domain", http.StatusConflict)
			return
		}
	}

	if userAgent := r.Header.Get("User-Agent"); !strings.Contains(userAgent, "workos-go/") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := json.Marshal(
		models.Organization{
			ID:                               "organization_id",
			Name:                             "Foo Corp",
			AllowProfilesOutsideOrganization: false,
			Domains: []models.OrganizationDomain{
				models.OrganizationDomain{
					ID:     "organization_domain_id",
					Domain: "foo-corp.com",
				},
				models.OrganizationDomain{
					ID:     "organization_domain_id_2",
					Domain: "foo-corp.io",
				},
			},
		})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
