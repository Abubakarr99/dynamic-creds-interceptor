package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Abubakarr99/dynamic-creds-interceptor/interceptor"
)

func TestAWSProviderSuccess(t *testing.T) {
	payload := map[string]interface{}{
		"body": map[string]interface{}{"ref": "refs/heads/main"},
	}
	jsonData, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(jsonData))
	req.Header.Set("X-Provider", "aws")
	rec := httptest.NewRecorder()

	handleInterceptorWithProvider(rec, req, &interceptor.MockProvider{})

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	var resp map[string]interface{}
	json.NewDecoder(rec.Body).Decode(&resp)
	if !resp["continue"].(bool) {
		t.Errorf("expected continue=true")
	}

	ext := resp["extensions"].(map[string]interface{})
	if ext["awsAccessKeyId"] != "MOCK_ACCESS_KEY_ID" {
		t.Errorf("unexpected awsAccessKeyId: %v", ext["awsAccessKeyId"])
	}
}

func TestUnknownProvider(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set("X-Provider", "unknown")
	rec := httptest.NewRecorder()

	handleInterceptor(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 for unknown provider, got %d", rec.Code)
	}
}
