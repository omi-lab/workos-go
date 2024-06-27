package workos_errors

import (
	"errors"
	"net/http"
)

func IsBadRequest(err error) bool {
	var httpError HTTPError
	return errors.As(err, &httpError) && httpError.Code == http.StatusBadRequest
}

func IsAuthenticationError(err error) bool {
	return errors.As(err, &ErrorEmailVerificationRequired{}) ||
		errors.As(err, &ErrorMFAChallenge{}) ||
		errors.As(err, &ErrorMFAEnrollment{}) ||
		errors.As(err, &ErrorOrganizationAuthenticationMethodsRequired{}) ||
		errors.As(err, &ErrorOrganizationSelectionRequired{}) ||
		errors.As(err, &ErrorSSORequired{}) ||
		errors.As(err, &ErrorInvalidCredentials{})
}
