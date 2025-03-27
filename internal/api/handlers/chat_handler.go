package handlers

import (
	"github.com/arjkashyap/erlic.ai/internal/directory"
	"github.com/arjkashyap/erlic.ai/internal/directory/activedir"
)

type ChatHandler struct {
	dirManager directory.DirectoryManager
}

func NewChatHandler() *ChatHandler {

	return &ChatHandler{
		dirManager: activedir.NewADManager("", "", "", "", false),
	}
}

func handlePrompt() {

}
