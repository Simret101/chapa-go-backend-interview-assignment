package controllers

import (
	"chapa/internal/domain"
	errors "chapa/pkg/config"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	TransactionUseCase domain.TransactionUseCase
}

func NewTransactionController(tuc domain.TransactionUseCase) *TransactionController {
	return &TransactionController{
		TransactionUseCase: tuc,
	}
}

func (tc *TransactionController) Initiate(c *gin.Context) {
	var txn domain.Transaction
	if err := c.ShouldBindJSON(&txn); err != nil {
		code := errors.GetStatusCode(errors.ErrBadRequest)
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	if err := tc.TransactionUseCase.InitiatePayment(c.Request.Context(), &txn); err != nil {
		code := errors.GetStatusCode(err)
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "transaction initiated", "txn_ref": txn.TxnRef})
}

func (tc *TransactionController) Verify(c *gin.Context) {
	txnRef := c.Param("txnRef")
	txn, err := tc.TransactionUseCase.VerifyTransaction(c.Request.Context(), txnRef)
	if err != nil {
		code := errors.GetStatusCode(err)
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, txn)
}

func (tc *TransactionController) Webhook(c *gin.Context) {

	signature := c.GetHeader("Chapa-Signature")
	if signature == "" {
		signature = c.GetHeader("x-chapa-signature")
	}
	if signature == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing signature header"})
		return
	}

	var payload map[string]interface{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Println("Error binding JSON payload:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	payloadBytes, _ := json.Marshal(payload)
	log.Println("Payload Received:", string(payloadBytes))

	if err := tc.TransactionUseCase.ProcessWebhook(c.Request.Context(), payload, signature); err != nil {
		code := errors.GetStatusCode(err)
		log.Println("Webhook processing failed:", err)
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "webhook processed successfully"})
}
