package handlers

import (
	"net/http"

	"github.com/arjkashyap/erlic.ai/internal/logger"
	"github.com/arjkashyap/erlic.ai/internal/utils"
	"github.com/markbates/goth/gothic"
)

type DashboardHandler struct {
	// Add any dependencies here if needed in the future
}

func NewDashboardHandler() *DashboardHandler {
	return &DashboardHandler{}
}

// GetDashboard returns basic dashboard data for the authenticated user
func (dh *DashboardHandler) GetDashboard(w http.ResponseWriter, r *http.Request) {
	session, _ := gothic.Store.Get(r, gothic.SessionName)
	userID := session.Values["user_id"].(int)

	// Placeholder dashboard data
	dashboardData := utils.Envelope{
		"user_id": userID,
		"message": "Welcome to your dashboard",
		"status":  "active",
	}

	utils.WriteJSON(w, http.StatusOK, dashboardData, nil, logger.Logger)
}
