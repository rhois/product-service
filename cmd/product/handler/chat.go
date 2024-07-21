package handler

import (
	"encoding/json"
	"net/http"
	"product/internal/presenter"
	"product/internal/usecase/chat"

	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	ChatService chat.Service
}

func (h *ChatHandler) AskAIHandler(c *gin.Context) {
	var queryReq presenter.ChatRequest
	err := json.NewDecoder(c.Request.Body).Decode(&queryReq)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	responseText, err := h.ChatService.GenerateResponse(c, queryReq.Query)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	response := presenter.ChatResponse{
		Response: responseText,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(c.Writer).Encode(response)
}
