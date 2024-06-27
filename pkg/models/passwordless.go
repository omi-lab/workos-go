package models

// PasswordlessSession contains data about a WorkOS Passwordless Session.
type PasswordlessSession struct {
	// The Passwordless Session's unique identifier.
	ID string `json:"id"`

	// The email of the user to authenticate.
	Email string `json:"email"`

	// ISO-8601 datetime at which the Passwordless Session link expires.
	ExpiresAt string `json:"expires_at"`

	// The link for the user to authenticate with.
	Link string `json:"link"`
}

// CreateSessionType represents the type of a Passwordless Session.
type PasswordlessSessionType string

// Constants that enumerate the available PasswordlessSessionType values.
const (
	PasswordlessSessionTypeMagicLink PasswordlessSessionType = "MagicLink"
)
