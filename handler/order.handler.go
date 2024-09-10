package handler

import (
	"github.com/affrianr/gin-rest-api-h8/domain"
	"github.com/affrianr/gin-rest-api-h8/usecase"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderUsecase *usecase.OrderUsecase
}

func NewOrderHandler(orderUC *usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{orderUC}
}

func (handler *OrderHandler) CreateOrder(ctx *gin.Context) {
	var order domain.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"success": false,
			"message": err.Error(),
		})
		return
	}

	if err := handler.orderUsecase.CreateOrder(&order); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"success": false,
			"message": "Failed to create order",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"success": true,
		"data":    order,
	})
}

func (handler *OrderHandler) GetOrder(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	order, err := handler.orderUsecase.GetOrder(uint(id))

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"success": false,
			"message": "Order not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"success": true,
		"data":    order,
	})
}

func (handler *OrderHandler) UpdateOrder(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var order domain.Order

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"success": false,
			"message": err.Error(),
		})
		return
	}

	order.ID = uint(id)

	if err := handler.orderUsecase.UpdateOrder(&order); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"success": false,
			"message": "Failed to update order",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"success": true,
		"data":    order,
	})
}

func (handler *OrderHandler) DeleteOrder(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := handler.orderUsecase.DeleteOrder(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"success": false,
			"message": "Failed to delete order",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"success": true,
		"message": "Order deleted successfully",
	})
}

func (handler *OrderHandler) ListOrders(ctx *gin.Context) {
	order, err := handler.orderUsecase.ListOrders()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"success": false,
			"message": "Failed to fetch orders",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"success": true,
		"data":    order,
	})
}
