package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	httpcommon "github.com/vucongthanh92/go-base-project/helper/http_common"
	"github.com/vucongthanh92/go-base-project/helper/validation"
	"github.com/vucongthanh92/go-base-project/internal/application/category"
	"github.com/vucongthanh92/go-base-project/internal/domain/models"
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
// @Summary create category
// @Accept json
// @Produce json
// @Param params body models.CreateCategoryReq true "CreateCategoryReq"
// @Router /api/v1/category [post]
// @Success	200
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	req := models.CreateCategoryReq{}
	if err := validation.GetBodyParamsHTTP(c, &req); err != nil {
		return
	}

	err := h.categoryService.CreateCategory(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(err.Error(), httpcommon.RequestInvalid, ""))
		return
	}

	c.JSON(http.StatusOK, nil)
}
