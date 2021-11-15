package main

import (
	"go_e-commerce-api/allproducts"
	"go_e-commerce-api/book"
	"go_e-commerce-api/handler"
	"go_e-commerce-api/hoodie"
	"go_e-commerce-api/laptop"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	// connect to database
	dsn := "root:@tcp(127.0.0.1:3306)/go-ecommerce?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection Error")
	}

	// auto migrate (auto add table)
	db.AutoMigrate(&allproducts.AllProducts{}, &book.Book{}, &hoodie.Hoodie{}, &laptop.Laptop{})

	// API Versioning
	v1 := r.Group("/v1")

	// 		BOOK
	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	v1.POST("/products/book", bookHandler.PostBooksHandler)
	v1.GET("/products/book", bookHandler.GetBooksList)
	v1.GET("/products/book/:id", bookHandler.GetBookById)
	v1.PUT("/products/book/:id", bookHandler.UpdateBook)
	v1.DELETE("/products/book/:id", bookHandler.DeleteBook)
	
	//		Laptop


	//		Hoodie
	hoodieRepository := hoodie.NewRepository(db)
	hoodieService := hoodie.NewService(hoodieRepository)
	hoodieHandler := handler.NewHoodieHandler(hoodieService)

	v1.POST("/products/hoodie", hoodieHandler.PostBooksHandler)
	v1.GET("/products/hoodie", hoodieHandler.GetBooksList)
	v1.GET("/products/hoodie/:id", hoodieHandler.GetBookById)
	v1.PUT("/products/hoodie/:id", hoodieHandler.UpdateBook)
	v1.DELETE("/products/hoodie/:id", hoodieHandler.DeleteBook)

	// 
	r.Run(":3000")
}