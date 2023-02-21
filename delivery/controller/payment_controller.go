package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-payment/model"
	"simple-payment/usecase"
	"strconv"
	"strings"
)

type PaymentController struct {
	rg             *gin.RouterGroup
	paymentUseCase usecase.PaymentUseCase
}

// @Summary Get all payments
// @Tags payment
// @success 200 {object} model.PaymentResponse{data=[]model.Payment}
// @Router /api/payments [get]
func (cc *PaymentController) getPayments(ctx *gin.Context) {

	payments, err := cc.paymentUseCase.Payments()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get payments",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Get all payments",
		"payments": payments,
	})
}

// @Summary Create new payment
// @Tags payment
// @Param string body model.PaymentRequest true "Payment"
// @Success 200 {object} model.PaymentResponse
// @Failure 400
// @Router /api/payments [post]
func (cc *PaymentController) createNewPayment(ctx *gin.Context) {
	newPayment := new(model.Payment)

	if err := ctx.ShouldBindJSON(newPayment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to bind JSON",
			"error":   err.Error(),
		})
		return
	}

	if err := cc.paymentUseCase.Insert(newPayment); err != nil {
		var errorMessage = err.Error()

		if strings.Contains(err.Error(), "pq: insert or update on table") {
			errorMessage = "sender ID or receiver ID or bank account number does not exist"
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create new payment",
			"error":   errorMessage,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success create new payment",
	})
}

// @Summary Get payment by ID
// @Tags payment
// @Param id path int true "Payment ID"
// @Success 200 {object} model.Payment
// @Router /api/payments/{id} [get]
func (cc *PaymentController) getPaymentById(ctx *gin.Context) {
	paymentId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse payment ID",
			"error":   err.Error(),
		})
		return
	}

	payment, err := cc.paymentUseCase.PaymentById(paymentId)
	if err != nil {
		errorMessage := err.Error()

		if strings.Contains(err.Error(), "sql: no rows in result set") {
			errorMessage = "Payment with that ID does not exist"
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get payment",
			"error":   errorMessage,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get payment by ID",
		"payment": payment,
	})
}

func NewPaymentController(rg *gin.RouterGroup, paymentUseCase usecase.PaymentUseCase) *PaymentController {
	controller := PaymentController{
		rg:             rg,
		paymentUseCase: paymentUseCase,
	}

	controller.rg.GET("/payments", controller.getPayments)
	controller.rg.POST("/payments", controller.createNewPayment)
	controller.rg.GET("/payments/:id", controller.getPaymentById)

	return &controller
}
