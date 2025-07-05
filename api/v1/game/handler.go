package game

import (
	"errors"
	"net/http"
	"quizzly-v2/internal/contracts"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	playerIDHeader = "X-Player-ID"
)

type Handler struct {
	usecase contracts.GameUsecase
}

func NewHandler(usecase contracts.GameUsecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) Register(router *gin.RouterGroup) {
	gameGroup := router.Group("/api/game")
	gameGroup.GET("/:game_id", h.getCurrentState)
}

func (h *Handler) getCurrentState(c *gin.Context) {
	gameIDStr := c.Param("game_id")
	gameID, err := uuid.Parse(gameIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid game_id format"})
		return
	}

	playerIDStr := c.GetHeader(playerIDHeader)
	if playerIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "header X-Player-ID is required"})
		return
	}
	playerID, err := uuid.Parse(playerIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid player_id format"})
		return
	}

	state, err := h.usecase.GetCurrentState(c.Request.Context(), gameID, playerID)
	if err != nil {
		if errors.Is(err, contracts.ErrGameNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := toStateResponse(state)
	c.JSON(http.StatusOK, response)
}
