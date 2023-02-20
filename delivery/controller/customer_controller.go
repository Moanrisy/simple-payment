package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CustomerController struct {
	rg *gin.RouterGroup
}

func (cc *CustomerController) getCustomers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get all customers",
		"customers": "",
	})
}


func (cc *CustomerController) createNewCustomer(ctx *gin.Context) {
ctx.JSON(http.StatusOK, gin.H{
		"message": "Create new customer",
		"customers": "",
	})
}

func (cc *CustomerController) getCustomerById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get customer by ID",
		"customers": "",
	})
}

func (cc *CustomerController) updateCustomerById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update customer by ID",
		"customers": "",
	})
}

func (cc *CustomerController) deleteCustomerById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete customer by ID",
		"customers": "",
	})
}

func NewCustomerController(rg *gin.RouterGroup) *CustomerController {
	controller := CustomerController{
		rg: rg,
	}

	controller.rg.GET("/customers", controller.getCustomers)
	controller.rg.POST("/customers", controller.createNewCustomer)
	controller.rg.GET("/customers/:id", controller.getCustomerById)
	controller.rg.PUT("/customers/:id", controller.updateCustomerById)
	controller.rg.DELETE("/customers/:id", controller.deleteCustomerById)

	return & controller
}