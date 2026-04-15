package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func userIDFromContext(c *gin.Context, fallback uuid.UUID) uuid.UUID {
	userID, exists := c.Get("user_id")
	if !exists {
		return fallback
	}
	parsedUserID, ok := userID.(uuid.UUID)
	if !ok {
		return fallback
	}
	return parsedUserID
}

func paginationParams(c *gin.Context) (int, int) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil || offset < 0 {
		offset = 0
	}

	return limit, offset
}
