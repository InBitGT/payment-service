package handlers

import (
	"net/http"
	"payment-service/internal/model"
	"payment-service/internal/services"
	"strconv"

	"github.com/InBitGT/senti-shared-library/pkg"
	"github.com/InBitGT/senti-shared-library/pkg/common"
	"github.com/gin-gonic/gin"
)

type SuscriptionHandler struct {
	service services.SuscriptionService
}

func NewSuscriptionHandler(service services.SuscriptionService) *SuscriptionHandler {
	return &SuscriptionHandler{service: service}
}

func (h *SuscriptionHandler) Create(c *gin.Context) {
	var data model.Suscription
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse(400, pkg.ERR_VALIDATION, "Datos inválidos", nil))
		return
	}
	if err := h.service.Create(&data); err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse(500, pkg.ERR_INTERNAL_ERROR, err.Error(), nil))
		return
	}
	c.JSON(http.StatusCreated, common.SuccessResponse(pkg.SUCCESS_CREATED, data, pkg.HTTP_CREATED))
}

func (h *SuscriptionHandler) GetAll(c *gin.Context) {
	roles, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse(404, pkg.ERR_NOT_FOUND, "Roles no encontrados", nil))
		return
	}
	c.JSON(http.StatusOK, common.SuccessResponse(pkg.SUCCESS_RETRIEVED, roles, pkg.HTTP_OK))
}

func (h *SuscriptionHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse(400, pkg.ERR_VALIDATION, "ID inválido", nil))
		return
	}
	data, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse(404, pkg.ERR_NOT_FOUND, "data no encontrado", nil))
		return
	}
	c.JSON(http.StatusOK, common.SuccessResponse(pkg.SUCCESS_RETRIEVED, data, pkg.HTTP_OK))
}

func (h *SuscriptionHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse(400, pkg.ERR_VALIDATION, "ID inválido", nil))
		return
	}
	var data model.Suscription
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse(400, pkg.ERR_VALIDATION, "Datos inválidos", nil))
		return
	}
	data.ID = uint(id)
	err = h.service.Update(&data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse(500, pkg.ERR_INTERNAL_ERROR, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.SuccessResponse(pkg.SUCCESS_UPDATED, "exito", pkg.HTTP_OK))
}

func (h *SuscriptionHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse(400, pkg.ERR_VALIDATION, "ID inválido", nil))
		return
	}
	err = h.service.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse(404, pkg.ERR_NOT_FOUND, "Data no encontrado", nil))
		return
	}
	c.JSON(http.StatusOK, common.SuccessResponse(pkg.SUCCESS_DELETED, "exito", pkg.HTTP_OK))
}
