package handlers

import (
	"fastfood-api/internal/models"
	"fastfood-api/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MenuHandler struct {
	repo *repository.MenuRepository
}

func NewMenuHandler(repo *repository.MenuRepository) *MenuHandler {
	return &MenuHandler{repo: repo}
}

func (h *MenuHandler) GetAllItems(c *gin.Context) {
	items, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *MenuHandler) GetItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	item, err := h.repo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *MenuHandler) CreateItem(c *gin.Context) {
	var item models.MenuItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.Create(item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, item)
}

func (h *MenuHandler) UpdateItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var item models.MenuItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item.ID = uint(id)
	if err := h.repo.Update(item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *MenuHandler) DeleteItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}