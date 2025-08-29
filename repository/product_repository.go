package repository

import (
	"errors"
	"go_product_catalog/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *models.Product) error {
	if err := product.Validate(); err != nil {
		return err
	}
	return r.db.Create(product).Error
}

func (r *ProductRepository) GetByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := r.db.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) Update(id uint, updatedProduct *models.Product) error {
	if err := updatedProduct.Validate(); err != nil {
		return err
	}

	var existing models.Product
	if err := r.db.First(&existing, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}
		return err
	}

	existing.Name = updatedProduct.Name
	existing.Description = updatedProduct.Description
	existing.Images = updatedProduct.Images

	return r.db.Save(&existing).Error
}

func (r *ProductRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("product not found")
	}
	return nil
}
