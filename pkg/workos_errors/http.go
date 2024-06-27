package workos_errors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// TryGetHTTPError returns an error when the http response contains invalid
// status code.
func TryGetHTTPError(r *http.Response) error {
	if r.StatusCode >= 200 && r.StatusCode < 300 {
		return nil
	}

	var msg string

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return HTTPError{
			Code:        r.StatusCode,
			Status:      r.Status,
			RequestID:   r.Header.Get("X-Request-ID"),
			Message:     err.Error(),
			ErrorCode:   "",
			Errors:      nil,
			FieldErrors: nil,
		}
	}

	if isJsonResponse(r) {
		return getJsonErrorMessage(body, r.StatusCode, r.Status, r.Header.Get("X-Request-ID"))
	} else {
		msg = string(body)
	}

	return HTTPError{
		Code:        r.StatusCode,
		Status:      r.Status,
		RequestID:   r.Header.Get("X-Request-ID"),
		Message:     msg,
		ErrorCode:   "",
		Errors:      nil,
		FieldErrors: nil,
	}
}

func isJsonResponse(r *http.Response) bool {
	return strings.Contains(r.Header.Get("Content-Type"), "application/json")
}

func getJsonErrorMessage(b []byte, statusCode int, status string, requestID string) error {
	if statusCode == 422 {
		var unprocesableEntityPayload struct {
			Message          string       `json:"message"`
			Error            string       `json:"error"`
			ErrorDescription string       `json:"error_description"`
			FieldErrors      []FieldError `json:"errors"`
			Code             string       `json:"code"`
		}

		if err := json.Unmarshal(b, &unprocesableEntityPayload); err != nil {
			return HTTPError{
				Code:        statusCode,
				Status:      status,
				RequestID:   requestID,
				Message:     string(b),
				ErrorCode:   "",
				Errors:      nil,
				FieldErrors: nil,
			}
		}

		return HTTPError{
			Code:        statusCode,
			Status:      status,
			RequestID:   requestID,
			Message:     unprocesableEntityPayload.Message,
			ErrorCode:   unprocesableEntityPayload.Code,
			Errors:      nil,
			FieldErrors: unprocesableEntityPayload.FieldErrors,
		}
	}

	var payload struct {
		Message          string   `json:"message"`
		Error            string   `json:"error"`
		ErrorDescription string   `json:"error_description"`
		Errors           []string `json:"errors"`
		Code             string   `json:"code"`
	}

	if err := json.Unmarshal(b, &payload); err != nil {
		return HTTPError{
			Code:        statusCode,
			Status:      status,
			RequestID:   requestID,
			Message:     string(b),
			ErrorCode:   "",
			Errors:      nil,
			FieldErrors: nil,
		}
	}

	var e error
	switch payload.Code {
	case HTTPErrorCodeEmailVerificationRequired:
		e = ErrorEmailVerificationRequired{}
	case HTTPErrorCodeMFAEnrollment:
		e = ErrorMFAEnrollment{}
	case HTTPErrorCodeMFAChallenge:
		e = ErrorMFAChallenge{}
	case HTTPErrorCodeOrganizationSelectionRequired:
		e = ErrorOrganizationSelectionRequired{}
	case HTTPErrorCodeInvalidCredentials:
		e = ErrorInvalidCredentials{}
	}

	switch payload.Error {
	case HTTPErrorCodeSSORequired:
		e = ErrorSSORequired{}
	case HTTPErrorCodeOrganizationAuthenticationMethodsRequired:
		e = ErrorOrganizationAuthenticationMethodsRequired{}
	}

	if e != nil {
		if err := json.NewDecoder(bytes.NewReader(b)).Decode(&e); err != nil {
			return HTTPError{
				Code:        statusCode,
				Status:      status,
				RequestID:   requestID,
				Message:     string(b),
				ErrorCode:   "",
				Errors:      nil,
				FieldErrors: nil,
			}
		}
		return e
	}

	if payload.Error != "" && payload.ErrorDescription != "" {
		return HTTPError{
			Code:        statusCode,
			Status:      status,
			RequestID:   requestID,
			Message:     fmt.Sprintf("%s %s", payload.Error, payload.ErrorDescription),
			ErrorCode:   "",
			Errors:      nil,
			FieldErrors: nil,
		}
	} else if payload.Message != "" && len(payload.Errors) == 0 {
		return HTTPError{
			Code:        statusCode,
			Status:      status,
			RequestID:   requestID,
			Message:     payload.Message,
			ErrorCode:   "",
			Errors:      nil,
			FieldErrors: nil,
		}
	} else if payload.Message != "" && len(payload.Errors) > 0 {
		return HTTPError{
			Code:        statusCode,
			Status:      status,
			RequestID:   requestID,
			Message:     payload.Message,
			ErrorCode:   payload.Code,
			Errors:      payload.Errors,
			FieldErrors: nil,
		}
	}

	return HTTPError{
		Code:        statusCode,
		Status:      status,
		RequestID:   requestID,
		Message:     string(b),
		ErrorCode:   "",
		Errors:      nil,
		FieldErrors: nil,
	}
}

// HTTPError represents an http error.
type HTTPError struct {
	Code        int
	Status      string
	RequestID   string
	Message     string
	ErrorCode   string
	Errors      []string
	FieldErrors []FieldError
}

type FieldError struct {
	Field string
	Code  string
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("%s: request id %q: %s", e.Status, e.RequestID, e.Message)
}
