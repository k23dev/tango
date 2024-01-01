
package features

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/k23dev/tango/app/models"
	"github.com/k23dev/tango/app/views"
	"github.com/k23dev/tango/pkg/pagination"
	"github.com/k23dev/tango/pkg/webcore"
	"github.com/k23dev/tango/pkg/webcore/utils"
	"github.com/labstack/echo/v4"
)


func FindOneTanga(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))

	t := models.NewTanga()
	tanga, _ := t.FindOne(tapp.App.DB.Primary, id)
	if tanga != nil {
		return utils.Render(c, views.TangasShowOne(tapp.GetTitleAndVersion(), *tanga))
	} else {
		return c.Redirect(http.StatusMovedPermanently, "/404")
	}
}

func FindAllTangas(c echo.Context, tapp *webcore.TangoApp) error {
	queryPage := c.Param("page")
	var currentPage = 0
	if queryPage != "" {
		currentPage, _ = strconv.Atoi(queryPage)
	}

	t := models.NewTanga()
	counter, _ := t.Count(tapp.App.DB.Primary)
	pagination := pagination.NewPagination(currentPage, itemsPerPage, counter)
	t, _ := t.FindAllPagination(tapp.App.DB.Primary, itemsPerPage, currentPage)

	return utils.Render(c, views.TangasShowList(tapp.GetTitleAndVersion(), *t, *pagination))
}

func ShowFormTanga(c echo.Context, tapp *webcore.TangoApp, is_new bool) error {
	t := models.NewTanga()

	if is_new {
		return utils.Render(c, views.TangasFormCreate(tapp.GetTitleAndVersion()))
	} else {
		id, _ := strconv.Atoi(c.Param("id"))
		t, _ := t.FindOne(tapp.App.DB.Primary, id)
		return utils.Render(c, views.TangasFormUpdate(tapp.GetTitleAndVersion(), t))
	}
}

func CreateTanga(c echo.Context, tapp *webcore.TangoApp) error {
	// get the incoming values
	tDTO := models.TangaDTO{}
	if err := c.Bind(&tDTO); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	t := models.NewTanga()
	t.Create(tapp.App.DB.Primary, tDTO.Name)

	return c.Redirect(http.StatusMovedPermanently, "/tangas/")
}

func UpdateTanga(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// get the incoming values
	tDTO := models.TangaDTO{}
	if err := c.Bind(&tDTO); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	t := models.NewTanga()
	t.Name = strings.ToLower(tDTO.Name)

	t.Update(tapp.App.DB.Primary, id, t.Name)

	return c.Redirect(http.StatusMovedPermanently, "/tangas/")
}

func DeleteTanga(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))
	t := models.NewTanga()
	t.Delete(tapp.App.DB.Primary, id)

	return c.Redirect(http.StatusMovedPermanently, "/tangas/")
}
	