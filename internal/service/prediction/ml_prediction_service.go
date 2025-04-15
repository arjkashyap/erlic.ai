package prediction

import (
	"context"
)

// MLPredictor defines the standard contract for interacting with
// various underlying Machine Learning models (LLMs) for predictions.
// This allows for swapping providers (Vertex AI, Azure OpenAI, etc.) easily.
type MLPredictor interface {
	// Predict sends a prompt (with optional system instructions)
	// to the configured ML provider and returns the generated text content.
	Predict(ctx context.Context, userPrompt string) (generatedText string, err error)
}
