package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vucongthanh92/go-base-project/internal/application/supplier"
)

type SupplierHandler struct {
	supplierService supplier.SupllierService
}

func NewSupplierHandler(
	supplierService supplier.SupllierService,
) *SupplierHandler {
	return &SupplierHandler{
		supplierService: supplierService,
	}
}

// API get products list godoc
// @Tags Supplier
// @Summary search products with filter and return pagination
// @Accept json
// @Produce json
// @Param  params body entities.HomeMovingEstimateRequest true "ProductResponse List"
// @Router 	/api/v1/products [get]
// @Success	200 {object} httpcommon.SuccessResponse[entities.ResultOrder]
func (h *ProductHandler) CreateSupplier(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
