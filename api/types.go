package api

type InterceptorRequest struct {
	Context map[string]interface{} `json:"context"`
	Body    map[string]interface{} `json:"body"`
	Header  map[string][]string    `json:"header"`
}

type InterceptorResponse struct {
	Continue   bool                   `json:"continue"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}
