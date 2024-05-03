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

type $PC$Feature struct {
	ctx  echo.Context
	tapp *webcore.TangoApp
	db *gorm.DB
}

func New$PC$Feature(ctx echo.Context, tapp *webcore.TangoApp) *$PC$Feature {
	return &$PC$Feature{
		ctx:  ctx,
		tapp: tapp,
	}
}

func (f *$PC$Feature) SetCtx(ctx echo.Context) {
	f.ctx = ctx
}

func (f *$PC$Feature) SetDB(db *gorm.DB) {
	f.db = db
}

func (f *$PC$Feature) FindOne(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(f.ctx.Param("id"))

	$FL$ := models.New$SC$()
	$SL$, _ := $FL$.FindOne(f.tapp.App.DB.Primary, id)
	if $SL$ != nil {
		return utils.Render(ctx, views.$PC$ShowOne(f.tapp.GetTitleAndVersion(), *$SL$))
	} else {
		return f.ctx.Redirect(http.StatusMovedPermanently, "/404")
	}
}

func (f *$PC$Feature) FindAll(ctx echo.Context, tapp *webcore.TangoApp) error {
	queryPage := f.ctx.Param("page")
	var currentPage = 0
	if queryPage != "" {
		currentPage, _ = strconv.Atoi(queryPage)
	}

	$FL$ := models.New$SC$()
	counter, _ := $FL$.Count(f.tapp.App.DB.Primary)
	pagination := pagination.NewPagination(currentPage, itemsPerPage, counter)
	$FL$Buf, _ := $FL$.FindAllPagination(f.tapp.App.DB.Primary, itemsPerPage, currentPage)

	if $FL$Buf != nil {
		return utils.Render(ctx, views.$PC$ShowList(f.tapp.GetTitleAndVersion(), *$FL$Buf, *pagination))
	}else{
		return utils.Render(ctx, views.$PC$ShowListEmpty(f.tapp.GetTitleAndVersion()))
	}

}

func (f *$PC$Feature) ShowForm(ctx echo.Context, tapp *webcore.TangoApp, is_new bool) error {
	$FL$ := models.New$SC$()

	if is_new {
		return utils.Render(ctx, views.$PC$FormCreate(f.tapp.GetTitleAndVersion()))
	} else {
		id, _ := strconv.Atoi(f.ctx.Param("id"))
		$FL$, _ := $FL$.FindOne(f.tapp.App.DB.Primary, id)
		return utils.Render(ctx, views.$PC$FormUpdate(f.tapp.GetTitleAndVersion(), $FL$))
	}
}

func (f *$PC$Feature) Create(ctx echo.Context, tapp *webcore.TangoApp) error {
	// get the incoming values
	$FL$DTO := models.$SC$DTO{}
	if err := f.ctx.Bind(&$FL$DTO); err != nil {
		return f.ctx.String(http.StatusBadRequest, "Bad request")
	}

	$FL$ := models.New$SC$()
	$FL$.Create(f.tapp.App.DB.Primary, $FL$DTO)

	return f.ctx.Redirect(http.StatusMovedPermanently, "/$PL$/")
}

func (f *$PC$Feature) Update(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(f.ctx.Param("id"))

	// get the incoming values
	$FL$DTO := models.$SC$DTO{}
	if err := f.ctx.Bind(&$FL$DTO); err != nil {
		return f.ctx.String(http.StatusBadRequest, "Bad request")
	}

	$FL$ := models.New$SC$()
	$FL$.Update(f.tapp.App.DB.Primary, id, $FL$DTO)

	return f.ctx.Redirect(http.StatusMovedPermanently, "/$PL$/")
}

func (f *$PC$Feature) Delete(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(f.ctx.Param("id"))
	$FL$ := models.New$SC$()
	$FL$.Delete(f.tapp.App.DB.Primary, id)

	return f.ctx.Redirect(http.StatusMovedPermanently, "/$PL$/")
}
	`
	return t.Replacements.Replace(template)

}
