package laptop

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Laptop, error)
	FindByID(ID int) (Laptop, error)
	Create(laptop Laptop) (Laptop, error)
	Update(laptop Laptop) (Laptop, error)
	Delete(laptop Laptop) (Laptop, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Laptop, error) {
	var laptops []Laptop
	err := r.db.Find(&laptops).Error

	return laptops, err
}

func (r *repository) FindByID(ID int) (Laptop, error) {
	var laptops Laptop
	err := r.db.Find(&laptops, ID).Error

	return laptops, err
}

func (r *repository) Create(laptop Laptop) (Laptop, error) {
	err := r.db.Create(&laptop).Error

	return laptop, err
}

func (r *repository) Update(laptop Laptop) (Laptop, error) {
	err := r.db.Save(&laptop).Error

	return laptop, err
}

func (r *repository) Delete(laptop Laptop) (Laptop, error) {
	err := r.db.Delete(&laptop).Error

	return laptop, err
}
