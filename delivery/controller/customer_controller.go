package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-payment/model"
	"simple-payment/usecase"
	"strings"
)

type CustomerController struct {
	rg              *gin.RouterGroup
	customerUseCase usecase.CustomerUseCase
}

// @Summary Get all customers
// @Tags customer
// @success 200 {object} model.CustomerResponse{data=[]model.Customer}
// @Router /api/customers [get]
func (cc *CustomerController) getCustomers(ctx *gin.Context) {

	customers, err := cc.customerUseCase.Customers()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get customers",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get all customers",
		"data":    customers,
	})
}

// @Summary Create new customer
// @Tags customer
// @Param string body model.CustomerRequest true "Customer"
// @Success 200 {object} model.CustomerResponse
// @Failure 400
// @Router /api/customers [post]
func (cc *CustomerController) createNewCustomer(ctx *gin.Context) {
	newCustomer := new(model.Customer)

	if err := ctx.ShouldBindJSON(newCustomer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to bind JSON",
			"error":   err.Error(),
		})
		return
	}

	if err := cc.customerUseCase.Insert(newCustomer); err != nil {
		var errorMessage = err.Error()

		if strings.Contains(err.Error(), "pq: duplicate key value violates unique constraint") {
			errorMessage = "User with that ID already a customer"
		}

		if strings.Contains(err.Error(), "pq: insert or update on table") {
			errorMessage = "User with that ID does not exist, Please sign up first"
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create new customer",
			"error":   errorMessage,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success create new customer",
	})
}

// @Summary Get customer by ID
// @Tags customer
// @Success 200 {object} model.Customer
// @Router /api/customers/{id} [get]
func (cc *CustomerController) getCustomerById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":   "Get customer by ID",
		"customers": "",
	})
}

// @Summary Update customer by ID
// @Tags customer
// @Success 200 {object} model.Customer
// @Router /api/customers/id [put]
func (cc *CustomerController) updateCustomerById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":   "Update customer by ID",
		"customers": "",
	})
}

// @Summary Delete customer by ID
// @Tags customer
// @Success 200 {object} model.Customer
// @Router /api/customers/{id} [delete]
func (cc *CustomerController) deleteCustomerById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":   "Delete customer by ID",
		"customers": "",
	})
}

func NewCustomerController(rg *gin.RouterGroup, customerUseCase usecase.CustomerUseCase) *CustomerController {
	controller := CustomerController{
		rg:              rg,
		customerUseCase: customerUseCase,
	}

	controller.rg.GET("/customers", controller.getCustomers)
	controller.rg.POST("/customers", controller.createNewCustomer)
	controller.rg.GET("/customers/:id", controller.getCustomerById)
	controller.rg.PUT("/customers/:id", controller.updateCustomerById)
	controller.rg.DELETE("/customers/:id", controller.deleteCustomerById)

	return &controller
}
