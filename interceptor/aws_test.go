package interceptor

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockProvider struct {
	shouldFail bool
}

func (m *mockProvider) GetCredentials() (*Credentials, error) {
	if m.shouldFail {
		return nil, errors.New("mock assume role failure")
	}
	return &Credentials{
		AccessKeyID:     "MOCKKEY",
		SecretAccessKey: "MOCKSECRET",
		SessionToken:    "MOCKSESSION",
		Provider:        "MOCK",
	}, nil
}

func TestAWSSTSProvider_Success(t *testing.T) {
	mock := &mockProvider{}
	creds, err := mock.GetCredentials()
	assert.NoError(t, err)
	assert.Equal(t, "MOCKKEY", creds.AccessKeyID)
	assert.Equal(t, "MOCKSECRET", creds.SecretAccessKey)
	assert.Equal(t, "MOCKSESSION", creds.SessionToken)
	assert.Equal(t, "MOCK", creds.Provider)
}

func TestAWSSTSProvider_Failure(t *testing.T) {
	mock := &mockProvider{shouldFail: true}
	creds, err := mock.GetCredentials()
	assert.Error(t, err)
	assert.Nil(t, creds)
}
