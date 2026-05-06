package handler

import (
	"net/http"

	"github.com/abg216te/tgbt/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, _ := h.service.List()
	c.JSON(http.StatusOK, users)
}
