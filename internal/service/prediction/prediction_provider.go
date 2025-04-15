// internal/service/prediction/prediction_provider.go
package prediction

import (
	"fmt"

	"github.com/arjkashyap/erlic.ai/internal/logger"
)

// NewPredictionService is a factory function that creates and returns the

func NewPredictionService(mlProvider string) (MLPredictor, error) {
	logger.Logger.Info("Initializing ML Prediction Service provider: " + mlProvider)

	switch mlProvider {
	case "vertex":
		// Call the specific constructor which returns the concrete type and error
		predictor, err := NewVertexPredictor()
		if err != nil {
			logger.Logger.Error("Failed to initialize VertexPredictor", "error", err)
			return nil, fmt.Errorf("failed to create vertex predictor: %w", err)
		}
		logger.Logger.Info("Successfully initialized Vertex AI provider.")
		return predictor, nil

	case "azure":
		// Call the specific constructor

		return nil, fmt.Errorf("azure provider not implemented")

	default:
		return nil, fmt.Errorf("unsupported or invalid ML_PROVIDER configured: %s", mlProvider)
	}
}
