package handler

import (
	"net/http"

	"project-devboard/internal/dtos"
	"project-devboard/internal/services"
	"project-devboard/pkg/apperrors"
	"project-devboard/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeveloperDashboardHandler struct {
	service services.DeveloperDashboardService
}

func NewDeveloperDashboardHandler(service services.DeveloperDashboardService) *DeveloperDashboardHandler {
	return &DeveloperDashboardHandler{service: service}
}

// GetDashboardData godoc
// @Summary      Get developer dashboard data
// @Description  Get all profile details for the authenticated developer's dashboard
// @Tags         developer-dashboard
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.Response{data=dtos.DeveloperDashboardResponse}
// @Failure      401  {object}  response.Response  "Unauthorized"
// @Failure      403  {object}  response.Response  "Forbidden"
// @Failure      404  {object}  response.Response  "User not found"
// @Failure      500  {object}  response.Response  "Internal Server Error"
// @Router       /developer/dashboard [get]
func (h *DeveloperDashboardHandler) GetDashboardData(c *gin.Context) {
	userID := userIDFromContext(c, uuid.Nil)
	if userID == uuid.Nil {
		c.Error(apperrors.New(apperrors.Unauthorized, apperrors.ErrUnauthorized))
		return
	}

	user, err := h.service.GetDashboardData(userID)
	if err != nil {
		c.Error(err)
		return
	}

	res := dtos.NewDeveloperDashboardResponse(user)
	response.Success(c, http.StatusOK, res)
}
