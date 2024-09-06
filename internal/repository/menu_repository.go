package repository

import (
	"fastfood-api/internal/models"
	"gorm.io/gorm"
)

type MenuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) *MenuRepository {
	return &MenuRepository{db: db}
}

func (r *MenuRepository) GetAll() ([]models.MenuItem, error) {
	var items []models.MenuItem
	result := r.db.Find(&items)
	return items, result.Error
}

func (r *MenuRepository) GetByID(id uint) (models.MenuItem, error) {
	var item models.MenuItem
	result := r.db.First(&item, id)
	return item, result.Error
}

func (r *MenuRepository) Create(item models.MenuItem) error {
	return r.db.Create(&item).Error
}

func (r *MenuRepository) Update(item models.MenuItem) error {
	return r.db.Save(&item).Error
}

func (r *MenuRepository) Delete(id uint) error {
	return r.db.Delete(&models.MenuItem{}, id).Error
}