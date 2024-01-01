
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


func FindOneFoto(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))

	f := models.NewFoto()
	foto, _ := f.FindOne(tapp.App.DB.Primary, id)
	if foto != nil {
		return utils.Render(c, views.FotosShowOne(tapp.GetTitleAndVersion(), *foto))
	} else {
		return c.Redirect(http.StatusMovedPermanently, "/404")
	}
}

func FindAllFotos(c echo.Context, tapp *webcore.TangoApp) error {
	queryPage := c.Param("page")
	var currentPage = 0
	if queryPage != "" {
		currentPage, _ = strconv.Atoi(queryPage)
	}

	f := models.NewFoto()
	counter, _ := f.Count(tapp.App.DB.Primary)
	pagination := pagination.NewPagination(currentPage, itemsPerPage, counter)
	f, _ := f.FindAllPagination(tapp.App.DB.Primary, itemsPerPage, currentPage)

	return utils.Render(c, views.FotosShowList(tapp.GetTitleAndVersion(), *f, *pagination))
}

func ShowFormFoto(c echo.Context, tapp *webcore.TangoApp, is_new bool) error {
	f := models.NewFoto()

	if is_new {
		return utils.Render(c, views.FotosFormCreate(tapp.GetTitleAndVersion()))
	} else {
		id, _ := strconv.Atoi(c.Param("id"))
		f, _ := f.FindOne(tapp.App.DB.Primary, id)
		return utils.Render(c, views.FotosFormUpdate(tapp.GetTitleAndVersion(), f))
	}
}

func CreateFoto(c echo.Context, tapp *webcore.TangoApp) error {
	// get the incoming values
	fDTO := models.FotoDTO{}
	if err := c.Bind(&fDTO); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	f := models.NewFoto()
	f.Create(tapp.App.DB.Primary, fDTO.Name)

	return c.Redirect(http.StatusMovedPermanently, "/fotos/")
}

func UpdateFoto(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// get the incoming values
	fDTO := models.FotoDTO{}
	if err := c.Bind(&fDTO); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	f := models.NewFoto()
	f.Name = strings.ToLower(fDTO.Name)

	f.Update(tapp.App.DB.Primary, id, f.Name)

	return c.Redirect(http.StatusMovedPermanently, "/fotos/")
}

func DeleteFoto(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))
	f := models.NewFoto()
	f.Delete(tapp.App.DB.Primary, id)

	return c.Redirect(http.StatusMovedPermanently, "/fotos/")
}
	