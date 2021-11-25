package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"go_e-commerce-api/allproducts"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type allProductHandler struct {
	allproductsService allproducts.Service
}

func NewAllProductHandler(allproductsService allproducts.Service) *allProductHandler {
	return &allProductHandler{allproductsService}
}

// root handler murpakan bagian dari allproductsHandler struct
// digubakan untuk bisa mengakses lewat allproductshandler
func (h *allProductHandler) GetBooksList(c *gin.Context) {
	allproductss, err := h.allproductsService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var allproductssResponse []allproducts.AllProductResponse

	for _, b := range allproductss {
		allproductsResponse := converToAllProductResponse(b)
		allproductssResponse = append(allproductssResponse, allproductsResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": allproductssResponse,
	})
}

func (h *allProductHandler) GetBookById(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	// call service
	b, err := h.allproductsService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	allproductsResponse := converToAllProductResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": allproductsResponse,
	})
}

func (h *allProductHandler) GetBookByCategory(c *gin.Context) {
	category := c.Param("category")

	allproductss, err := h.allproductsService.FindByCategory(category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var allproductssResponse []allproducts.AllProductResponse

	for _, b := range allproductss {
		allproductsResponse := converToAllProductResponse(b)
		allproductssResponse = append(allproductssResponse, allproductsResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": allproductssResponse,
	})
}

func (h *allProductHandler) GetBookByProductName(c *gin.Context) {
	name_product := c.Param("name_product")
	email_user := c.Param("email_user")
	price := c.Param("price")

	allproductss, err := h.allproductsService.FindByNameProduct(name_product, price, email_user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var allproductssResponse []allproducts.AllProductResponse

	for _, b := range allproductss {
		allproductsResponse := converToAllProductResponse(b)
		allproductssResponse = append(allproductssResponse, allproductsResponse)
	}

	//BEWARE DONT TOUCH THIS CODE
	if allproductssResponse != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": allproductssResponse,
		})
	}
}

func (h *allProductHandler) GetBookByUser(c *gin.Context) {
	email_user := c.Param("email_user")

	allproductss, err := h.allproductsService.FindByUser(email_user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var allproductssResponse []allproducts.AllProductResponse

	for _, b := range allproductss {
		allproductsResponse := converToAllProductResponse(b)
		allproductssResponse = append(allproductssResponse, allproductsResponse)
	}

	//BEWARE DONT TOUCH THIS CODE
	if allproductssResponse != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": allproductssResponse,
		})
	}
}


func (h *allProductHandler) PostBooksHandler(c *gin.Context) {
	var allproductsRequest allproducts.AllProductRequest

	err := c.ShouldBindJSON(&allproductsRequest)

	if err != nil {
		// log.Fatal(err) -> kalau terjadi error, server mati
		for _, e := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("Error on filled %s, condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errMessage)

			// gunakan return untuk tidak melanjutkan yang dibawah
			return
		}
	}

	allproducts, err := h.allproductsService.Create(allproductsRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": allproducts,
	})
}

func (h *allProductHandler) UpdateBook(c *gin.Context) {
	var allproductsRequest allproducts.AllProductRequest

	err := c.ShouldBindJSON(&allproductsRequest)

	if err != nil {
		// log.Fatal(err) -> kalau terjadi error, server mati
		for _, e := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("Error on filled %s, condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errMessage)

			// gunakan return untuk tidak melanjutkan yang dibawah
			return
		}
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	allproducts, err := h.allproductsService.Update(id, allproductsRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": allproducts,
	})
}

func (h *allProductHandler) DeleteBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	// call service
	b, err := h.allproductsService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	allproductsResponse := converToAllProductResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data":    allproductsResponse,
		"Message": "Delete data success",
	})
}

func converToAllProductResponse(b allproducts.AllProduct) allproducts.AllProductResponse {
	return allproducts.AllProductResponse{
		Id:           b.Id,
		Name_product: b.Name_product,
		Image_url:    b.Image_url,
		Description:  b.Description,
		Price:        b.Price,
		Name_user:    b.Name_user,
		Email_user:   b.Email_user,
		Category:     b.Category,
	}
}
