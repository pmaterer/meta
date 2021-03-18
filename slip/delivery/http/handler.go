package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pmaterer/meta/slip"
)

type service interface {
	CreateSlip(slip slip.Slip) error
	GetSlip(id int64) (slip.Slip, error)
	GetAllSlips() ([]slip.Slip, error)
	UpdateSlip(slip slip.Slip) error
	DeleteSlip(id int64) error
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

func (h *Handler) GetSlip(g *gin.Context) {
	// TODO: handle slip not found
	paramID := g.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	slip, err := h.service.GetSlip(int64(id))
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, slip)
}

func (h *Handler) GetAllSlips(g *gin.Context) {
	slips, err := h.service.GetAllSlips()
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	g.JSON(http.StatusOK, slips)
}

func (h *Handler) UpdateSlip(g *gin.Context) {
	paramID := g.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var slip slip.Slip
	if err := g.ShouldBindJSON(&slip); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	slip.ID = int64(id)
	if err := h.service.UpdateSlip(slip); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *Handler) DeleteSlip(g *gin.Context) {
	paramID := g.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.service.DeleteSlip(int64(id))
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	g.JSON(http.StatusOK, gin.H{"message": "OK"})
}
