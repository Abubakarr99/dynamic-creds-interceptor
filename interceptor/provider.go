package interceptor

type Credentials struct {
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
	Provider        string
}

type CredentialProvider interface {
	GetCredentials() (*Credentials, error)
}
