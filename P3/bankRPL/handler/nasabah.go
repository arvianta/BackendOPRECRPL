package handler

import (
	"net/http"

	"github.com/arvianta/BackendOPRECRPL/P3/bankRPL/common"
	"github.com/arvianta/BackendOPRECRPL/P3/bankRPL/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NasabahHandler struct {
	DB *gorm.DB
}

func (h *NasabahHandler) HandleInsertNasabah(ctx *gin.Context) {
	var nasabah entity.Nasabah

	err := ctx.ShouldBind(&nasabah)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: err.Error(),
				Data:    nil,
			})
		return
	}
	tx := h.DB.Create(&nasabah)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: tx.Error.Error(),
				Data:    nil,
			})
		return
	}

	ctx.JSON(http.StatusOK,
		common.Response{
			Status:  true,
			Message: "nasabah created successfully",
			Data:    nasabah,
		})
}

func (h *NasabahHandler) HandleUpdateNasabah(ctx *gin.Context) {
	id := ctx.Param("id")

	var nasabah entity.Nasabah

	err := ctx.ShouldBind(&nasabah)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: err.Error(),
				Data:    nil,
			})
		return
	}

	tx := h.DB.Model(&nasabah).Where("id = ?", id).Updates(&nasabah)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: tx.Error.Error(),
				Data:    nil,
			})
		return
	}

	ctx.JSON(http.StatusOK,
		common.Response{
			Status:  true,
			Message: "nasabah updated successfully",
			Data:    nasabah,
		})
}

func (h *NasabahHandler) HandleDeleteNasabah(ctx *gin.Context) {
	var nasabah entity.Nasabah

	err := h.DB.First(&nasabah, ctx.Param("id")).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: "failed to find nasabah with given ID",
				Data:    nil,
			})
		return
	}

	tx := h.DB.Delete(&nasabah)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: tx.Error.Error(),
				Data:    nil,
			})
		return
	}

	ctx.JSON(http.StatusOK,
		common.Response{
			Status:  true,
			Message: "nasabah deleted successfully",
			Data:    nil,
		})
}
