package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-payment/model"
	"simple-payment/usecase"
	"strconv"
	"strings"
)

type MerchantController struct {
	rg              *gin.RouterGroup
	merchantUseCase usecase.MerchantUseCase
}

// @Summary Get all merchants
// @Tags merchant
// @success 200 {object} model.MerchantResponse{data=[]model.Merchant}
// @Router /api/merchants [get]
func (cc *MerchantController) getMerchants(ctx *gin.Context) {

	merchants, err := cc.merchantUseCase.Merchants()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get merchants",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":   "Get all merchants",
		"merchants": merchants,
	})
}

// @Summary Create new merchant
// @Tags merchant
// @Param string body model.MerchantRequest true "Merchant"
// @Success 200 {object} model.MerchantResponse
// @Failure 400
// @Router /api/merchants [post]
func (cc *MerchantController) createNewMerchant(ctx *gin.Context) {
	newMerchant := new(model.Merchant)

	if err := ctx.ShouldBindJSON(newMerchant); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to bind JSON",
			"error":   err.Error(),
		})
		return
	}

	if err := cc.merchantUseCase.Insert(newMerchant); err != nil {
		var errorMessage = err.Error()

		if strings.Contains(err.Error(), "pq: duplicate key value violates unique constraint") {
			errorMessage = "User with that ID already have a merchant"
		}

		if strings.Contains(err.Error(), "pq: insert or update on table") {
			errorMessage = "User with that ID does not exist, Please sign up first"
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create new merchant",
			"error":   errorMessage,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success create new merchant",
	})
}

// @Summary Get merchant by ID
// @Tags merchant
// @Param id path int true "Merchant ID"
// @Success 200 {object} model.Merchant
// @Router /api/merchants/{id} [get]
func (cc *MerchantController) getMerchantById(ctx *gin.Context) {
	merchantId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse merchant ID",
			"error":   err.Error(),
		})
		return
	}

	merchant, err := cc.merchantUseCase.MerchantById(merchantId)
	if err != nil {
		errorMessage := err.Error()

		if strings.Contains(err.Error(), "sql: no rows in result set") {
			errorMessage = "Merchant with that ID does not exist"
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get merchant",
			"error":   errorMessage,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Get merchant by ID",
		"customer": merchant,
	})
}

// @Summary Delete merchant by ID
// @Tags merchant
// @Param id path int true "Merchant ID"
// @Success 200 {object} model.Merchant
// @Router /api/merchants/{id} [delete]
func (cc *MerchantController) deleteMerchantById(ctx *gin.Context) {
	merchantId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse merchant ID",
			"error":   err.Error(),
		})
		return
	}

	if err := cc.merchantUseCase.Delete(merchantId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to delete merchant",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete merchant success",
	})
}

func NewMerchantController(rg *gin.RouterGroup, merchantUseCase usecase.MerchantUseCase) *MerchantController {
	controller := MerchantController{
		rg:              rg,
		merchantUseCase: merchantUseCase,
	}

	controller.rg.GET("/merchants", controller.getMerchants)
	controller.rg.POST("/merchants", controller.createNewMerchant)
	controller.rg.GET("/merchants/:id", controller.getMerchantById)
	controller.rg.DELETE("/merchants/:id", controller.deleteMerchantById)

	return &controller
}
