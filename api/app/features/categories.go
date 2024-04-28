package features

import (
	"net/http"
	"strconv"

	"github.com/k23dev/tango/app/models"
	"github.com/k23dev/tango/pkg/webcore"
	"github.com/labstack/echo/v4"
)

const categoriesPagination = false

const categoriesPaginationItemsPerPage = 15

func FindOneCategory(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	c := models.NewCategory()
	category, err := c.FindOne(tapp.App.DB.Primary, id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err)
	}
	return ctx.JSON(http.StatusOK, category.ConvertToDTO())
}

func FindAllCategories(ctx echo.Context, tapp *webcore.TangoApp) error {
	var cBuf *[]models.Category
	c := models.NewCategory()

	if categoriesPagination == true {
		queryPage := ctx.Param("page")
		currentPage := 0
		if queryPage != "" {
			currentPage, _ = strconv.Atoi(queryPage)
		}

		// total de registros en la db
		// counter, _ := c.Count(tapp.App.DB.Primary)
		// pagination := pagination.NewPagination(currentPage,categoriesPaginationItemsPerPage,counter)

		cBuf, _ = c.FindAllPagination(tapp.App.DB.Primary, categoriesPaginationItemsPerPage, currentPage)
	} else {
		cBuf, _ = c.FindAll(tapp.App.DB.Primary)
	}

	return ctx.JSON(http.StatusOK, cBuf)

}

func CreateCategory(ctx echo.Context, tapp *webcore.TangoApp) error {
	// get the incoming values
	cDTO := models.CategoryDTO{}
	if err := ctx.Bind(&cDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, "")
	}

	c := models.NewCategory()
	cBuf, err := c.Create(tapp.App.DB.Primary, cDTO)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusCreated, cBuf.ConvertToDTO())
}

func UpdateCategory(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	// get the incoming values
	cDTO := models.CategoryDTO{}
	if err := ctx.Bind(&cDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, "")
	}

	c := models.NewCategory()
	cBuf, err := c.Update(tapp.App.DB.Primary, id, cDTO)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, cBuf.ConvertToDTO())
}

func DeleteCategory(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	c := models.NewCategory()
	cBuf, err := c.Delete(tapp.App.DB.Primary, id)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, cBuf.ConvertToDTO())
}
