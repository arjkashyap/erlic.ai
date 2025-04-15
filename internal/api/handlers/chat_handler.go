package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/arjkashyap/erlic.ai/internal/config"
	"github.com/arjkashyap/erlic.ai/internal/logger"
	"github.com/arjkashyap/erlic.ai/internal/service/directory"
	"github.com/arjkashyap/erlic.ai/internal/service/directory/activedir"
	"github.com/arjkashyap/erlic.ai/internal/service/prediction"
	"go.uber.org/zap"
)

type ChatHandler struct {
	dirManager directory.DirectoryManager
	predictor  prediction.MLPredictor
}

type ChatRequest struct {
	Message string `json:"message"`
}

type ChatResponse struct {
	Response      string                    `json:"response"`
	FunctionCalls []prediction.FunctionCall `json:"function_calls"`
}

func NewChatHandler(cfg *config.Config) *ChatHandler {
	predictor, err := prediction.NewPredictionService(cfg.MLProvider)
	if err != nil {
		logger.Logger.Error("Failed to initialize ML Prediction Service provider: "+cfg.MLProvider, zap.Error(err))
		panic(err)
	}

	return &ChatHandler{
		dirManager: activedir.NewADManager("", "", "", "", "", false, ""),
		predictor:  predictor,
	}
}

func (ch *ChatHandler) HandlePrompt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ChatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	logger.Logger.Infof("Received chat message: %s", req.Message)
	logger.Logger.Info("Predicting chat message")

	response, err := ch.predictor.Predict(context.Background(), req.Message)
	if err != nil {
		logger.Logger.Error("Failed to predict chat message", zap.Error(err))
		http.Error(w, "ML Prediction service: Internal server error", http.StatusInternalServerError)
		return
	}

	// Deserialize the response into our types
	var functionCalls []prediction.FunctionCall
	if err := json.Unmarshal([]byte(response), &functionCalls); err != nil {
		logger.Logger.Error("Failed to deserialize prediction response", zap.Error(err))
		http.Error(w, "Failed to process prediction response", http.StatusInternalServerError)
		return
	}

	// Create response with the deserialized data
	chatResponse := ChatResponse{
		Response:      response,
		FunctionCalls: functionCalls,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(chatResponse); err != nil {
		logger.Logger.Error("Failed to encode response", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
