package controllers

import (
	"chapa/internal/domain"
	errors "chapa/pkg/config"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransferController struct {
	TransferUseCase domain.TransferUseCase
}

func NewTransferController(tuc domain.TransferUseCase) *TransferController {
	return &TransferController{
		TransferUseCase: tuc,
	}
}

func (tc *TransferController) Initiate(c *gin.Context) {
	var transfer domain.Transfer
	if err := c.ShouldBindJSON(&transfer); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid transfer data"})
		return
	}

	if err := tc.TransferUseCase.InitiateTransfer(c.Request.Context(), &transfer); err != nil {
		log.Println("Error initiating transfer:", err)
		code := errors.GetStatusCode(err)
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "transfer initiated", "transfer_ref": transfer.TransferRef})
}

func (tc *TransferController) Verify(c *gin.Context) {
	ref := c.Param("ref")
	log.Println("Verifying transfer with reference:", ref)
	transfer, err := tc.TransferUseCase.CheckTransferStatus(c.Request.Context(), ref)
	if err != nil {
		log.Println("Error verifying transfer:", err)
		code := errors.GetStatusCode(err)
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transfer)
}

func (tc *TransferController) Webhook(c *gin.Context) {
	signature := c.GetHeader("Chapa-Signature")
	if signature == "" {
		signature = c.GetHeader("x-chapa-signature")
	}
	if signature == "" {
		log.Println("Missing signature header")
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

	if err := tc.TransferUseCase.ProcessWebhook(c.Request.Context(), payload, signature); err != nil {
		log.Println("Webhook processing failed:", err)
		code := errors.GetStatusCode(err)
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "webhook processed successfully"})
}
