package models

import "time"

type WebhookEventName = string

const (
	WebhookEventNameEmailVerificationCreated WebhookEventName = "email_verification.created"
	WebhookEventNamePasswordResetCreated     WebhookEventName = "password_reset.created"
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
