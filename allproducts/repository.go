package allproducts

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]AllProduct, error)
	FindByID(ID int) (AllProduct, error)
	FindByCategory(category string) ([]AllProduct, error)
	FindByUser(email_user string) ([]AllProduct, error)
	FindByNameProduct(name_product string, price string, email_user string) ([]AllProduct, error)
	Create(allProduct AllProduct) (AllProduct, error)
	Update(allProduct AllProduct) (AllProduct, error)
	Delete(allProduct AllProduct) (AllProduct, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]AllProduct, error) {
	var allProducts []AllProduct
	err := r.db.Find(&allProducts).Error

	return allProducts, err
}

func (r *repository) FindByID(ID int) (AllProduct, error) {
	var allProducts AllProduct
	err := r.db.Find(&allProducts, ID).Error

	return allProducts, err
}

func (r *repository) FindByCategory(category string) ([]AllProduct, error) {
	var allProducts []AllProduct
	// err := r.db.Where("title = ?", title).First(&books).Error

	err := r.db.Where("category LIKE ?", category).Find(&allProducts).Error

	return allProducts, err
}

func (r *repository) FindByUser(email_user string) ([]AllProduct, error) {
	var allProducts []AllProduct
	// err := r.db.Where("title = ?", title).First(&books).Error

	err := r.db.Where("email_user LIKE ?", email_user).Find(&allProducts).Error

	return allProducts, err
}

func (r *repository) FindByNameProduct(name_product string, price string, email_user string) ([]AllProduct, error) {
	var allProducts []AllProduct
	// err := r.db.Where("title = ?", title).First(&books).Error

    value := fmt.Sprintf("%%%s%%", name_product)
	err := r.db.Where("name_product LIKE ? AND email_user = ? AND price = ?", value, email_user, price).Find(&allProducts).Error

	// err := r.db.Where("email_user LIKE ?", name_product).Find(&allProducts).Error

	return allProducts, err
}

func (r *repository) Create(allProduct AllProduct) (AllProduct, error) {
	err := r.db.Create(&allProduct).Error

	return allProduct, err
}

func (r *repository) Update(allProduct AllProduct) (AllProduct, error) {
	err := r.db.Save(&allProduct).Error

	return allProduct, err
}

func (r *repository) Delete(allProduct AllProduct) (AllProduct, error) {
	err := r.db.Delete(&allProduct).Error

	return allProduct, err
}
