package main

import (
	"go_e-commerce-api/allproducts"
	"go_e-commerce-api/book"
	"go_e-commerce-api/hoodie"
	"go_e-commerce-api/laptop"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()

	// connect to database
	dsn := "root:@tcp(127.0.0.1:3306)/go-ecommerce?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection Error")
	}

	db.AutoMigrate(&allproducts.AllProducts{}, &book.Book{}, &hoodie.Hoodie{}, &laptop.Laptop{})


	//
	router.Run(":3000")
}