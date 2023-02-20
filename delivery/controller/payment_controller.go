package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PaymentController struct {
	rg *gin.RouterGroup
}

// @Summary Get all payments
// @Tags payment
// @Success 200 {object} model.Payment
// @Router /api/payments [get]
func (cc *PaymentController) getPayments(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get all payments",
		"payments": "",
	})
}

// @Summary Create new payment
// @Tags payment
// @Success 200
// @Router /api/payments [post]
func (cc *PaymentController) createNewPayment(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Create new payment",
		"payments": "",
	})
}

// @Summary Get payment by ID
// @Tags payment
// @Success 200 {object} model.Payment
// @Router /api/payments/{id} [get]
func (cc *PaymentController) getPaymentById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get payment by ID",
		"payments": "",
	})
}

func NewPaymentController(rg *gin.RouterGroup) *PaymentController {
	controller := PaymentController{
		rg: rg,
	}

	controller.rg.GET("/payments", controller.getPayments)
	controller.rg.POST("/payments", controller.createNewPayment)
	controller.rg.GET("/payments/:id", controller.getPaymentById)

	return & controller
}
