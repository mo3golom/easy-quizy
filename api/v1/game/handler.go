package game

import (
	"easy-quizy/internal/contracts"
	"errors"
	"fmt"
	"net/http"

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
	gameGroup.POST("/:game_id/accept-answer", h.acceptAnswer)
	gameGroup.GET("/:game_id/reset", h.resetGame)
	gameGroup.GET("/daily", h.getDailyGame)
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

func (h *Handler) acceptAnswer(c *gin.Context) {
	gameIDStr := c.Param("game_id")
	gameID, err := uuid.Parse(gameIDStr)
	if err != nil {
		fmt.Printf("Invalid game ID: %s\n", gameIDStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid game_id format"})
		return
	}

	playerIDStr := c.GetHeader(playerIDHeader)
	if playerIDStr == "" {
		fmt.Printf("Missing X-Player-ID header\n")
		c.JSON(http.StatusBadRequest, gin.H{"error": "header X-Player-ID is required"})
		return
	}
	fmt.Printf("Player ID: %s\n", playerIDStr)
	playerID, err := uuid.Parse(playerIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid player_id format"})
		return
	}

	var req AcceptAnswerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body: " + err.Error()})
		return
	}

	out, err := h.usecase.AcceptAnswer(c.Request.Context(), &contracts.AcceptAnswersIn{
		GameID:     gameID,
		PlayerID:   playerID,
		QuestionID: req.QuestionID,
		Answer:     req.AnswerID,
	})
	if err != nil {
		if errors.Is(err, contracts.ErrGameNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, AcceptAnswerResponse{
		IsCorrect:   out.IsCorrect,
		Explanation: out.Explanation,
	})
}

func (h *Handler) resetGame(c *gin.Context) {
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

	err = h.usecase.Reset(c.Request.Context(), gameID, playerID)
	if err != nil {
		if errors.Is(err, contracts.ErrGameNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (h *Handler) getDailyGame(c *gin.Context) {
	game, err := h.usecase.GetDaily(c.Request.Context())
	if err != nil {
		if errors.Is(err, contracts.ErrGameNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, GetDailyGameResponse{
		GameID: game.ID.String(),
	})
}
