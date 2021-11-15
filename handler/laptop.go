package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"go_e-commerce-api/laptop"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type laptopHandler struct {
	laptopService laptop.Service
}

func NewLaptopHandler(laptopService laptop.Service) *laptopHandler {
	return &laptopHandler{laptopService}
}

// root handler murpakan bagian dari laptopHandler struct
// digubakan untuk bisa mengakses lewat laptophandler
func (h *laptopHandler) GetBooksList(c *gin.Context) {
	laptops, err := h.laptopService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var laptopsResponse []laptop.LaptopResponse

	for _, b := range laptops {
		laptopResponse := converToLaptopResponse(b)
		laptopsResponse = append(laptopsResponse, laptopResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": laptopsResponse,
	})
}

func (h *laptopHandler) GetBookById(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	// call service
	b, err := h.laptopService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	laptopResponse := converToLaptopResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": laptopResponse,
	})
}

func (h *laptopHandler) PostBooksHandler(c *gin.Context) {
	var laptopRequest laptop.LaptopRequest

	err := c.ShouldBindJSON(&laptopRequest)

	if err != nil {
		// log.Fatal(err) -> kalau terjadi error, server mati
		for _, e := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("Error on filled %s, condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errMessage)

			// gunakan return untuk tidak melanjutkan yang dibawah
			return
		}
	}

	laptop, err := h.laptopService.Create(laptopRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": laptop,
	})
}

func (h *laptopHandler) UpdateBook(c *gin.Context) {
	var laptopRequest laptop.LaptopRequest

	err := c.ShouldBindJSON(&laptopRequest)

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
	laptop, err := h.laptopService.Update(id, laptopRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": laptop,
	})
}

func (h *laptopHandler) DeleteBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	// call service
	b, err := h.laptopService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	laptopResponse := converToLaptopResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data":    laptopResponse,
		"Message": "Delete data success",
	})
}

func converToLaptopResponse(b laptop.Laptop) laptop.LaptopResponse {
	return laptop.LaptopResponse{
		Id:           b.Id,
		Name_product: b.Name_product,
		Image_url:    b.Image_url,
		Description:  b.Description,
		Price:        b.Price,
		Name_user:    b.Name_user,
		Email_user:   b.Email_user,
	}
}
