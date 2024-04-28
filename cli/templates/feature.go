package templates

func (t *Templates) Feature() string {

	t.setReplacements()

	template := `
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


func FindOne$SC$(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	$FL$ := models.New$SC$()
	$SL$, _ := $FL$.FindOne(tapp.App.DB.Primary, id)
	if $SL$ != nil {
		return utils.Render(ctx, views.$PC$ShowOne(tapp.GetTitleAndVersion(), *$SL$))
	} else {
		return ctx.Redirect(http.StatusMovedPermanently, "/404")
	}
}

func FindAll$PC$(ctx echo.Context, tapp *webcore.TangoApp) error {
	queryPage := ctx.Param("page")
	var currentPage = 0
	if queryPage != "" {
		currentPage, _ = strconv.Atoi(queryPage)
	}

	$FL$ := models.New$SC$()
	counter, _ := $FL$.Count(tapp.App.DB.Primary)
	pagination := pagination.NewPagination(currentPage, itemsPerPage, counter)
	$FL$Buf, _ := $FL$.FindAllPagination(tapp.App.DB.Primary, itemsPerPage, currentPage)

	if $FL$Buf != nil {
		return utils.Render(ctx, views.$PC$ShowList(tapp.GetTitleAndVersion(), *$FL$Buf, *pagination))
	}else{
		return utils.Render(ctx, views.$PC$ShowListEmpty(tapp.GetTitleAndVersion()))
	}

}

func ShowForm$SC$(ctx echo.Context, tapp *webcore.TangoApp, is_new bool) error {
	$FL$ := models.New$SC$()

	if is_new {
		return utils.Render(ctx, views.$PC$FormCreate(tapp.GetTitleAndVersion()))
	} else {
		id, _ := strconv.Atoi(ctx.Param("id"))
		$FL$, _ := $FL$.FindOne(tapp.App.DB.Primary, id)
		return utils.Render(ctx, views.$PC$FormUpdate(tapp.GetTitleAndVersion(), $FL$))
	}
}

func Create$SC$(ctx echo.Context, tapp *webcore.TangoApp) error {
	// get the incoming values
	$FL$DTO := models.$SC$DTO{}
	if err := ctx.Bind(&$FL$DTO); err != nil {
		return ctx.String(http.StatusBadRequest, "Bad request")
	}

	$FL$ := models.New$SC$()
	$FL$.Create(tapp.App.DB.Primary, $FL$DTO)

	return ctx.Redirect(http.StatusMovedPermanently, "/$PL$/")
}

func Update$SC$(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	// get the incoming values
	$FL$DTO := models.$SC$DTO{}
	if err := ctx.Bind(&$FL$DTO); err != nil {
		return ctx.String(http.StatusBadRequest, "Bad request")
	}

	$FL$ := models.New$SC$()
	$FL$.Update(tapp.App.DB.Primary, id, $FL$DTO)

	return ctx.Redirect(http.StatusMovedPermanently, "/$PL$/")
}

func Delete$SC$(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	$FL$ := models.New$SC$()
	$FL$.Delete(tapp.App.DB.Primary, id)

	return ctx.Redirect(http.StatusMovedPermanently, "/$PL$/")
}
	`
	return t.Replacements.Replace(template)

}
