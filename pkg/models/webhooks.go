package models

import "time"

type WebhookEventName = string

const (
	WebhookEventNameEmailVerificationCreated WebhookEventName = "email_verification.created"
	WebhookEventNamePasswordResetCreated     WebhookEventName = "password_reset.created"
	WebhookEventNameConnectionActivated      WebhookEventName = "connection.activated"
	WebhookEventNameConnectionDeactivated    WebhookEventName = "connection.deactivated"
	WebhookEventNameConnectionDeleted        WebhookEventName = "connection.deleted"
	WebhookEventNameInvitationCreated        WebhookEventName = "invitations.created"
)

type WebhookEvent struct {
	ID        string           `json:"id"`
	Event     WebhookEventName `json:"event"`
	CreatedAt time.Time        `json:"createdAt"`
}

type WebhookEventEmailVerificationCreated struct {
	WebhookEvent
	Data struct {
		Object    string    `json:"object"`
		ID        string    `json:"id"`
		Email     string    `json:"email"`
		UserID    string    `json:"user_id"`
		ExpiresAt time.Time `json:"expires_at"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
}

type WebhookEventPasswordResetCreated struct {
	WebhookEvent
	Data struct {
		Object    string    `json:"object"`
		ID        string    `json:"id"`
		Email     string    `json:"email"`
		UserID    string    `json:"user_id"`
		ExpiresAt time.Time `json:"expires_at"`
		CreatedAt time.Time `json:"created_at"`
	}
}

type WebhookEventConnectionDomain struct {
	Object string `json:"object"`
	ID     string `json:"id"`
	Domain string `json:"domain"`
}

type WebhookEventConnection struct {
	WebhookEvent
	Data struct {
		Object         string                         `json:"object"`
		ID             string                         `json:"id"`
		OrganizationID string                         `json:"organization_id"`
		ExternalKey    string                         `json:"external_key"`
		State          ConnectionState                `json:"state"`
		Status         ConnectionStatus               `json:"status"`
		Domains        []WebhookEventConnectionDomain `json:"domains"`
		ExpiresAt      time.Time                      `json:"expires_at"`
		CreatedAt      time.Time                      `json:"created_at"`
	}
}

type WebhookEventInvitationCreated struct {
	WebhookEvent
	Data struct {
		Object         string          `json:"object"`
		ID             string          `json:"id"`
		Email          string          `json:"email"`
		State          InvitationState `json:"state"`
		OrganizationID string          `json:"organization_id"`
		InviterUserID  string          `json:"inviter_user_id"`
		AcceptedAt     *time.Time      `json:"accepted_at"`
		RevokedAt      *time.Time      `json:"revoked_at"`
		ExpiresAt      time.Time       `json:"expires_at"`
		CreatedAt      time.Time       `json:"created_at"`
		UpdatedAt      time.Time       `json:"updated_at"`
	}
}
