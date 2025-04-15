package prediction

import "github.com/arjkashyap/erlic.ai/internal/env"

type VertexAIConfig struct {
	APIKey            string
	Endpoint          string
	Model             string
	SystemInstruction string
}

func NewVertexAIConfig() *VertexAIConfig {
	return &VertexAIConfig{
		APIKey:            env.GetString("VERTEX_AI_API_KEY", ""),
		Endpoint:          env.GetString("VERTEX_AI_ENDPOINT", ""),
		Model:             env.GetString("VERTEX_AI_MODEL", ""),
		SystemInstruction: SystemInstruction,
	}
}
