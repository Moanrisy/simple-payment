package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-payment/model"
	"simple-payment/usecase"
	"strconv"
	"strings"
)

type BankController struct {
	rg          *gin.RouterGroup
	bankUseCase usecase.BankUseCase
}

// @Summary Get all banks
// @Tags 3. Bank endpoints
// @success 200 {object} model.BankResponse{data=[]model.Bank}
// @Router /api/banks [get]
func (cc *BankController) getBanks(ctx *gin.Context) {

	banks, err := cc.bankUseCase.Banks()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get banks",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get all banks",
		"banks":   banks,
	})
}

// @Summary Create new bank
// @Tags 3. Bank endpoints
// @Param string body model.BankRequest true "Bank"
// @Success 200 {object} model.BankResponse
// @Failure 400
// @Router /api/banks [post]
func (cc *BankController) createNewBank(ctx *gin.Context) {
	newBank := new(model.Bank)

	if err := ctx.ShouldBindJSON(newBank); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to bind JSON",
			"error":   err.Error(),
		})
		return
	}

	if err := cc.bankUseCase.Insert(newBank); err != nil {
		var errorMessage = err.Error()

		if strings.Contains(err.Error(), "pq: duplicate key value violates unique constraint") {
			errorMessage = "Bank with that account number already exist"
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create new bank",
			"error":   errorMessage,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success create new bank",
	})
}

// @Summary Get bank by ID
// @Tags 3. Bank endpoints
// @Param id path int true "Bank ID"
// @Success 200 {object} model.Bank
// @Router /api/banks/{id} [get]
func (cc *BankController) getBankById(ctx *gin.Context) {
	bankId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse bank ID",
			"error":   err.Error(),
		})
		return
	}

	bank, err := cc.bankUseCase.BankById(bankId)
	if err != nil {
		errorMessage := err.Error()

		if strings.Contains(err.Error(), "sql: no rows in result set") {
			errorMessage = "Bank with that ID does not exist"
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get bank",
			"error":   errorMessage,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get bank by ID",
		"banks":   bank,
	})
}

// @Summary Delete bank by ID
// @Tags 3. Bank endpoints
// @Param id path int true "Bank ID"
// @Success 200
// @Router /api/banks/{id} [delete]
func (cc *BankController) deleteBankById(ctx *gin.Context) {
	bankId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse bank ID",
			"error":   err.Error(),
		})
		return
	}

	if err := cc.bankUseCase.Delete(bankId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to delete bank",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete bank success",
	})
}

func NewBankController(rg *gin.RouterGroup, bankUseCase usecase.BankUseCase) *BankController {
	controller := BankController{
		rg:          rg,
		bankUseCase: bankUseCase,
	}

	controller.rg.GET("/banks", controller.getBanks)
	controller.rg.POST("/banks", controller.createNewBank)
	controller.rg.GET("/banks/:id", controller.getBankById)
	controller.rg.DELETE("/banks/:id", controller.deleteBankById)

	return &controller
}
