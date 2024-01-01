
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


func FindOneGato(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))

	g := models.NewGato()
	gato, _ := g.FindOne(tapp.App.DB.Primary, id)
	if gato != nil {
		return utils.Render(c, views.GatosShowOne(tapp.GetTitleAndVersion(), *gato))
	} else {
		return c.Redirect(http.StatusMovedPermanently, "/404")
	}
}

func FindAllGatos(c echo.Context, tapp *webcore.TangoApp) error {
	queryPage := c.Param("page")
	var currentPage = 0
	if queryPage != "" {
		currentPage, _ = strconv.Atoi(queryPage)
	}

	g := models.NewGato()
	counter, _ := g.Count(tapp.App.DB.Primary)
	pagination := pagination.NewPagination(currentPage, itemsPerPage, counter)
	g, _ := g.FindAllPagination(tapp.App.DB.Primary, itemsPerPage, currentPage)

	return utils.Render(c, views.GatosShowList(tapp.GetTitleAndVersion(), *g, *pagination))
}

func ShowFormGato(c echo.Context, tapp *webcore.TangoApp, is_new bool) error {
	g := models.NewGato()

	if is_new {
		return utils.Render(c, views.GatosFormCreate(tapp.GetTitleAndVersion()))
	} else {
		id, _ := strconv.Atoi(c.Param("id"))
		g, _ := g.FindOne(tapp.App.DB.Primary, id)
		return utils.Render(c, views.GatosFormUpdate(tapp.GetTitleAndVersion(), g))
	}
}

func CreateGato(c echo.Context, tapp *webcore.TangoApp) error {
	// get the incoming values
	gDTO := models.GatoDTO{}
	if err := c.Bind(&gDTO); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	g := models.NewGato()
	g.Create(tapp.App.DB.Primary, gDTO.Name)

	return c.Redirect(http.StatusMovedPermanently, "/gatos/")
}

func UpdateGato(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// get the incoming values
	gDTO := models.GatoDTO{}
	if err := c.Bind(&gDTO); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	g := models.NewGato()
	g.Name = strings.ToLower(gDTO.Name)

	g.Update(tapp.App.DB.Primary, id, g.Name)

	return c.Redirect(http.StatusMovedPermanently, "/gatos/")
}

func DeleteGato(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))
	g := models.NewGato()
	g.Delete(tapp.App.DB.Primary, id)

	return c.Redirect(http.StatusMovedPermanently, "/gatos/")
}
	