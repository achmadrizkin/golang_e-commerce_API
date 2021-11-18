package main

import (
	"go_e-commerce-api/allproducts"
	"go_e-commerce-api/book"
	"go_e-commerce-api/handler"
	"go_e-commerce-api/hoodie"
	"go_e-commerce-api/laptop"
	"go_e-commerce-api/transaction"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	// connect to database
	// PLEASE CREATE go-ecommerce database first.
	dsn := "root:@tcp(127.0.0.1:3306)/go-ecommerce?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection Error")
	}

	// auto migrate (auto add table)
	db.AutoMigrate(&allproducts.AllProduct{}, &book.Book{}, &hoodie.Hoodie{}, &laptop.Laptop{}, &transaction.Transaction{})

	// API Versioning
	v1 := r.Group("/v1")


	// All Products
	allProductRepository := allproducts.NewRepository(db)
	allProductService := allproducts.NewService(allProductRepository)
	allProductHandler := handler.NewAllProductHandler(allProductService)

	v1.POST("/products", allProductHandler.PostBooksHandler)
	v1.GET("/products", allProductHandler.GetBooksList)
	v1.GET("/products/:id", allProductHandler.GetBookById)
	v1.GET("/products/category/:category", allProductHandler.GetBookByCategory)
	v1.GET("/products/user/:email_user", allProductHandler.GetBookByUser)
	v1.PUT("/products/:id", allProductHandler.UpdateBook)
	v1.DELETE("/products/:id", allProductHandler.DeleteBook)

	// Transaction
	transactionRepository := transaction.NewRepository(db)
	transactionService := transaction.NewService(transactionRepository)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	v1.POST("/transaction", transactionHandler.PostBooksHandler)
	v1.GET("/transaction", transactionHandler.GetBooksList)
	v1.GET("/transaction/:id", transactionHandler.GetBookById)
	v1.PUT("/transaction/:id", transactionHandler.UpdateBook)
	v1.DELETE("/transaction/:id", transactionHandler.DeleteBook)


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
	laptopRepository := laptop.NewRepository(db)
	laptopService := laptop.NewService(laptopRepository)
	laptopHandler := handler.NewLaptopHandler(laptopService)

	v1.POST("/products/laptop", laptopHandler.PostBooksHandler)
	v1.GET("/products/laptop", laptopHandler.GetBooksList)
	v1.GET("/products/laptop/:id", laptopHandler.GetBookById)
	v1.PUT("/products/laptop/:id", laptopHandler.UpdateBook)
	v1.DELETE("/products/laptop/:id", laptopHandler.DeleteBook)

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