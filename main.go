package main

import (
	"encoding/json"
	"github.com/Abubakarr99/dynamic-creds-interceptor/api"
	"github.com/Abubakarr99/dynamic-creds-interceptor/interceptor"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handleInterceptor)
	port := "8080"
	log.Printf("Starting Dynamic Credentials Interceptor on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleInterceptor(w http.ResponseWriter, r *http.Request) {
	var req api.InterceptorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	providerName := r.Header.Get("X-Provider")
	var provider interceptor.CredentialProvider

	switch providerName {
	case "aws":
		provider = &interceptor.AWSSTSProvider{
			RoleArn:     os.Getenv("AWS_ROLE_ARN"),
			SessionName: "tekton-session",
			Region:      os.Getenv("AWS_REGION"),
		}
	default:
		http.Error(w, "unsupported provider", http.StatusBadRequest)
		return
	}

	handleInterceptorWithProvider(w, r, provider)
}

func handleInterceptorWithProvider(w http.ResponseWriter, r *http.Request, provider interceptor.CredentialProvider) {
	var req api.InterceptorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	creds, err := provider.GetCredentials()
	if err != nil {
		resp := api.InterceptorResponse{Continue: false}
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp := api.InterceptorResponse{
		Continue: true,
		Extensions: map[string]interface{}{
			"awsAccessKeyId":     creds.AccessKeyID,
			"awsSecretAccessKey": creds.SecretAccessKey,
			"awsSessionToken":    creds.SessionToken,
			"credentialProvider": creds.Provider,
		},
	}
	json.NewEncoder(w).Encode(resp)
}
