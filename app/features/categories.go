package features

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/k23dev/tango/app/models"
	"github.com/k23dev/tango/app/views"
	"github.com/k23dev/tango/pkg/webcore"
	"github.com/k23dev/tango/pkg/webcore/utils"
	"github.com/labstack/echo/v4"
)

func FindOneCategory(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))

	cat := models.NewCategory()
	category, _ := cat.FindOne(tapp.App.DB.Primary, id)
	if category.ID != 0 {
		return utils.Render(c, views.CategoriesShowOne(tapp.GetTitleAndVersion(), *category))
	} else {
		return utils.RenderNotFound(c)
	}
}

func FindAllCategories(c echo.Context, tapp *webcore.TangoApp) error {
	cat := models.NewCategory()
	categories, _ := cat.FindAll(tapp.App.DB.Primary)
	return utils.Render(c, views.CategoriesShowList(tapp.GetTitleAndVersion(), categories))
}

func ShowFormCategory(c echo.Context, tapp *webcore.TangoApp, is_new bool) error {
	cat := models.NewCategory()

	if is_new {
		return utils.Render(c, views.CategoriesFormCreate(tapp.GetTitleAndVersion()))
	} else {
		id, _ := strconv.Atoi(c.Param("id"))
		cat, _ := cat.FindOne(tapp.App.DB.Primary, id)
		return utils.Render(c, views.CategoriesFormUpdate(tapp.GetTitleAndVersion(), cat))
	}
}

func CreateCategory(c echo.Context, tapp *webcore.TangoApp) error {
	// get the incoming values
	catDTO := models.CategoryDTO{}
	if err := c.Bind(&catDTO); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	cat := models.NewCategory()
	cat.Create(tapp.App.DB.Primary, catDTO.Name)
	// return c.String(http.StatusOK, "Categoría creada "+category.Name)
	return c.Redirect(http.StatusMovedPermanently, "/categories/")
}

func UpdateCategory(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// get the incoming values
	catDTO := models.CategoryDTO{}
	if err := c.Bind(&catDTO); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	cat := models.NewCategory()
	cat.Name = strings.ToLower(catDTO.Name)

	cat.Update(tapp.App.DB.Primary, id, cat.Name)
	// return c.String(http.StatusOK, "Categoría actualizada "+cat.Name)
	return c.Redirect(http.StatusMovedPermanently, "/categories/")
}

func DeleteCategory(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))
	cat := models.NewCategory()
	cat.Delete(tapp.App.DB.Primary, id)
	// return c.String(http.StatusOK, "Categoría creada "+category.Name)
	return c.Redirect(http.StatusMovedPermanently, "/categories/")
}
