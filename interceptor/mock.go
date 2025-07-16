package interceptor

import "errors"

// MockProvider is a simple mock implementation of CredentialProvider
type MockProvider struct {
	Fail bool
}

func (m *MockProvider) GetCredentials() (*Credentials, error) {
	if m.Fail {
		return nil, errors.New("mock assume role failure")
	}
	return &Credentials{
		AccessKeyID:     "MOCK_ACCESS_KEY_ID",
		SecretAccessKey: "MOCK_SECRET_ACCESS_KEY",
		SessionToken:    "MOCK_SESSION_TOKEN",
		Provider:        "MOCK",
	}, nil
}
