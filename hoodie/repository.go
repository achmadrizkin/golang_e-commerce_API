package hoodie

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Hoodie, error)
	FindByID(ID int) (Hoodie, error)
	Create(hoodie Hoodie) (Hoodie, error)
	Update(hoodie Hoodie) (Hoodie, error)
	Delete(hoodie Hoodie) (Hoodie, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Hoodie, error) {
	var hoodies []Hoodie
	err := r.db.Find(&hoodies).Error

	return hoodies, err
}

func (r *repository) FindByID(ID int) (Hoodie, error) {
	var hoodies Hoodie
	err := r.db.Find(&hoodies, ID).Error

	return hoodies, err
}

func (r *repository) Create(hoodie Hoodie) (Hoodie, error) {
	err := r.db.Create(&hoodie).Error

	return hoodie, err
}

func (r *repository) Update(hoodie Hoodie) (Hoodie, error) {
	err := r.db.Save(&hoodie).Error

	return hoodie, err
}

func (r *repository) Delete(hoodie Hoodie) (Hoodie, error) {
	err := r.db.Delete(&hoodie).Error

	return hoodie, err
}
