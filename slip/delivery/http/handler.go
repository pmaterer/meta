package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pmaterer/meta/slip"
)

type service interface {
	CreateSlip(slip slip.Slip) error
}

type Handler struct {
	service service
}

func NewHandler(s service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) CreateSlip(g *gin.Context) {
	var slip slip.Slip
	if err := g.ShouldBindJSON(&slip); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateSlip(slip); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"message": "OK"})
}
