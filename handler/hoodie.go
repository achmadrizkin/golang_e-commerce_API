package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"go_e-commerce-api/hoodie"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type hoodieHandler struct {
	hoodieService hoodie.Service
}

func NewHoodieHandler(hoodieService hoodie.Service) *hoodieHandler {
	return &hoodieHandler{hoodieService}
}

// root handler murpakan bagian dari hoodieHandler struct
// digubakan untuk bisa mengakses lewat hoodiehandler
func (h *hoodieHandler) GetBooksList(c *gin.Context) {
	hoodies, err := h.hoodieService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var hoodiesResponse []hoodie.HoodieResponse

	for _, b := range hoodies {
		hoodieResponse := converToHoodieResponse(b)
		hoodiesResponse = append(hoodiesResponse, hoodieResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": hoodiesResponse,
	})
}

func (h *hoodieHandler) GetBookById(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	// call service
	b, err := h.hoodieService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	hoodieResponse := converToHoodieResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": hoodieResponse,
	})
}

func (h *hoodieHandler) PostBooksHandler(c *gin.Context) {
	var hoodieRequest hoodie.HoodieRequest

	err := c.ShouldBindJSON(&hoodieRequest)

	if err != nil {
		// log.Fatal(err) -> kalau terjadi error, server mati
		for _, e := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("Error on filled %s, condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errMessage)

			// gunakan return untuk tidak melanjutkan yang dibawah
			return
		}
	}

	hoodie, err := h.hoodieService.Create(hoodieRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": hoodie,
	})
}

func (h *hoodieHandler) UpdateBook(c *gin.Context) {
	var hoodieRequest hoodie.HoodieRequest

	err := c.ShouldBindJSON(&hoodieRequest)

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
	hoodie, err := h.hoodieService.Update(id, hoodieRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": hoodie,
	})
}

func (h *hoodieHandler) DeleteBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	// call service
	b, err := h.hoodieService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	hoodieResponse := converToHoodieResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data":    hoodieResponse,
		"Message": "Delete data success",
	})
}

func converToHoodieResponse(b hoodie.Hoodie) hoodie.HoodieResponse {
	return hoodie.HoodieResponse{
		Id:           b.Id,
		Name_product: b.Name_product,
		Image_url:    b.Image_url,
		Description:  b.Description,
		Price:        b.Price,
		Name_user:    b.Name_user,
		Email_user:   b.Email_user,
	}
}
