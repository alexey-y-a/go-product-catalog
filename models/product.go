package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string         `json:"name" validate:"required,min=2,max=100"`
	Description string         `json:"description" validate:"required"`
	Images      pq.StringArray `json:"images" gorm:"type:text[]" validate:"dive,uri"`
}

func (p *Product) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}
