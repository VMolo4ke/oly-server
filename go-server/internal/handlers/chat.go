package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"oly-server/internal/middleware"
	"oly-server/internal/services"
)

type ChatHandler struct {
	chatService *services.ChatService
}

func NewChatHandler(chatService *services.ChatService) *ChatHandler {
	return &ChatHandler{chatService: chatService}
}

func getUserIDFromContext(ctx context.Context) string {
	if userID := ctx.Value(middleware.UserIDKey); userID != nil {
		if str, ok := userID.(string); ok {
			return str
		}
	}
	return ""
}

func getUserIDFromHeader(r *http.Request) string {
	userID := r.Header.Get("X-User-Id")
	return userID
}

func (h *ChatHandler) GetUserChats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := getUserIDFromHeader(r)
	if userID == "" {
		http.Error(w, "X-User-Id header required", http.StatusBadRequest)
		return
	}

	response, err := h.chatService.GetUserChats(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *ChatHandler) CreateChat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input services.CreateChatInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.chatService.Create(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *ChatHandler) AddParticipant(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	chatID := strings.TrimPrefix(r.URL.Path, "/chat/participants/")
	if chatID == "" || chatID == r.URL.Path {
		http.Error(w, "Chat ID required", http.StatusBadRequest)
		return
	}

	var input struct {
		UserID string `json:"userId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.chatService.AddParticipant(chatID, input.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
