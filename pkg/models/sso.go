package models

// ConnectionType represents a connection type.
type ConnectionType string

// Constants that enumerate the available connection types.
const (
	ConnectionTypeADFSSAML              ConnectionType = "ADFSSAML"
	ConnectionTypeAdpOidc               ConnectionType = "AdpOidc"
	ConnectionTypeAuth0SAML             ConnectionType = "Auth0SAML"
	ConnectionTypeAzureSAML             ConnectionType = "AzureSAML"
	ConnectionTypeCasSAML               ConnectionType = "CasSAML"
	ConnectionTypeCloudflareSAML        ConnectionType = "CloudflareSAML"
	ConnectionTypeClassLinkSAML         ConnectionType = "ClassLinkSAML"
	ConnectionTypeCyberArkSAML          ConnectionType = "CyberArkSAML"
	ConnectionTypeDuoSAML               ConnectionType = "DuoSAML"
	ConnectionTypeGenericOIDC           ConnectionType = "GenericOIDC"
	ConnectionTypeGenericSAML           ConnectionType = "GenericSAML"
	ConnectionTypeGoogleOAuth           ConnectionType = "GoogleOAuth"
	ConnectionTypeGoogleSAML            ConnectionType = "GoogleSAML"
	ConnectionTypeJumpCloudSAML         ConnectionType = "JumpCloudSAML"
	ConnectionTypeKeycloakSAML          ConnectionType = "KeycloakSAML"
	ConnectionTypeLastPassSAML          ConnectionType = "LastPassSAML"
	ConnectionTypeLoginGovOidc          ConnectionType = "LoginGovOidc"
	ConnectionTypeMagicLink             ConnectionType = "MagicLink"
	ConnectionTypeMicrosoftOAuth        ConnectionType = "MicrosoftOAuth"
	ConnectionTypeMiniOrangeSAML        ConnectionType = "MiniOrangeSAML"
	ConnectionTypeNetIqSAML             ConnectionType = "NetIqSAML"
	ConnectionTypeOktaSAML              ConnectionType = "OktaSAML"
	ConnectionTypeOneLoginSAML          ConnectionType = "OneLoginSAML"
	ConnectionTypeOracleSAML            ConnectionType = "OracleSAML"
	ConnectionTypePingFederateSAML      ConnectionType = "PingFederateSAML"
	ConnectionTypePingOneSAML           ConnectionType = "PingOneSAML"
	ConnectionTypeRipplingSAML          ConnectionType = "RipplingSAML"
	ConnectionTypeSalesforceSAML        ConnectionType = "SalesforceSAML"
	ConnectionTypeShibbolethSAML        ConnectionType = "ShibbolethSAML"
	ConnectionTypeShibbolethGenericSAML ConnectionType = "ShibbolethGenericSAML"
	ConnectionTypeSimpleSamlPhpSAML     ConnectionType = "SimpleSamlPhpSAML"
	ConnectionTypeVMwareSAML            ConnectionType = "VMwareSAML"
)

// ConnectionDomain represents the domain records associated with a Connection.
type ConnectionDomain struct {
	// Connection Domain unique identifier.
	ID string `json:"id"`

	// Domain for a Connection record.
	Domain string `json:"domain"`
}

// ConnectionStatus represents a Connection's linked status.
//
// Deprecated: Please use ConnectionState instead.
type ConnectionStatus string

// Constants that enumerate the available Connection's linked statuses.
const (
	ConnectionStatusLinked   ConnectionStatus = "linked"
	ConnectionStatusUnlinked ConnectionStatus = "unlinked"
)

// ConnectionState indicates whether a Connection is able to authenticate users.
type ConnectionState string

// Constants that enumerate a Connection's possible states.
const (
	ConnectionStateDraft      ConnectionState = "draft"
	ConnectionStateActive     ConnectionState = "active"
	ConnectionStateInactive   ConnectionState = "inactive"
	ConnectionStateValidating ConnectionState = "validating"
)

// Connection represents a Connection record.
type Connection struct {
	// Connection unique identifier.
	ID string `json:"id"`

	// Connection linked status. Deprecated; use State instead.
	Status ConnectionStatus `json:"status"`

	// Connection linked state.
	State ConnectionState `json:"state"`

	// Connection name.
	Name string `json:"name"`

	// Connection provider type.
	ConnectionType ConnectionType `json:"connection_type"`

	// Organization ID.
	OrganizationID string `json:"organization_id"`

	// Domain records for the Connection.
	Domains []ConnectionDomain `json:"domains"`

	// The timestamp of when the Connection was created.
	CreatedAt string `json:"created_at"`

	// The timestamp of when the Connection was updated.
	UpdatedAt string `json:"updated_at"`
}
