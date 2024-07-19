package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/omi-lab/workos-go/v4/pkg/common"
)

// UserEmail contains data about a Directory User's e-mail address.
type DirectoryUserEmail struct {
	// Flag to indicate if this e-mail is primary.
	Primary bool

	// Directory User's e-mail.
	Value string

	// Type of e-mail (ex. work).
	Type string
}

// UserGroup contains data about a Directory User's groups.
type DirectoryUserGroup struct {
	// Description of the record.
	Object string

	// The Group's identifier.
	ID string

	// The Group's Name.
	Name string
}

// UserState represents the active state of a Directory User.
type DirectoryUserState string

// Constants that enumerate the state of a Directory User.
const (
	DirectoryUserStateActive   DirectoryUserState = "active"
	DirectoryUserStateInactive DirectoryUserState = "inactive"
)

// User contains data about a provisioned Directory User.
type DirectoryUser struct {
	// The User's unique identifier.
	ID string `json:"id"`

	// The User's unique identifier assigned by the Directory Provider.
	IdpID string `json:"idp_id"`

	// The identifier of the Directory the Directory User belongs to.
	DirectoryID string `json:"directory_id"`

	// The identifier for the Organization in which the Directory resides.
	OrganizationID string `json:"organization_id"`

	// The User's username.
	Username string `json:"username"`

	// The User's e-mails.
	Emails []DirectoryUserEmail `json:"emails"`

	// The User's groups.
	Groups []DirectoryUserGroup `json:"groups"`

	// The User's first name.
	FirstName string `json:"first_name"`

	// The User's last name.
	LastName string `json:"last_name"`

	// The User's job title.
	JobTitle string `json:"job_title"`

	// The User's state.
	State DirectoryUserState `json:"state"`

	// The User's raw attributes in raw encoded JSON.
	RawAttributes json.RawMessage `json:"raw_attributes"`

	// The User's custom attributes in raw encoded JSON.
	CustomAttributes json.RawMessage `json:"custom_attributes"`

	// The User's created at date
	CreatedAt time.Time `json:"created_at"`

	// The User's updated at date
	UpdatedAt time.Time `json:"updated_at"`

	// The role given to this Directory User
	Role common.RoleResponse `json:"role,omitempty"`
}

// PrimaryEmail is a method for finding a user's primary email (when applicable)
func (r DirectoryUser) PrimaryEmail() (string, error) {
	for _, v := range r.Emails {
		if v.Primary {
			return v.Value, nil
		}
	}
	return "", errors.New("no primary email for this user found")
}

// Group contains data about a provisioned Directory Group.
type DirectoryGroup struct {
	// The Group's unique identifier.
	ID string `json:"id"`

	// The Group's name.
	Name string `json:"name"`

	// The Group's unique identifier assigned by the Directory Provider.
	IdpID string `json:"idp_id"`

	// The identifier of the Directory the group belongs to.
	DirectoryID string `json:"directory_id"`

	// The identifier for the Organization in which the Directory resides.
	OrganizationID string `json:"organization_id"`

	// The Group's created at date.
	CreatedAt time.Time `json:"created_at"`

	// The Group's updated at date.
	UpdatedAt time.Time `json:"updated_at"`

	// The Group's raw attributes in raw encoded JSON.
	RawAttributes json.RawMessage `json:"raw_attributes"`
}

// DirectoryType represents a Directory type.
type DirectoryType string

// Constants that enumerate the available Directory types.
const (
	DirectoryTypeAzureSCIMV2_0   DirectoryType = "azure scim v2.0"
	DirectoryTypeBambooHr        DirectoryType = "bamboohr"
	DirectoryTypeBreatheHr       DirectoryType = "breathe hr"
	DirectoryTypeCezanneHr       DirectoryType = "cezanne hr"
	DirectoryTypeCyberArk        DirectoryType = "cyberark scim v2.0"
	DirectoryTypeFourthHr        DirectoryType = "fourth hr"
	DirectoryTypeGSuiteDirectory DirectoryType = "gsuite directory"
	DirectoryTypeGenericSCIMV2_0 DirectoryType = "generic scim v2.0"
	DirectoryTypeHibob           DirectoryType = "hibob"
	DirectoryTypeJumpCloud       DirectoryType = "jump cloud scim v2.0"
	DirectoryTypeOktaSCIMV2_0    DirectoryType = "okta scim v2.0"
	DirectoryTypeOneLogin        DirectoryType = "onelogin scim v2.0"
	DirectoryTypePeopleHr        DirectoryType = "people hr"
	DirectoryTypePersonio        DirectoryType = "personio"
	DirectoryTypePingFederate    DirectoryType = "pingfederate scim v2.0"
	DirectoryTypeRippling        DirectoryType = "rippling scim v2.0"
	DirectoryTypeSFTP            DirectoryType = "sftp"
	DirectoryTypeSFTPWorkday     DirectoryType = "sftp workday"
	DirectoryTypeWorkday         DirectoryType = "workday"
)

// DirectoryState represents if a Directory is linked or unlinked.
type DirectoryState string

// Constants that enumerate the linked status of a Directory.
const (
	DirectoryStateLinked             DirectoryState = "linked"
	DirectoryStateUnlinked           DirectoryState = "unlinked"
	DirectoryStateInvalidCredentials DirectoryState = "invalid_credentials"
)

// Directory contains data about a project's directory.
type Directory struct {
	// Directory unique identifier.
	ID string `json:"id"`

	// Directory name.
	Name string `json:"name"`

	// Directory domain.
	Domain string `json:"domain"`

	// Externally used identifier for the Directory.
	ExternalKey string `json:"external_key"`

	// Type of the directory.
	Type DirectoryType `json:"type"`

	// Linked status for the Directory.
	State DirectoryState `json:"state"`

	// The user's directory provider's Identifier
	IdpID string `json:"idp_id"`

	// Identifier for the Directory's Organization.
	OrganizationID string `json:"organization_id"`

	// The timestamp of when the Directory was created.
	CreatedAt time.Time `json:"created_at"`

	// The timestamp of when the Directory was updated.
	UpdatedAt time.Time `json:"updated_at"`
}
