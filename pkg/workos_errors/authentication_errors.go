package workos_errors

import "github.com/omi-lab/workos-go/v4/pkg/models"

type ErrorEmailVerificationRequired struct {
	Code                       string `json:"code"`
	Message                    string `json:"message"`
	PendingAuthenticationToken string `json:"pending_authentication_token"`
	Email                      string `json:"email"`
	EmailVerificationID        string `json:"email_verification_id"`
}

func (e ErrorEmailVerificationRequired) Error() string {
	return e.Message
}

type ErrorMFAChallenge struct {
	Code                       string `json:"code"`
	Message                    string `json:"message"`
	PendingAuthenticationToken string `json:"pending_authentication_token"`
	AuthenticationFactors      []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"authentication_factors"`
	User models.User `json:"user"`
}

func (e ErrorMFAChallenge) Error() string {
	return e.Message
}

type ErrorMFAEnrollment struct {
	Code                       string      `json:"code"`
	Message                    string      `json:"message"`
	PendingAuthenticationToken string      `json:"pending_authentication_token"`
	User                       models.User `json:"user"`
}

func (e ErrorMFAEnrollment) Error() string {
	return e.Message
}

type ErrorOrganizationAuthenticationMethodsRequiredAuthMethods struct {
	AppleOAuth     bool `json:"apple_oauth"`
	GitHubOAuth    bool `json:"github_oauth"`
	GoogleOAuth    bool `json:"google_oauth"`
	MagicAuth      bool `json:"magic_auth"`
	MicrosoftOAuth bool `json:"microsoft_oauth"`
	Password       bool `json:"password"`
}

type ErrorOrganizationAuthenticationMethodsRequired struct {
	Message          string                                                    `json:"error"`
	ErrorDescription string                                                    `json:"error_description"`
	Email            string                                                    `json:"email"`
	SSOConnectionIDs []string                                                  `json:"sso_connection_ids"`
	AuthMethods      ErrorOrganizationAuthenticationMethodsRequiredAuthMethods `json:"auth_methods"`
}

func (e ErrorOrganizationAuthenticationMethodsRequired) Error() string {
	return e.Message
}

type ErrorOrganizationSelectionRequired struct {
	Code                       string `json:"code"`
	Message                    string `json:"message"`
	PendingAuthenticationToken string `json:"pending_authentication_token"`
	Organization               []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"organization"`
	User models.User `json:"user"`
}

func (e ErrorOrganizationSelectionRequired) Error() string {
	return e.Message
}

type ErrorSSORequired struct {
	Message                    string   `json:"error"`
	ErrorDescription           string   `json:"error_description"`
	Email                      string   `json:"email"`
	ConnectionIDs              []string `json:"connection_ids"`
	PendingAuthenticationToken string   `json:"pending_authentication_token"`
}

func (e ErrorSSORequired) Error() string {
	return e.Message
}

type ErrorInvalidCredentials struct {
	Message string `json:"error"`
}

func (e ErrorInvalidCredentials) Error() string {
	return e.Message
}
