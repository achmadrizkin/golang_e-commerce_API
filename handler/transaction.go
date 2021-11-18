package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"go_e-commerce-api/transaction"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type transactionHandler struct {
	transactionService transaction.Service
}

func NewTransactionHandler(transactionService transaction.Service) *transactionHandler {
	return &transactionHandler{transactionService}
}

// root handler murpakan bagian dari transactionHandler struct
// digubakan untuk bisa mengakses lewat transactionhandler
func (h *transactionHandler) GetBooksList(c *gin.Context) {
	transactions, err := h.transactionService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var transactionsResponse []transaction.TransactionResponse

	for _, b := range transactions {
		transactionResponse := converToTransactionResponse(b)
		transactionsResponse = append(transactionsResponse, transactionResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": transactionsResponse,
	})
}

func (h *transactionHandler) GetBookById(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	// call service
	b, err := h.transactionService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	transactionResponse := converToTransactionResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": transactionResponse,
	})
}

func (h *transactionHandler) PostBooksHandler(c *gin.Context) {
	var transactionRequest transaction.TransactionRequest

	err := c.ShouldBindJSON(&transactionRequest)

	if err != nil {
		// log.Fatal(err) -> kalau terjadi error, server mati
		for _, e := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("Error on filled %s, condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errMessage)

			// gunakan return untuk tidak melanjutkan yang dibawah
			return
		}
	}

	transaction, err := h.transactionService.Create(transactionRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": transaction,
	})
}

func (h *transactionHandler) UpdateBook(c *gin.Context) {
	var transactionRequest transaction.TransactionRequest

	err := c.ShouldBindJSON(&transactionRequest)

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
	transaction, err := h.transactionService.Update(id, transactionRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": transaction,
	})
}

func (h *transactionHandler) DeleteBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	// call service
	b, err := h.transactionService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	transactionResponse := converToTransactionResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data":    transactionResponse,
		"Message": "Delete data success",
	})
}

func converToTransactionResponse(b transaction.Transaction) transaction.TransactionResponse {
	return transaction.TransactionResponse{
		Id:           b.Id,
		Name_product: b.Name_product,
		Image_url:    b.Image_url,
		Description:  b.Description,
		Price:        b.Price,
		Name_user:    b.Name_user,
		Email_user:   b.Email_user,
		Name_buyer:   b.Name_buyer,
		Email_buyer:  b.Email_buyer,
	}
}
