package models

import "time"

type AuditLogEvent struct {
	// Represents the activity performed by the actor.
	Action string `json:"action"`

	// The schema version of the event
	Version int `json:"version,omitempty"`

	// The time when the event occurred.
	// Defaults to time.Now().
	OccurredAt time.Time `json:"occurred_at"`

	// Describes the entity that generated the event
	Actor AuditLogEventActor `json:"actor"`

	// List of event target
	Targets []AuditLogEventTarget `json:"targets"`

	// Attributes of event context
	Context AuditLogEventContext `json:"context"`

	// Event metadata.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// Context describes the event location and user agent
type AuditLogEventContext struct {
	// Place from where the event is fired
	Location string `json:"location"`

	// User Agent identity information of the event actor
	UserAgent string `json:"user_agent"`
}

// Target describes event entity's
type AuditLogEventTarget struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Type string `json:"type"`

	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// Actor describes the entity that generated the event
type AuditLogEventActor struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Type string `json:"type"`

	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// AuditLogExportState represents the active state of an AuditLogExport.
type AuditLogExportState string

// Constants that enumerate the state of a AuditLogExport.
const (
	AuditLogExportStateReady   AuditLogExportState = "Ready"
	AuditLogExportStatePending AuditLogExportState = "Pending"
	AuditLogExportStateError   AuditLogExportState = "Error"
)

type AuditLogExportObject string

const AuditLogExportObjectName AuditLogExportObject = "audit_log_export"

type AuditLogExport struct {
	// Object will always be set to 'audit_log_export'
	Object AuditLogExportObject `json:"object"`

	// AuditLogExport identifier
	ID string `json:"id"`

	// State is the active state of AuditLogExport
	State AuditLogExportState `json:"state"`

	// URL for downloading the exported logs
	URL string `json:"url"`

	// AuditLogExport's created at date
	CreatedAt string `json:"created_at"`

	// AuditLogExport's updated at date
	UpdatedAt string `json:"updated_at"`
}
