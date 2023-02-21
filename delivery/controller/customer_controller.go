package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-payment/model"
	"simple-payment/usecase"
	"time"
)

type CustomerController struct {
	rg              *gin.RouterGroup
	customerUseCase usecase.CustomerUseCase
}

// @Summary Get all customers
// @Tags customer
// @Success 200 {object} model.Customer
// @Router /api/customers [get]
func (cc *CustomerController) getCustomers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":   "Get all customers",
		"customers": "",
	})
}

// @Summary Create new customer
// @Tags customer
// @Success 200
// @Router /api/customers [post]
func (cc *CustomerController) createNewCustomer(ctx *gin.Context) {
	newCustomer := model.Customer{
		CustomerId: "1",
		UserId:     "1",
		Name:       "John Doe",
		Balance:    "100000",
		CreatedAt:  time.Now(),
	}

	if err := cc.customerUseCase.Insert(&newCustomer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create new customer",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":   "Create new customer",
		"customers": "",
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
