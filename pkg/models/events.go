package models

import (
	"encoding/json"
	"time"
)

const (
	// Connection Events
	EventConnectionActivated   = "connection.activated"
	EventConnectionDeactivated = "connection.deactived"
	EventConnectionDeleted     = "connection.deleted"
	// Directory Events
	EventDirectoryActivated = "dsync.activated"
	EventDirectoryDeleted   = "dsync.deleted"
	// Directory User Events
	EventDirectoryUserCreated = "dsync.user.created"
	EventDirectoryUserUpdated = "dsync.user.updated"
	EventDirectoryUserDeleted = "dsync.user.deleted"
	// Directory Group Events
	EventDirectoryGroupCreated     = "dsync.group.created"
	EventDirectoryGroupUpdated     = "dsync.group.updated"
	EventDirectoryGroupDeleted     = "dsync.group.deleted"
	EventDirectoryGroupUserAdded   = "dsync.group.user_added"
	EventDirectroyGroupUserRemoved = "dsync.group.user_removed"
	// User Management Events
	EventUserCreated                   = "user.created"
	EventUserUpdated                   = "user.updated"
	EventUserDeleted                   = "user.deleted"
	EventOrganizationMembershipAdded   = "organization_membership.added" // Deprecated: use OrganizationMembershipCreated instead
	EventOrganizationMembershipCreated = "organization_membership.created"
	EventOrganizationMembershipDeleted = "organization_membership.deleted"
	EventOrganizationMembershipUpdated = "organization_membership.updated"
	EventOrganizationMembershipRemoved = "organization_membership.removed" // Deprecated: use OrganizationMembershipDeleted instead
	EventSessionCreated                = "session.created"
	EventEmailVerificationCreated      = "email_verification.created"
	EventInvitationCreated             = "invitation.created"
	EventMagicAuthCreated              = "magic_auth.created"
	EventPasswordResetCreated          = "password_reset.created"
)

// Event contains data about a particular Event.
type Event struct {
	// The Event's unique identifier.
	ID string `json:"id"`

	// The type of Event.
	Event string `json:"event"`

	// The Event's data in raw encoded JSON.
	Data json.RawMessage `json:"data"`

	// The Event's created at date.
	CreatedAt time.Time `json:"created_at"`
}
