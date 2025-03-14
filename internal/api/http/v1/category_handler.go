package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vucongthanh92/go-base-project/internal/application/category"
)

type CategoryHandler struct {
	categoryService category.CategoryService
}

func NewCategoryHandler(
	categoryService category.CategoryService,
) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

// API get products list godoc
// @Tags Category
// @Summary search products with filter and return pagination
// @Accept json
// @Produce json
// @Param  params body entities.HomeMovingEstimateRequest true "ProductResponse List"
// @Router 	/api/v1/products [get]
// @Success	200 {object} httpcommon.SuccessResponse[entities.ResultOrder]
func (h *CategoryHandler) CreateCategory(c *gin.Context) {

	c.JSON(http.StatusOK, nil)
}
