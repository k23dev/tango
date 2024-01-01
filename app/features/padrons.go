
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


func FindOnePadron(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))

	p := models.NewPadron()
	padron, _ := p.FindOne(tapp.App.DB.Primary, id)
	if padron != nil {
		return utils.Render(c, views.PadronsShowOne(tapp.GetTitleAndVersion(), *padron))
	} else {
		return c.Redirect(http.StatusMovedPermanently, "/404")
	}
}

func FindAllPadrons(c echo.Context, tapp *webcore.TangoApp) error {
	queryPage := c.Param("page")
	var currentPage = 0
	if queryPage != "" {
		currentPage, _ = strconv.Atoi(queryPage)
	}

	p := models.NewPadron()
	counter, _ := p.Count(tapp.App.DB.Primary)
	pagination := pagination.NewPagination(currentPage, itemsPerPage, counter)
	p, _ := p.FindAllPagination(tapp.App.DB.Primary, itemsPerPage, currentPage)

	return utils.Render(c, views.PadronsShowList(tapp.GetTitleAndVersion(), *p, *pagination))
}

func ShowFormPadron(c echo.Context, tapp *webcore.TangoApp, is_new bool) error {
	p := models.NewPadron()

	if is_new {
		return utils.Render(c, views.PadronsFormCreate(tapp.GetTitleAndVersion()))
	} else {
		id, _ := strconv.Atoi(c.Param("id"))
		p, _ := p.FindOne(tapp.App.DB.Primary, id)
		return utils.Render(c, views.PadronsFormUpdate(tapp.GetTitleAndVersion(), p))
	}
}

func CreatePadron(c echo.Context, tapp *webcore.TangoApp) error {
	// get the incoming values
	pDTO := models.PadronDTO{}
	if err := c.Bind(&pDTO); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	p := models.NewPadron()
	p.Create(tapp.App.DB.Primary, pDTO.Name)

	return c.Redirect(http.StatusMovedPermanently, "/padrons/")
}

func UpdatePadron(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// get the incoming values
	pDTO := models.PadronDTO{}
	if err := c.Bind(&pDTO); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	p := models.NewPadron()
	p.Name = strings.ToLower(pDTO.Name)

	p.Update(tapp.App.DB.Primary, id, p.Name)

	return c.Redirect(http.StatusMovedPermanently, "/padrons/")
}

func DeletePadron(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))
	p := models.NewPadron()
	p.Delete(tapp.App.DB.Primary, id)

	return c.Redirect(http.StatusMovedPermanently, "/padrons/")
}
	