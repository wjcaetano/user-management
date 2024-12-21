package handler

import (
	"net/http"
	"user-management/app/user/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userService service.UserService
}

func NewHandler(s service.UserService) *Handler {
	return &Handler{userService: s}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	inputUser := req.ToEntity()

	user, err := h.userService.CreateUser(inputUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := ToResponse(user)

	c.JSON(http.StatusOK, response)
}
