package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MerchantController struct {
	rg *gin.RouterGroup
}

// @Summary Get all merchants
// @Tags merchant
// @Success 200 {object} model.Merchant
// @Router /api/merchants [get]
func (cc *MerchantController) getMerchants(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get all merchants",
		"merchants": "",
	})
}

// @Summary Create new merchant
// @Tags merchant
// @Success 200
// @Router /api/merchants [post]
func (cc *MerchantController) createNewMerchant(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Create new merchant",
		"merchants": "",
	})
}

// @Summary Get merchant by ID
// @Tags merchant
// @Success 200 {object} model.Merchant
// @Router /api/merchants/{id} [get]
func (cc *MerchantController) getMerchantById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get merchant by ID",
		"merchants": "",
	})
}

// @Summary Update merchant by ID
// @Tags merchant
// @Success 200 {object} model.Merchant
// @Router /api/merchants/id [put]
func (cc *MerchantController) updateMerchantById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update merchant by ID",
		"merchants": "",
	})
}

// @Summary Delete merchant by ID
// @Tags merchant
// @Success 200 {object} model.Merchant
// @Router /api/merchants/{id} [delete]
func (cc *MerchantController) deleteMerchantById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete merchant by ID",
		"merchants": "",
	})
}

func NewMerchantController(rg *gin.RouterGroup) *MerchantController {
	controller := MerchantController{
		rg: rg,
	}

	controller.rg.GET("/merchants", controller.getMerchants)
	controller.rg.POST("/merchants", controller.createNewMerchant)
	controller.rg.GET("/merchants/:id", controller.getMerchantById)
	controller.rg.PUT("/merchants/:id", controller.updateMerchantById)
	controller.rg.DELETE("/merchants/:id", controller.deleteMerchantById)

	return & controller
}
