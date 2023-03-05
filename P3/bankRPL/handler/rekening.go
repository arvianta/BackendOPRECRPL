package handler

import (
	"net/http"

	"github.com/arvianta/BackendOPRECRPL/P3/bankRPL/common"
	"github.com/arvianta/BackendOPRECRPL/P3/bankRPL/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RekeningHandler struct {
	DB  *gorm.DB
	DBA *gorm.Association
}

func (h *RekeningHandler) HandleInsertRekening(ctx *gin.Context) {
	var rekening entity.Rekening
	err := ctx.ShouldBind(&rekening)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: err.Error(),
				Data:    nil,
			})
		return
	}

	tx := h.DB.Create(&rekening)
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
			Message: "Rekening berhasil didaftar",
			Data:    rekening,
		})
}

func (h *RekeningHandler) HandleGetAllRekeningByNama(ctx *gin.Context) {
	nama := ctx.Param("nama")

	var nasabah entity.Nasabah
	tx := h.DB.Preload("Rekening").Where("nama = ?", nama).First(&nasabah)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: tx.Error.Error(),
				Data:    nil,
			})
		return
	}

	rekening := make([]string, len(nasabah.Rekening))
	for i, r := range nasabah.Rekening {
		rekening[i] = r.Number
	}

	ctx.JSON(http.StatusOK,
		common.Response{
			Status:  true,
			Message: "nasabah rekenings fetched successfully",
			Data:    rekening,
		})
}

func (h *RekeningHandler) HandleDeleteNasabahRekening(ctx *gin.Context) {
	nama := ctx.Param("nama")
	rekeningNumber := ctx.Param("rekeningNumber")

	var nasabah entity.Nasabah
	tx := h.DB.Preload("Rekening").Where("nama = ?", nama).First(&nasabah)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: tx.Error.Error(),
				Data:    nil,
			})
		return
	}

	var rekeningToDelete entity.Rekening
	found := false
	for _, r := range nasabah.Rekening {
		if r.Number == rekeningNumber {
			rekeningToDelete = r
			found = true
			break
		}
	}
	if !found {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: "rekening number is not found for nasabah",
				Data:    nil,
			})
		return
	}

	tx2 := h.DB.Model(&nasabah).Association("Rekening").Delete(rekeningToDelete)
	if tx2 != nil {
		ctx.JSON(http.StatusInternalServerError,
			common.Response{
				Status:  false,
				Message: "failed to delete rekening for nasabah",
				Data:    nil,
			})
		return
	}

	ctx.JSON(http.StatusOK,
		common.Response{
			Status:  true,
			Message: "rekening deleted successfully",
			Data:    nil,
		})
}
