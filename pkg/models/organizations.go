package models

// OrganizationDomain contains data about an Organization's Domains.
type OrganizationDomain struct {
	// The Organization Domain's unique identifier.
	ID string `json:"id"`

	// The domain value
	Domain string `json:"domain"`
}

// Organization contains data about a WorkOS Organization.
type Organization struct {
	// The Organization's unique identifier.
	ID string `json:"id"`

	// The Organization's name.
	Name string `json:"name"`

	// Whether Connections within the Organization allow profiles that are
	// outside of the Organization's configured User Email Domains.
	//
	// Deprecated: If you need to allow sign-ins from any email domain, contact support@workos.com.
	AllowProfilesOutsideOrganization bool `json:"allow_profiles_outside_organization"`

	// The Organization's Domains.
	Domains []OrganizationDomain `json:"domains"`

	// The timestamp of when the Organization was created.
	CreatedAt string `json:"created_at"`

	// The timestamp of when the Organization was updated.
	UpdatedAt string `json:"updated_at"`
}

type OrganizationDomainDataState string

const (
	OrganizationDomainDataStateVerified OrganizationDomainDataState = "verified"
	OrganizationDomainDataStatePending  OrganizationDomainDataState = "pending"
)

// OrganizationDomainData contains data used to create an OrganizationDomain.
type OrganizationDomainData struct {
	// The domain's value.
	Domain string `json:"domain"`

	// The domain's state.
	State OrganizationDomainDataState `json:"state"`
}
