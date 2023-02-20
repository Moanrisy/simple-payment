package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BankController struct {
	rg *gin.RouterGroup
}

// @Summary Get all banks
// @Tags bank
// @Success 200 {object} model.Bank
// @Router /api/banks [get]
func (cc *BankController) getBanks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get all banks",
		"banks": "",
	})
}

// @Summary Create new bank
// @Tags bank
// @Success 200
// @Router /api/banks [post]
func (cc *BankController) createNewBank(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Create new bank",
		"banks": "",
	})
}

// @Summary Get bank by ID
// @Tags bank
// @Success 200 {object} model.Bank
// @Router /api/banks/{id} [get]
func (cc *BankController) getBankById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get bank by ID",
		"banks": "",
	})
}

// @Summary Update bank by ID
// @Tags bank
// @Success 200 {object} model.Bank
// @Router /api/banks/id [put]
func (cc *BankController) updateBankById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update bank by ID",
		"banks": "",
	})
}

// @Summary Delete bank by ID
// @Tags bank
// @Success 200 {object} model.Bank
// @Router /api/banks/{id} [delete]
func (cc *BankController) deleteBankById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete bank by ID",
		"banks": "",
	})
}

func NewBankController(rg *gin.RouterGroup) *BankController {
	controller := BankController{
		rg: rg,
	}

	controller.rg.GET("/banks", controller.getBanks)
	controller.rg.POST("/banks", controller.createNewBank)
	controller.rg.GET("/banks/:id", controller.getBankById)
	controller.rg.PUT("/banks/:id", controller.updateBankById)
	controller.rg.DELETE("/banks/:id", controller.deleteBankById)

	return & controller
}
