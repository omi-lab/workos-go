package workos_errors

import (
	"fmt"
	"net/http"
	"testing"
)

func TestIsBadRequest(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "bad request",
			args: args{err: HTTPError{
				Code: http.StatusBadRequest,
			}},
			want: true,
		},
		{
			name: "internal server error",
			args: args{err: HTTPError{
				Code: http.StatusInternalServerError,
			}},
			want: false,
		},
		{
			name: "unknown error",
			args: args{err: fmt.Errorf("unknown error")},
			want: false,
		},
		{
			name: "nil",
			args: args{err: nil},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBadRequest(tt.args.err); got != tt.want {
				t.Errorf("IsBadRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAuthenticationError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "email verification required",
			args: args{err: ErrorEmailVerificationRequired{}},
			want: true,
		},
		{
			name: "mfa enrollment",
			args: args{err: ErrorMFAEnrollment{}},
			want: true,
		},
		{
			name: "mfa challenge",
			args: args{err: ErrorMFAChallenge{}},
			want: true,
		},
		{
			name: "organization selection required",
			args: args{err: ErrorOrganizationSelectionRequired{}},
			want: true,
		},
		{
			name: "sso required",
			args: args{err: ErrorSSORequired{}},
			want: true,
		},
		{
			name: "organization authentication methods required",
			args: args{err: ErrorOrganizationAuthenticationMethodsRequired{}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAuthenticationError(tt.args.err); got != tt.want {
				t.Errorf("IsAuthenticationError() = %v, want %v", got, tt.want)
			}
		})
	}
}
