package prediction

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type VertexAIRequest struct {
	Contents          []VertexAIContent         `json:"contents"`
	SystemInstruction VertexAISystemInstruction `json:"systemInstruction"`
}

type VertexAIContent struct {
	Parts []VertexAIPart `json:"parts"`
	Role  string         `json:"role"`
}

type VertexAIPart struct {
	Text string `json:"text"`
}

type VertexAISystemInstruction struct {
	Parts []VertexAIPart `json:"parts"`
}

type VertexAIResponse struct {
	Candidates    []VertexAICandidate   `json:"candidates"`
	UsageMetadata VertexAIUsageMetadata `json:"usageMetadata"`
	ModelVersion  string                `json:"modelVersion"`
}

type VertexAICandidate struct {
	Content      VertexAIContent `json:"content"`
	FinishReason string          `json:"finishReason"`
	AvgLogprobs  float64         `json:"avgLogprobs"`
}

type VertexAIUsageMetadata struct {
	PromptTokenCount        int                    `json:"promptTokenCount"`
	CandidatesTokenCount    int                    `json:"candidatesTokenCount"`
	TotalTokenCount         int                    `json:"totalTokenCount"`
	PromptTokensDetails     []VertexAITokenDetails `json:"promptTokensDetails"`
	CandidatesTokensDetails []VertexAITokenDetails `json:"candidatesTokensDetails"`
}

type VertexAITokenDetails struct {
	Modality   string `json:"modality"`
	TokenCount int    `json:"tokenCount"`
}

type VertexPredictor struct {
	httpClient *http.Client
	config     *VertexAIConfig
}

func NewVertexPredictor() (*VertexPredictor, error) {
	config := NewVertexAIConfig()
	if config.APIKey == "" || config.Endpoint == "" {
		return nil, fmt.Errorf("vertex AI configuration (API Key, Endpoint) is missing")
	}
	return &VertexPredictor{
		httpClient: &http.Client{Timeout: 30 * time.Second},
		config:     config,
	}, nil
}

func (vp *VertexPredictor) Predict(ctx context.Context, userPrompt string) (string, error) {
	// Construct the request URL
	requestURL := fmt.Sprintf("%s?key=%s",
		vp.config.Endpoint,
		url.QueryEscape(vp.config.APIKey))

	// Create the request body
	reqBody := VertexAIRequest{
		Contents: []VertexAIContent{
			{
				Parts: []VertexAIPart{{Text: userPrompt}},
				Role:  "user",
			},
		},
		SystemInstruction: VertexAISystemInstruction{
			Parts: []VertexAIPart{{Text: vp.config.SystemInstruction}},
		},
	}

	// Marshal the request body
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := vp.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Parse the response
	var predictionResp VertexAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&predictionResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	// Extract the prediction text
	if len(predictionResp.Candidates) == 0 || len(predictionResp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no prediction in response")
	}

	return predictionResp.Candidates[0].Content.Parts[0].Text, nil
}
