package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vucongthanh92/go-base-project/internal/application/product"
)

type ProductHandler struct {
	productService product.ProductService
}

func NewProductHandler(
	productService product.ProductService,
) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// API get products list godoc
// @Tags Product
// @Summary search products with filter and return pagination
// @Accept json
// @Produce json
// @Param  params body entities.HomeMovingEstimateRequest true "ProductResponse List"
// @Router 	/api/v1/products [get]
// @Success	200 {object} httpcommon.SuccessResponse[entities.ResultOrder]
func (h *ProductHandler) GetProductList(c *gin.Context) {

	c.JSON(http.StatusOK, nil)
}
