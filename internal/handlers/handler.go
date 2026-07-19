package handlers

import (
	"log/slog"
	"net/http"
	"url-shortener/internal/models"
	"url-shortener/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	serv *services.Service
	log  *slog.Logger
}

func New(serv *services.Service, logger *slog.Logger) (h *Handler) {
	return &Handler{
		serv: serv,
		log:  logger,
	}
}

func (h *Handler) Set(c *gin.Context) {
	var link models.Link

	if err := c.ShouldBindJSON(&link); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.log.Debug("json formated", slog.Any("link", link))

	urlShort, err := h.serv.Set(&link)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, urlShort)
}

func (h *Handler) Get(c *gin.Context) {
	urlShort := c.Param("url")
	link, err := h.serv.Get(urlShort)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, link)
}
