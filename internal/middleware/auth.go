package middleware

import (
	"easy-quizy/internal/contracts"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	UserIDKey = "userID"
)

// AuthMiddleware validates X-Player-ID and X-Source headers and retrieves user
func AuthMiddleware(userUsecase contracts.UserUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		playerID := c.GetHeader("X-Player-ID")
		source := c.GetHeader("X-Source")

		// Validate headers are not empty
		if playerID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "X-Player-ID header is required"})
			c.Abort()
			return
		}

		if source == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "X-Source header is required"})
			c.Abort()
			return
		}

		// Retrieve user
		user, err := userUsecase.RetrieveUser(c.Request.Context(), playerID, source)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Store internal UserID in gin context
		c.Set(UserIDKey, user.ID)
		c.Next()
	}
}

// GetUserID retrieves the user ID from gin context
func GetUserID(c *gin.Context) (uuid.UUID, bool) {
	userID, exists := c.Get(UserIDKey)
	if !exists {
		return uuid.Nil, false
	}

	id, ok := userID.(uuid.UUID)
	return id, ok
}
