package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iotassss/domainmodel/internal/usecase"
)

type UserHandler struct {
	createUser    usecase.CreateUserUseCase
	getUserByUUID usecase.GetUserByUUIDUseCase
}

func NewUserHandler(
	createUser usecase.CreateUserUseCase,
	getUserByUUID usecase.GetUserByUUIDUseCase,
) *UserHandler {
	return &UserHandler{
		createUser:    createUser,
		getUserByUUID: getUserByUUID,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.createUser.Execute(c.Request.Context(), req.FirstName, req.LastName, req.Email, req.Password)
	if err != nil {
		handleDomainError(c, err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (uc *UserHandler) GetUserByUUID(c *gin.Context) {
	uuid := c.Param("uuid")
	user, err := uc.getUserByUUID.Execute(c.Request.Context(), uuid)
	if err != nil {
		handleDomainError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
