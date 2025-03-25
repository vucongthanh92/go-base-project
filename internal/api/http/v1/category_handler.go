package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vucongthanh92/go-base-project/helper/constants"
	httpcommon "github.com/vucongthanh92/go-base-project/helper/http_common"
	"github.com/vucongthanh92/go-base-project/helper/validation"
	"github.com/vucongthanh92/go-base-project/internal/domain/interfaces"
	"github.com/vucongthanh92/go-base-project/internal/domain/models"
)

type CategoryHandler struct {
	categoryService interfaces.CategoryServiceI
}

func NewCategoryHandler(categoryService interfaces.CategoryServiceI) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

// API CreateCategory godoc
// @Tags Category
// @Summary create category by name
// @Accept json
// @Produce json
// @Param params body models.CreateCategoryReq true "CreateCategoryReq"
// @Router /api/v1/category [post]
// @Success	200 {object} entities.Category
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	req := models.CreateCategoryReq{}

	if err := validation.GetBodyParamsHTTP(c, &req); err != nil {
		return
	}

	res, resErr := h.categoryService.CreateCategory(c, req)
	if resErr != nil {
		resErr.ExposeHttpError(c)
		return
	}

	c.JSON(http.StatusOK, httpcommon.NewSuccessResponse(res))
}

// API UpdateCategory godoc
// @Tags Category
// @Summary update category by id
// @Accept json
// @Produce json
// @Param params body models.CreateCategoryReq true "UpdateCategoryReq"
// @Router /api/v1/category [put]
// @Success	200 {object} entities.Category
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	var (
		req = models.UpdateCategoryReq{}
		err error
	)

	if err := validation.GetBodyParamsHTTP(c, &req); err != nil {
		return
	}

	objectID := c.Param("id")
	req.ID, err = strconv.ParseUint(objectID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse("Invalid ID", constants.REQUEST_INVALID, ""))
		return
	}

	res, resErr := h.categoryService.UpdateCategory(c, req)
	if resErr != nil {
		resErr.ExposeHttpError(c)
		return
	}

	c.JSON(http.StatusOK, httpcommon.NewSuccessResponse(res))
}

// API DeleteCategoryByID godoc
// @Tags Category
// @Summary delete category by id
// @Accept json
// @Produce json
// @Param params body models.CreateCategoryReq true "UpdateCategoryReq"
// @Router /api/v1/category [put]
// @Success	200 {object} entities.Category
func (h *CategoryHandler) DeleteCategoryByID(c *gin.Context) {
	var req = struct {
		UpdatedAt time.Time `json:"updated_at"`
	}{}

	if err := validation.GetBodyParamsHTTP(c, &req); err != nil {
		return
	}

	objectID := c.Param("id")
	categoryID, err := strconv.ParseUint(objectID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse("Invalid ID", constants.REQUEST_INVALID, ""))
		return
	}

	resErr := h.categoryService.DeleteCategoryByID(c, categoryID, req.UpdatedAt)
	if resErr != nil {
		resErr.ExposeHttpError(c)
		return
	}

	c.JSON(http.StatusOK, httpcommon.NewSuccessResponse[any](nil))
}

// API GetCategoryList godoc
// @Tags Category
// @Summary get list categories
// @Accept json
// @Produce json
// @Router /api/v1/category [get]
// @Success 200 {object} []entities.Category
func (h *CategoryHandler) GetCategoryList(c *gin.Context) {

	res, resErr := h.categoryService.GetCategoryList(c)
	if resErr != nil {
		resErr.ExposeHttpError(c)
		return
	}

	c.JSON(http.StatusOK, httpcommon.NewSuccessResponse(res))
}
