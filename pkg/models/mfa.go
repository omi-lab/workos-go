package models

// Type represents the type of Authentication Factor
type FactorType string

// Constants that enumerate the available Types.
const (
	FactorTypeSMS  FactorType = "sms"
	FactorTypeTOTP FactorType = "totp"
)

type Factor struct {
	// The authentication factor's unique ID
	ID string `json:"id"`

	// The name of the response type
	Object string `json:"object"`

	// The timestamp of when the request was created.
	CreatedAt string `json:"created_at"`

	// The timestamp of when the request was updated.
	UpdatedAt string `json:"updated_at"`

	// The type of request either 'sms' or 'totp'
	Type FactorType `json:"type"`

	// Details of the totp response will be 'null' if using sms
	TOTP TOTPDetails `json:"totp"`

	// Details of the sms response will be 'null' if using totp
	SMS SMSDetails `json:"sms"`
}

type TOTPDetails struct {
	QRCode string `json:"qr_code"`
	Secret string `json:"secret"`
	URI    string `json:"uri"`
}

type SMSDetails struct {
	PhoneNumber string `json:"phone_number"`
}

type Challenge struct {
	// The authentication challenge's unique ID
	ID string `json:"id"`

	// The name of the response type.
	Object string `json:"object"`

	// The timestamp of when the request was created.
	CreatedAt string `json:"created_at"`

	// The timestamp of when the request was updated.
	UpdatedAt string `json:"updated_at"`

	// The timestamp of when the request expires.
	ExpiresAt string `json:"expires_at"`

	// The authentication factor Id used to create the request.
	FactorID string `json:"authentication_factor_id"`
}
