package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iotassss/domainmodel/internal/domain"
)

func handleDomainError(c *gin.Context, err error) {
	switch err.(type) {
	case *domain.ValidationError:
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case *domain.NotFoundError:
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case *domain.ConflictError:
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	case *domain.ServerError:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
