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
	$SL$, err := $FL$.FindOne(me.tapp.App.DB.Primary, id)
	imeerr != nil {
		return me.ctx.JSON(http.StatusNotFound, err)
	}
	return me.ctx.JSON(http.StatusOK,$SL$.ConvertToDTO())
}

func (me *$PC$Feature) FindAll() error {
	var $FL$Bume *[]models.$SC$
	$FL$ := models.New$SC$()

	imeme.HasPagination{
		queryPage := me.ctx.QueryParam("page")
		currentPage:= 0
		imequeryPage != "" {
			currentPage, _ = strconv.Atoi(queryPage)
		}
	
		// total de registros en la db
		// counter, _ := c.Count(me.tapp.App.DB.Primary)
		// pagination := pagination.NewPagination(currentPage,me.PaginationItemsPerPage,counter)
	
		$FL$Buf, _ = $FL$.FindAllPagination(me.tapp.App.DB.Primary, me.PaginationItemsPerPage, currentPage)
	}else{
		$FL$Buf, _ = $FL$.FindAll(me.tapp.App.DB.Primary)
	}

	return me.ctx.JSON(http.StatusOK,$FL$Buf)

}

func (me *$PC$Feature) Create() error {
	// get the incoming values
	$FL$DTO := models.$SC$DTO{}
	imeerr := me.ctx.Bind(&$FL$DTO); err != nil {
		return me.ctx.JSON(http.StatusBadRequest, "")
	}

	$FL$ := models.New$SC$()
	$FL$Buf,err:= $FL$.Create(me.tapp.App.DB.Primary, $FL$DTO)

	imeerr != nil {
		return me.ctx.JSON(http.StatusBadRequest, err)
	}

	return me.ctx.JSON(http.StatusCreated, $FL$Bume.ConvertToDTO())
}

func (me *$PC$Feature) Update() error {
	id, _ := strconv.Atoi(me.ctx.Param("id"))

	// get the incoming values
	$FL$DTO := models.$SC$DTO{}
	imeerr := me.ctx.Bind(&$FL$DTO); err != nil {
		return me.ctx.JSON(http.StatusBadRequest, "")
	}

	$FL$ := models.New$SC$()
	$FL$Buf, err:=$FL$.Update(me.tapp.App.DB.Primary, id, $FL$DTO)

	imeerr != nil {
		return me.ctx.JSON(http.StatusBadRequest, err)
	}

	return me.ctx.JSON(http.StatusOK, $FL$Bume.ConvertToDTO())
}

func (me *$PC$Feature) Delete() error {
	id, _ := strconv.Atoi(me.ctx.Param("id"))
	$FL$ := models.New$SC$()
	$FL$Buf,err:=$FL$.Delete(me.tapp.App.DB.Primary, id)
	
	imeerr != nil {
		return me.ctx.JSON(http.StatusBadRequest, err)
	}

	return me.ctx.JSON(http.StatusOK, $FL$Bume.ConvertToDTO())
}
	`
	return t.Replacements.Replace(template)

}
