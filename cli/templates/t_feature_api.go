package templates

func (t *Templates) FeatureAPI() string {

	t.setReplacements()

	template := `
package features

import (
	"net/http"
	"strconv"

	"github.com/k23dev/tango/app/models"
	"github.com/k23dev/tango/pkg/webcore"
	"github.com/labstack/echo/v4"
)

type $PC$Feature struct {
	ctx  echo.Context
	tapp *webcore.TangoApp
	db *gorm.DB
	HasPagination bool
	PaginationItemsPerPage int
}

func New$PC$Feature(tapp *webcore.TangoApp) *$PC$Feature {
	return &$PC$Feature{
		tapp: tapp,
		HasPagination:false,
		PaginationItemsPerPage:15,
		db:tapp.App.DB.Primary,
	}
}

func (me *$PC$Feature) SetCtx(ctx echo.Context) {
	me.ctx = ctx
}

func (me *$PC$Feature) SetDB(db *gorm.DB) {
	me.db = db
}

func (me *$PC$Feature) FindOne() error {
	id, _ := strconv.Atoi(me.ctx.Param("id"))

	$FL$ := models.New$SC$()
	$SL$, err := $FL$.FindOne(me.db, id)
	if err != nil {
		return me.ctx.JSON(http.StatusNotFound, err)
	}
	return me.ctx.JSON(http.StatusOK,$SL$.ConvertToDTO())
}

func (me *$PC$Feature) FindAll() error {
	var $FL$Buf *[]models.$SC$
	$FL$ := models.New$SC$()

	if me.HasPagination{
		queryPage := me.ctx.QueryParam("page")
		currentPage:= 0
		if queryPage != "" {
			currentPage, _ = strconv.Atoi(queryPage)
		}
	
		// total de registros en la db
		// counter, _ := c.Count(me.db)
		// pagination := pagination.NewPagination(currentPage,me.PaginationItemsPerPage,counter)
	
		$FL$Buf, _ = $FL$.FindAllPagination(me.db, me.PaginationItemsPerPage, currentPage)
	}else{
		$FL$Buf, _ = $FL$.FindAll(me.db)
	}

	return me.ctx.JSON(http.StatusOK,$FL$Buf)

}

func (me *$PC$Feature) Create() error {
	// get the incoming values
	$FL$DTO := models.$SC$DTO{}
	if err := me.ctx.Bind(&$FL$DTO); err != nil {
		return me.ctx.JSON(http.StatusBadRequest, "")
	}

	$FL$ := models.New$SC$()
	$FL$Buf,err:= $FL$.Create(me.db, $FL$DTO)

	if err != nil {
		return me.ctx.JSON(http.StatusBadRequest, err)
	}

	return me.ctx.JSON(http.StatusCreated, $FL$Buf.ConvertToDTO())
}

func (me *$PC$Feature) Update() error {
	id, _ := strconv.Atoi(me.ctx.Param("id"))

	// get the incoming values
	$FL$DTO := models.$SC$DTO{}
	if err := me.ctx.Bind(&$FL$DTO); err != nil {
		return me.ctx.JSON(http.StatusBadRequest, "")
	}

	$FL$ := models.New$SC$()
	$FL$Buf, err:=$FL$.Update(me.db, id, $FL$DTO)

	if err != nil {
		return me.ctx.JSON(http.StatusBadRequest, err)
	}

	return me.ctx.JSON(http.StatusOK, $FL$Buf.ConvertToDTO())
}

func (me *$PC$Feature) Delete() error {
	id, _ := strconv.Atoi(me.ctx.Param("id"))
	$FL$ := models.New$SC$()
	$FL$Buf,err:=$FL$.Delete(me.db, id)
	
	if err != nil {
		return me.ctx.JSON(http.StatusBadRequest, err)
	}

	return me.ctx.JSON(http.StatusOK, $FL$Buf.ConvertToDTO())
}
	`
	return t.Replacements.Replace(template)

}
