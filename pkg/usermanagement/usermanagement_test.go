package usermanagement

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/omi-lab/workos-go/v4/pkg/common"
	"github.com/omi-lab/workos-go/v4/pkg/models"

	"github.com/stretchr/testify/require"
)

func mockClient(s *httptest.Server) *Client {
	client := NewClient("")
	client.HTTPClient = s.Client()
	client.Endpoint = s.URL
	return client
}

func TestUserManagementGetUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(getUserTestHandler))
	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	expectedResponse := models.User{
		ID:            "user_01E3JC5F5Z1YJNPGVYWV9SX6GH",
		Email:         "marcelina@foo-corp.com",
		FirstName:     "Marcelina",
		LastName:      "Davis",
		EmailVerified: true,
		CreatedAt:     time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
		UpdatedAt:     time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
	}

	userRes, err := GetUser(context.Background(), GetUserOpts{
		User: "user_123",
	})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, userRes)
}

func TestUserManagementListUsers(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(listUsersTestHandler))

	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	expectedResponse := ListUsersResponse{
		Data: []models.User{
			{
				ID:            "user_01E3JC5F5Z1YJNPGVYWV9SX6GH",
				Email:         "marcelina@foo-corp.com",
				FirstName:     "Marcelina",
				LastName:      "Davis",
				EmailVerified: true,
				CreatedAt:     time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
				UpdatedAt:     time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
			},
		},
		ListMetadata: common.ListMetadata{
			After: "",
		},
	}

	userRes, err := ListUsers(context.Background(), ListUsersOpts{})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, userRes)
}

func TestUserManagementCreateUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(createUserTestHandler))
	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	expectedResponse := models.User{
		ID:            "user_01E3JC5F5Z1YJNPGVYWV9SX6GH",
		Email:         "marcelina@foo-corp.com",
		FirstName:     "Marcelina",
		LastName:      "Davis",
		EmailVerified: true,
		CreatedAt:     time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
		UpdatedAt:     time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
	}

	userRes, err := CreateUser(context.Background(), CreateUserOpts{
		Email:         "marcelina@gmail.com",
		FirstName:     "Marcelina",
		LastName:      "Davis",
		EmailVerified: true,
		Password:      "pass",
	})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, userRes)
}

func TestUserManagementCreateUserPasswordHash(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(createUserTestHandler))
	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	expectedResponse := models.User{
		ID:            "user_01E3JC5F5Z1YJNPGVYWV9SX6GH",
		Email:         "marcelina@foo-corp.com",
		FirstName:     "Marcelina",
		LastName:      "Davis",
		EmailVerified: true,
		CreatedAt:     time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
		UpdatedAt:     time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
	}

	userRes, err := CreateUser(context.Background(), CreateUserOpts{
		Email:            "marcelina@gmail.com",
		FirstName:        "Marcelina",
		LastName:         "Davis",
		EmailVerified:    true,
		PasswordHash:     "$2b$10$dXS6RadWKYIqs6vOwqKZceLuCIqz6S81t06.yOkGJbbfeO9go4fai",
		PasswordHashType: "bcrypt",
	})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, userRes)
}

func TestUserManagementUpdateUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(updateUserTestHandler))
	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	expectedResponse := models.User{
		ID:            "user_01E3JC5F5Z1YJNPGVYWV9SX6GH",
		Email:         "marcelina@foo-corp.com",
		FirstName:     "Marcelina",
		LastName:      "Davis",
		EmailVerified: true,
		CreatedAt:     time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
		UpdatedAt:     time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
	}

	userRes, err := UpdateUser(context.Background(), UpdateUserOpts{
		User:          "user_01E3JC5F5Z1YJNPGVYWV9SX6GH",
		FirstName:     "Marcelina",
		LastName:      "Davis",
		EmailVerified: true,
		Password:      "pass",
	})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, userRes)
}

func TestUserManagementUpdateUserPasswordHash(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(updateUserTestHandler))
	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	expectedResponse := models.User{
		ID:            "user_01E3JC5F5Z1YJNPGVYWV9SX6GH",
		Email:         "marcelina@foo-corp.com",
		FirstName:     "Marcelina",
		LastName:      "Davis",
		EmailVerified: true,
		CreatedAt:     time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
		UpdatedAt:     time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
	}

	userRes, err := UpdateUser(context.Background(), UpdateUserOpts{
		User:             "user_01E3JC5F5Z1YJNPGVYWV9SX6GH",
		FirstName:        "Marcelina",
		LastName:         "Davis",
		EmailVerified:    true,
		PasswordHash:     "$2b$10$dXS6RadWKYIqs6vOwqKZceLuCIqz6S81t06.yOkGJbbfeO9go4fai",
		PasswordHashType: "bcrypt",
	})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, userRes)
}

func TestUsersDeleteUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(deleteUserTestHandler))
	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	err := DeleteUser(context.Background(), DeleteUserOpts{
		User: "user_01E3JC5F5Z1YJNPGVYWV9SX6GH",
	})

	require.NoError(t, err)
	require.Equal(t, nil, err)
}

func TestUsersGetEmailVerification(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(getEmailVerificationTestHandler))
	defer server.Close()

	DefaultClient = mockClient(server)
	SetAPIKey("test")

	expectedResponse := models.EmailVerification{
		ID:        "email_verification_123",
		UserId:    "user_123",
		Email:     "marcelina@foo-corp.com",
		ExpiresAt: time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
		Code:      "123456",
		CreatedAt: time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
		UpdatedAt: time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
	}

	getByIDRes, err := GetEmailVerification(context.Background(), GetEmailVerificationOpts{
		EmailVerification: "email_verification_123",
	})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, getByIDRes)
}

func TestUsersSendVerificationEmail(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(sendVerificationEmailTestHandler))
	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	expectedResponse := UserResponse{
		User: models.User{
			ID:            "user_123",
			Email:         "marcelina@foo-corp.com",
			FirstName:     "Marcelina",
			LastName:      "Davis",
			EmailVerified: true,
			CreatedAt:     time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
			UpdatedAt:     time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
		},
	}

	userRes, err := SendVerificationEmail(context.Background(), SendVerificationEmailOpts{
		User: "user_123",
	})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, userRes)
}

func TestUserManagementVerifyEmail(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(verifyEmailCodeTestHandler))
	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	expectedResponse := UserResponse{
		User: models.User{
			ID:            "user_123",
			Email:         "marcelina@foo-corp.com",
			FirstName:     "Marcelina",
			LastName:      "Davis",
			EmailVerified: true,
		},
	}

	userRes, err := VerifyEmail(context.Background(), VerifyEmailOpts{
		User: "user_123",
		Code: "testToken",
	})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, userRes)
}

func TestUsersGetPasswordReset(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(getPasswordResetTestHandler))
	defer server.Close()

	DefaultClient = mockClient(server)
	SetAPIKey("test")

	expectedResponse := models.PasswordReset{
		ID:                 "password_reset_123",
		UserId:             "user_123",
		Email:              "marcelina@foo-corp.com",
		PasswordResetToken: "myToken",
		PasswordResetUrl:   "https://your-app.com/reset-password?token=myToken",
		ExpiresAt:          time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
		CreatedAt:          time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
	}

	getByIDRes, err := GetPasswordReset(context.Background(), GetPasswordResetOpts{
		PasswordReset: "password_reset_123",
	})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, getByIDRes)
}

func TestUserManagementCreatePasswordReset(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(CreatePasswordResetTestHandler))

	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	expectedResponse := models.PasswordReset{
		ID:                 "password_reset_123",
		UserId:             "user_123",
		Email:              "marcelina@foo-corp.com",
		PasswordResetToken: "myToken",
		PasswordResetUrl:   "https://your-app.com/reset-password?token=myToken",
		ExpiresAt:          time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
		CreatedAt:          time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
	}

	createRes, err := CreatePasswordReset(context.Background(), CreatePasswordResetOpts{
		Email: "marcelina@foo-corp.com",
	})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, createRes)
}

func TestUserManagementCreatePasswordResetChallenge(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(sendPasswordResetEmailTestHandler))
	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	err := SendPasswordResetEmail(context.Background(), SendPasswordResetEmailOpts{
		Email:            "marcelina@foo-corp.com",
		PasswordResetUrl: "https://example.com/reset",
	})

	require.NoError(t, err)
}

func TestUserManagementResetPassword(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(resetPasswordHandler))
	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	expectedResponse := UserResponse{
		User: models.User{
			ID: "user_123",

			Email:         "marcelina@foo-corp.com",
			FirstName:     "Marcelina",
			LastName:      "Davis",
			EmailVerified: true,
		},
	}

	userRes, err := ResetPassword(context.Background(), ResetPasswordOpts{
		Token: "testToken",
	})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, userRes)
}

func TestUserManagementGetAuthorizationURL(t *testing.T) {
	client := NewClient("test")

	u, err := client.GetAuthorizationURL(GetAuthorizationURLOpts{
		ClientID:    "client_123",
		Provider:    "GoogleOAuth",
		RedirectURI: "https://example.com/sso/workos/callback",
		State:       "custom state",
	})
	require.NoError(t, err)
	require.Equal(
		t,
		"https://api.workos.com/user_management/authorize?client_id=client_123&provider=GoogleOAuth&redirect_uri=https%3A%2F%2Fexample.com%2Fsso%2Fworkos%2Fcallback&response_type=code&state=custom+state",
		u.String(),
	)
}

func TestUserManagementAuthenticateWithCode(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(authenticationResponseTestHandler))

	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	expectedResponse := AuthenticateResponse{
		User: models.User{
			ID:        "testUserID",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "employee@foo-corp.com",
		},
		OrganizationID: "org_123",
		AccessToken:    "access_token",
		RefreshToken:   "refresh_token",
	}

	authenticationRes, err := AuthenticateWithCode(context.Background(), AuthenticateWithCodeOpts{})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, authenticationRes)
}

func TestUserManagementAuthenticateWithPassword(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(authenticationResponseTestHandler))

	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	expectedResponse := AuthenticateResponse{
		User: models.User{
			ID:        "testUserID",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "employee@foo-corp.com",
		},
		OrganizationID: "org_123",
		AccessToken:    "access_token",
		RefreshToken:   "refresh_token",
	}

	authenticationRes, err := AuthenticateWithPassword(context.Background(), AuthenticateWithPasswordOpts{})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, authenticationRes)
}

func TestUserManagementAuthenticateWithMagicAuth(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(authenticationResponseTestHandler))

	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	expectedResponse := AuthenticateResponse{
		User: models.User{
			ID:        "testUserID",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "employee@foo-corp.com",
		},
		OrganizationID: "org_123",
		AccessToken:    "access_token",
		RefreshToken:   "refresh_token",
	}

	authenticationRes, err := AuthenticateWithMagicAuth(context.Background(), AuthenticateWithMagicAuthOpts{})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, authenticationRes)
}

func TestUserManagementAuthenticateWithTOTP(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(authenticationResponseTestHandler))

	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	expectedResponse := AuthenticateResponse{
		User: models.User{
			ID:        "testUserID",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "employee@foo-corp.com",
		},
		OrganizationID: "org_123",
		AccessToken:    "access_token",
		RefreshToken:   "refresh_token",
	}

	authenticationRes, err := AuthenticateWithTOTP(context.Background(), AuthenticateWithTOTPOpts{})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, authenticationRes)
}

func TestUserManagementAuthenticateWithEmailVerificationCode(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(authenticationResponseTestHandler))

	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	expectedResponse := AuthenticateResponse{
		User: models.User{
			ID:        "testUserID",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "employee@foo-corp.com",
		},
		OrganizationID: "org_123",
		AccessToken:    "access_token",
		RefreshToken:   "refresh_token",
	}

	authenticationRes, err := AuthenticateWithEmailVerificationCode(context.Background(), AuthenticateWithEmailVerificationCodeOpts{})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, authenticationRes)
}

func TestUserManagementAuthenticateWithOrganizationSelection(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(authenticationResponseTestHandler))

	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	expectedResponse := AuthenticateResponse{
		User: models.User{
			ID:        "testUserID",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "employee@foo-corp.com",
		},
		OrganizationID: "org_123",
		AccessToken:    "access_token",
		RefreshToken:   "refresh_token",
	}

	authenticationRes, err := AuthenticateWithOrganizationSelection(context.Background(), AuthenticateWithOrganizationSelectionOpts{})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, authenticationRes)
}

func TestUsersGetMagicAuth(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(getMagicAuthTestHandler))
	defer server.Close()

	DefaultClient = mockClient(server)
	SetAPIKey("test")

	expectedResponse := models.MagicAuth{
		ID:        "magic_auth_123",
		UserId:    "user_123",
		Email:     "marcelina@foo-corp.com",
		ExpiresAt: time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
		Code:      "123456",
		CreatedAt: time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
		UpdatedAt: time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
	}

	getByIDRes, err := GetMagicAuth(context.Background(), GetMagicAuthOpts{
		MagicAuth: "magic_auth_123",
	})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, getByIDRes)
}

func TestUserManagementCreateMagicAuth(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(CreateMagicAuthTestHandler))

	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	expectedResponse := models.MagicAuth{
		ID:        "magic_auth_123",
		UserId:    "user_123",
		Email:     "marcelina@foo-corp.com",
		ExpiresAt: time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
		Code:      "123456",
		CreatedAt: time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
		UpdatedAt: time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
	}

	createRes, err := CreateMagicAuth(context.Background(), CreateMagicAuthOpts{
		Email: "marcelina@foo-corp.com",
	})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, createRes)
}

func TestUserManagementSendMagicAuthCode(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(sendMagicAuthCodeTestHandler))

	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	err := SendMagicAuthCode(context.Background(), SendMagicAuthCodeOpts{})

	require.NoError(t, err)
}

func TestUserManagementEnrollAuthFactor(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(enrollAuthFactorTestHandler))

	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	expectedResponse := EnrollAuthFactorResponse{
		Factor: models.Factor{
			ID:        "auth_factor_test123",
			CreatedAt: time.Date(2022, 2, 17, 22, 39, 26, 616, time.UTC),
			UpdatedAt: time.Date(2022, 2, 17, 22, 39, 26, 616, time.UTC),
			Type:      "generic_otp",
		},
		Challenge: models.Challenge{
			ID:        "auth_challenge_test123",
			CreatedAt: time.Date(2022, 2, 17, 22, 39, 26, 616, time.UTC),
			UpdatedAt: time.Date(2022, 2, 17, 22, 39, 26, 616, time.UTC),
			FactorID:  "auth_factor_test123",
			ExpiresAt: time.Date(2022, 2, 17, 22, 39, 26, 616, time.UTC),
		},
	}

	authenticationRes, err := EnrollAuthFactor(context.Background(), EnrollAuthFactorOpts{
		User:       "user_01E3JC5F5Z1YJNPGVYWV9SX6GH",
		Type:       models.FactorTypeTOTP,
		TOTPSecret: "testSecret",
	})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, authenticationRes)
}

func TestUserManagementListAuthFactors(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(listAuthFactorsTestHandler))

	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	expectedResponse := ListAuthFactorsResponse{
		Data: []models.Factor{
			{
				ID:        "auth_factor_test123",
				CreatedAt: time.Date(2022, 2, 17, 22, 39, 26, 616, time.UTC),
				UpdatedAt: time.Date(2022, 2, 17, 22, 39, 26, 616, time.UTC),
				Type:      "generic_otp",
			},
			{
				ID:        "auth_factor_test234",
				CreatedAt: time.Date(2022, 2, 17, 22, 39, 26, 616, time.UTC),
				UpdatedAt: time.Date(2022, 2, 17, 22, 39, 26, 616, time.UTC),
				Type:      "generic_otp",
			},
		},
	}

	authenticationRes, err := ListAuthFactors(context.Background(), ListAuthFactorsOpts{
		User: "user_01E3JC5F5Z1YJNPGVYWV9SX6GH",
	})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, authenticationRes)
}

func TestUserManagementGetOrganizationMembership(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(getOrganizationMembershipTestHandler))
	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	expectedResponse := models.OrganizationMembership{
		ID:             "om_01E4ZCR3C56J083X43JQXF3JK5",
		UserID:         "user_01E4ZCR3C5A4QZ2Z2JQXGKZJ9E",
		OrganizationID: "org_01E4ZCR3C56J083X43JQXF3JK5",
		Status:         models.OrganizationMembershipStatusActive,
		CreatedAt:      time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
		UpdatedAt:      time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
	}

	userRes, err := GetOrganizationMembership(context.Background(), GetOrganizationMembershipOpts{
		OrganizationMembership: "om_01E4ZCR3C56J083X43JQXF3JK5",
	})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, userRes)
}

func TestUserManagementListOrganizationMemberships(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(listOrganizationMembershipsTestHandler))

	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	expectedResponse := ListOrganizationMembershipsResponse{
		Data: []models.OrganizationMembership{
			{
				ID:             "om_01E4ZCR3C56J083X43JQXF3JK5",
				UserID:         "user_01E4ZCR3C5A4QZ2Z2JQXGKZJ9E",
				OrganizationID: "org_01E4ZCR3C56J083X43JQXF3JK5",
				Status:         models.OrganizationMembershipStatusActive,
				CreatedAt:      time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
				UpdatedAt:      time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
			},
		},
		ListMetadata: common.ListMetadata{
			After: "",
		},
	}

	userRes, err := ListOrganizationMemberships(context.Background(), ListOrganizationMembershipsOpts{})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, userRes)
}

func TestUserManagementCreateOrganizationMembership(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(createOrganizationMembershipTestHandler))
	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	expectedRole := common.RoleResponse{
		Slug: "member",
	}

	expectedResponse := models.OrganizationMembership{
		ID:             "om_01E4ZCR3C56J083X43JQXF3JK5",
		UserID:         "user_01E4ZCR3C5A4QZ2Z2JQXGKZJ9E",
		OrganizationID: "org_01E4ZCR3C56J083X43JQXF3JK5",
		Status:         models.OrganizationMembershipStatusActive,
		Role:           expectedRole,
		CreatedAt:      time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
		UpdatedAt:      time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
	}

	userRes, err := CreateOrganizationMembership(context.Background(), CreateOrganizationMembershipOpts{
		UserID:         "user_01E4ZCR3C5A4QZ2Z2JQXGKZJ9E",
		OrganizationID: "org_01E4ZCR3C56J083X43JQXF3JK5",
	})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, userRes)
}

func TestUsersDeleteOrganizationMembership(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(deleteOrganizationMembershipTestHandler))
	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	err := DeleteOrganizationMembership(context.Background(), DeleteOrganizationMembershipOpts{
		OrganizationMembership: "om_01E4ZCR3C56J083X43JQXF3JK5",
	})

	require.NoError(t, err)
	require.Equal(t, nil, err)
}

func TestUsersUpdateOrganizationMembership(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(updateOrganizationMembershipTestHandler))
	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	expectedRole := common.RoleResponse{
		Slug: "member",
	}

	expectedResponse := models.OrganizationMembership{
		ID:             "om_01E4ZCR3C56J083X43JQXF3JK5",
		UserID:         "user_01E4ZCR3C5A4QZ2Z2JQXGKZJ9E",
		OrganizationID: "org_01E4ZCR3C56J083X43JQXF3JK5",
		Status:         models.OrganizationMembershipStatusActive,
		Role:           expectedRole,
		CreatedAt:      time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
		UpdatedAt:      time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
	}

	body, err := UpdateOrganizationMembership(
		context.Background(),
		"om_01E4ZCR3C56J083X43JQXF3JK5",
		UpdateOrganizationMembershipOpts{
			RoleSlug: "member",
		},
	)

	require.NoError(t, err)
	require.Equal(t, expectedResponse, body)
}

func TestUserManagementDeactivateOrganizationMembership(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(deactivateOrganizationMembershipTestHandler))
	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	expectedResponse := models.OrganizationMembership{
		ID:             "om_01E4ZCR3C56J083X43JQXF3JK5",
		UserID:         "user_01E4ZCR3C5A4QZ2Z2JQXGKZJ9E",
		OrganizationID: "org_01E4ZCR3C56J083X43JQXF3JK5",
		Status:         models.OrganizationMembershipStatusInactive,
		CreatedAt:      time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
		UpdatedAt:      time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
	}

	userRes, err := DeactivateOrganizationMembership(context.Background(), DeactivateOrganizationMembershipOpts{
		OrganizationMembership: "om_01E4ZCR3C56J083X43JQXF3JK5",
	})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, userRes)
}

func TestUserManagementReactivateOrganizationMembership(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(reactivateOrganizationMembershipTestHandler))
	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	expectedResponse := models.OrganizationMembership{
		ID:             "om_01E4ZCR3C56J083X43JQXF3JK5",
		UserID:         "user_01E4ZCR3C5A4QZ2Z2JQXGKZJ9E",
		OrganizationID: "org_01E4ZCR3C56J083X43JQXF3JK5",
		Status:         models.OrganizationMembershipStatusActive,
		CreatedAt:      time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
		UpdatedAt:      time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
	}

	userRes, err := ReactivateOrganizationMembership(context.Background(), ReactivateOrganizationMembershipOpts{
		OrganizationMembership: "om_01E4ZCR3C56J083X43JQXF3JK5",
	})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, userRes)
}

func TestUsersGetInvitation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(getInvitationTestHandler))
	defer server.Close()

	DefaultClient = mockClient(server)
	SetAPIKey("test")

	expectedResponse := models.Invitation{
		ID:                  "invitation_123",
		Email:               "marcelina@foo-corp.com",
		State:               models.InvitationStatePending,
		Token:               "myToken",
		AcceptInvitationUrl: "https://your-app.com/invite?invitation_token=myToken",
		ExpiresAt:           time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
		CreatedAt:           time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
		UpdatedAt:           time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
	}

	getByIDRes, err := GetInvitation(context.Background(), GetInvitationOpts{
		Invitation: "invitation_123",
	})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, getByIDRes)
}

func TestUsersFindInvitationByToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(findInvitationByTokenTestHandler))
	defer server.Close()

	DefaultClient = mockClient(server)
	SetAPIKey("test")

	expectedResponse := models.Invitation{
		ID:                  "invitation_123",
		Email:               "marcelina@foo-corp.com",
		State:               models.InvitationStatePending,
		Token:               "myToken",
		AcceptInvitationUrl: "https://your-app.com/invite?invitation_token=myToken",
		ExpiresAt:           time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
		CreatedAt:           time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
		UpdatedAt:           time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
	}

	getByIDRes, err := FindInvitationByToken(context.Background(), FindInvitationByTokenOpts{
		InvitationToken: "myToken",
	})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, getByIDRes)
}

func TestUsersListInvitations(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(listInvitationsTestHandler))
	defer server.Close()

	DefaultClient = mockClient(server)
	SetAPIKey("test")

	expectedResponse :=
		ListInvitationsResponse{
			Data: []models.Invitation{
				{
					ID:                  "invitation_123",
					Email:               "marcelina@foo-corp.com",
					State:               models.InvitationStatePending,
					Token:               "myToken",
					AcceptInvitationUrl: "https://your-app.com/invite?invitation_token=myToken",
					ExpiresAt:           time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
					CreatedAt:           time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
					UpdatedAt:           time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
				},
			},
			ListMetadata: common.ListMetadata{
				After: "",
			},
		}

	listRes, err := ListInvitations(context.Background(), ListInvitationsOpts{
		Email: "marcelina@foo-corp.com",
	})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, listRes)
}

func TestUsersSendInvitation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(SendInvitationTestHandler))

	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	expectedResponse := models.Invitation{
		ID:                  "invitation_123",
		Email:               "marcelina@foo-corp.com",
		State:               models.InvitationStatePending,
		Token:               "myToken",
		AcceptInvitationUrl: "https://your-app.com/invite?invitation_token=myToken",
		ExpiresAt:           time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
		CreatedAt:           time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
		UpdatedAt:           time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
	}

	createRes, err := SendInvitation(context.Background(), SendInvitationOpts{
		Email:          "marcelina@foo-corp.com",
		OrganizationID: "org_123",
		ExpiresInDays:  7,
		InviterUserID:  "user_123",
		RoleSlug:       "admin",
	})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, createRes)
}

func TestUsersRevokeInvitation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(RevokeInvitationTestHandler))

	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	expectedResponse := models.Invitation{
		ID:                  "invitation_123",
		Email:               "marcelina@foo-corp.com",
		State:               models.InvitationStatePending,
		Token:               "myToken",
		AcceptInvitationUrl: "https://your-app.com/invite?invitation_token=myToken",
		ExpiresAt:           time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
		CreatedAt:           time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
		UpdatedAt:           time.Date(2021, 6, 25, 19, 07, 33, 155, time.UTC),
	}

	revokeRes, err := RevokeInvitation(context.Background(), RevokeInvitationOpts{
		Invitation: "invitation_123",
	})

	require.NoError(t, err)
	require.Equal(t, expectedResponse, revokeRes)
}

func TestUserManagementGetJWKSURL(t *testing.T) {
	client := NewClient("test")

	u, err := client.GetJWKSURL("client_123")
	require.NoError(t, err)
	require.Equal(
		t,
		"https://api.workos.com/sso/jwks/client_123",
		u.String(),
	)
}

func TestUsersManagementGetLogoutURL(t *testing.T) {
	client := NewClient("test")

	u, err := client.GetLogoutURL(GetLogoutURLOpts{SessionID: "session_abc"})
	require.NoError(t, err)
	require.Equal(
		t,
		"https://api.workos.com/user_management/sessions/logout?session_id=session_abc",
		u.String(),
	)
}

func TestUsersRevokeSession(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(RevokeSessionTestHandler))

	defer server.Close()

	DefaultClient = mockClient(server)

	SetAPIKey("test")

	err := RevokeSession(context.Background(), RevokeSessionOpts{
		SessionID: "session_123",
	})

	require.NoError(t, err)
}
