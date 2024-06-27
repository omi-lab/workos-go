package models

// GenerateLinkIntent represents the intent of an Admin Portal.
type GenerateLinkIntent string

// Constants that enumerate the available GenerateLinkIntent types.
const (
	GenerateLinkIntentSSO        GenerateLinkIntent = "sso"
	GenerateLinkIntentDSync      GenerateLinkIntent = "dsync"
	GenerateLinkIntentAuditLogs  GenerateLinkIntent = "audit_logs"
	GenerateLinkIntentLogStreams GenerateLinkIntent = "log_streams"
)
