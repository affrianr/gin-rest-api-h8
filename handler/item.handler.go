package handler

import (
	"net/http"
	"strconv"

	"github.com/affrianr/gin-rest-api-h8/domain"
	"github.com/affrianr/gin-rest-api-h8/usecase"

	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	ItemUsecase *usecase.ItemUsecase
}

func NewItemHandler(itemUC *usecase.ItemUsecase) *ItemHandler {
	return &ItemHandler{itemUC}
}

func (handler *ItemHandler) CreateItem(ctx *gin.Context) {
	var item domain.Item
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"success": false,
			"message": err.Error(),
		})
		return
	}

	if err := handler.ItemUsecase.CreateItem(&item); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"success": false,
			"message": "Failed to create item",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"success": true,
		"data":    item,
	})
}

func (handler *ItemHandler) GetItem(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	item, err := handler.ItemUsecase.GetItem(uint(id))

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"success": false,
			"message": "Item not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"success": true,
		"data":    item,
	})
}

func (handler *ItemHandler) UpdateItem(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var item domain.Item

	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"success": false,
			"message": err.Error(),
		})
		return
	}

	item.ID = uint(id)

	if err := handler.ItemUsecase.UpdateItem(&item); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"success": false,
			"message": "Failed to update item",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"success": true,
		"data":    item,
	})
}

func (handler *ItemHandler) DeleteItem(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := handler.ItemUsecase.DeleteItem(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"success": false,
			"message": "Failed to delete item",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"success": true,
		"message": "Item deleted successfully",
	})
}

func (handler *ItemHandler) ListItems(ctx *gin.Context) {
	item, err := handler.ItemUsecase.ListItems()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"success": false,
			"message": "Failed to fetch items",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"success": true,
		"data":    item,
	})
}
