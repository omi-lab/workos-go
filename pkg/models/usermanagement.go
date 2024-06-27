package models

import "github.com/omi-lab/workos-go/v4/pkg/common"

type EmailVerification struct {
	ID        string `json:"id"`
	UserId    string `json:"user_id"`
	Email     string `json:"email"`
	ExpiresAt string `json:"expires_at"`
	Code      string `json:"code"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// InvitationState represents the state of an Invitation.
type InvitationState string

// Constants that enumerate the state of an Invitation.
const (
	InvitationStatePending  InvitationState = "pending"
	InvitationStateAccepted InvitationState = "accepted"
	InvitationStateExpired  InvitationState = "expired"
	InvitationStateRevoked  InvitationState = "revoked"
)

type Invitation struct {
	ID                  string          `json:"id"`
	Email               string          `json:"email"`
	State               InvitationState `json:"state"`
	AcceptedAt          string          `json:"accepted_at,omitempty"`
	RevokedAt           string          `json:"revoked_at,omitempty"`
	Token               string          `json:"token"`
	AcceptInvitationUrl string          `json:"accept_invitation_url"`
	OrganizationID      string          `json:"organization_id,omitempty"`
	InviterUserID       string          `json:"inviter_user_id,omitempty"`
	ExpiresAt           string          `json:"expires_at"`
	CreatedAt           string          `json:"created_at"`
	UpdatedAt           string          `json:"updated_at"`
}

type MagicAuth struct {
	ID        string `json:"id"`
	UserId    string `json:"user_id"`
	Email     string `json:"email"`
	ExpiresAt string `json:"expires_at"`
	Code      string `json:"code"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type PasswordReset struct {
	ID                 string `json:"id"`
	UserId             string `json:"user_id"`
	Email              string `json:"email"`
	PasswordResetToken string `json:"password_reset_token"`
	PasswordResetUrl   string `json:"password_reset_url"`
	ExpiresAt          string `json:"expires_at"`
	CreatedAt          string `json:"created_at"`
}

// OrganizationMembershipStatus represents the status of an Organization Membership.
type OrganizationMembershipStatus string

// Constants that enumerate the status of an Organization Membership.
const (
	OrganizationMembershipStatusActive                        OrganizationMembershipStatus = "active"
	OrganizationMembershipStatusInactive                      OrganizationMembershipStatus = "inactive"
	OrganizationMembershipStatusPendingOrganizationMembership OrganizationMembershipStatus = "pending"
)

// OrganizationMembership contains data about a particular OrganizationMembership.
type OrganizationMembership struct {
	// The Organization Membership's unique identifier.
	ID string `json:"id"`

	// The ID of the User.
	UserID string `json:"user_id"`

	// The ID of the Organization.
	OrganizationID string `json:"organization_id"`

	// The role given to this Organization Membership
	Role common.RoleResponse `json:"role"`

	// The Status of the Organization.
	Status OrganizationMembershipStatus `json:"status"`

	// CreatedAt is the timestamp of when the OrganizationMembership was created.
	CreatedAt string `json:"created_at"`

	// UpdatedAt is the timestamp of when the OrganizationMembership was updated.
	UpdatedAt string `json:"updated_at"`
}

// User contains data about a particular User.
type User struct {

	// The User's unique identifier.
	ID string `json:"id"`

	// The User's first name.
	FirstName string `json:"first_name"`

	// The User's last name.
	LastName string `json:"last_name"`

	// The User's email.
	Email string `json:"email"`

	// The timestamp of when the User was created.
	CreatedAt string `json:"created_at"`

	// The timestamp of when the User was updated.
	UpdatedAt string `json:"updated_at"`

	// Whether the User email is verified.
	EmailVerified bool `json:"email_verified"`

	// A URL reference to an image representing the User.
	ProfilePictureURL string `json:"profile_picture_url"`
}

// Represents User identities obtained from external identity providers.
type Identity struct {
	// The unique ID of the user in the external identity provider.
	IdpID string `json:"idp_id"`
	// The type of the identity.
	Type string `json:"type"`
	// The type of OAuth provider for the identity.
	Provider string `json:"provider"`
}
