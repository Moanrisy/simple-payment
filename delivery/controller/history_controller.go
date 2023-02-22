package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-payment/delivery/middleware"
	"simple-payment/usecase"
)

type HistoryController struct {
	rg             *gin.RouterGroup
	historyUseCase usecase.HistoryUseCase
	tokenMdw       middleware.AuthTokenMiddleware
}

// @Summary Get all histories
// @Tags logs history
// @success 200 {object} model.LogHistoryResponse{data=[]model.LogHistory}
// @Router /api/logs/history [get]
func (cc *HistoryController) getHistorys(ctx *gin.Context) {

	histories, err := cc.historyUseCase.Historys()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get histories",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":   "Get all histories",
		"histories": histories,
	})
}

func NewHistoryController(rg *gin.RouterGroup, historyUseCase usecase.HistoryUseCase, tokenMdw middleware.AuthTokenMiddleware) *HistoryController {
	controller := HistoryController{
		rg:             rg,
		historyUseCase: historyUseCase,
		tokenMdw:       tokenMdw,
	}

	protectedGroup := rg.Group("", tokenMdw.RequiredToken())

	protectedGroup.GET("/logs/history", controller.getHistorys)

	return &controller
}
