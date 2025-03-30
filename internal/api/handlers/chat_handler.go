package handlers

import (
	"net/http"

	"github.com/arjkashyap/erlic.ai/internal/directory"
	"github.com/arjkashyap/erlic.ai/internal/directory/activedir"
)

type ChatHandler struct {
	dirManager directory.DirectoryManager
}

func NewChatHandler() *ChatHandler {

	return &ChatHandler{
		dirManager: activedir.NewADManager("", "", "", "", "", false, ""),
	}
}

func (ch *ChatHandler) handlePrompt(w http.ResponseWriter, r *http.Request) {

}
