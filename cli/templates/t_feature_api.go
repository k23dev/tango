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

func (f *$PC$Feature) SetCtx(ctx echo.Context) {
	f.ctx = ctx
}

func (f *$PC$Feature) SetDB(db *gorm.DB) {
	f.db = db
}

func (f *$PC$Feature) FindOne() error {
	id, _ := strconv.Atoi(f.ctx.Param("id"))

	$FL$ := models.New$SC$()
	$SL$, err := $FL$.FindOne(f.tapp.App.DB.Primary, id)
	if err != nil {
		return f.ctx.JSON(http.StatusNotFound, err)
	}
	return f.ctx.JSON(http.StatusOK,$SL$.ConvertToDTO())
}

func (f *$PC$Feature) FindAll() error {
	var $FL$Buf *[]models.$SC$
	$FL$ := models.New$SC$()

	if f.HasPagination{
		queryPage := f.ctx.Param("page")
		currentPage:= 0
		if queryPage != "" {
			currentPage, _ = strconv.Atoi(queryPage)
		}
	
		// total de registros en la db
		// counter, _ := c.Count(f.tapp.App.DB.Primary)
		// pagination := pagination.NewPagination(currentPage,f.PaginationItemsPerPage,counter)
	
		$FL$Buf, _ = $FL$.FindAllPagination(f.tapp.App.DB.Primary, f.PaginationItemsPerPage, currentPage)
	}else{
		$FL$Buf, _ = $FL$.FindAll(f.tapp.App.DB.Primary)
	}

	return f.ctx.JSON(http.StatusOK,$FL$Buf)

}

func (f *$PC$Feature) Create() error {
	// get the incoming values
	$FL$DTO := models.$SC$DTO{}
	if err := f.ctx.Bind(&$FL$DTO); err != nil {
		return f.ctx.JSON(http.StatusBadRequest, "")
	}

	$FL$ := models.New$SC$()
	$FL$Buf,err:= $FL$.Create(f.tapp.App.DB.Primary, $FL$DTO)

	if err != nil {
		return f.ctx.JSON(http.StatusBadRequest, err)
	}

	return f.ctx.JSON(http.StatusCreated, $FL$Buf.ConvertToDTO())
}

func (f *$PC$Feature) Update() error {
	id, _ := strconv.Atoi(f.ctx.Param("id"))

	// get the incoming values
	$FL$DTO := models.$SC$DTO{}
	if err := f.ctx.Bind(&$FL$DTO); err != nil {
		return f.ctx.JSON(http.StatusBadRequest, "")
	}

	$FL$ := models.New$SC$()
	$FL$Buf, err:=$FL$.Update(f.tapp.App.DB.Primary, id, $FL$DTO)

	if err != nil {
		return f.ctx.JSON(http.StatusBadRequest, err)
	}

	return f.ctx.JSON(http.StatusOK, $FL$Buf.ConvertToDTO())
}

func (f *$PC$Feature) Delete() error {
	id, _ := strconv.Atoi(f.ctx.Param("id"))
	$FL$ := models.New$SC$()
	$FL$Buf,err:=$FL$.Delete(f.tapp.App.DB.Primary, id)
	
	if err != nil {
		return f.ctx.JSON(http.StatusBadRequest, err)
	}

	return f.ctx.JSON(http.StatusOK, $FL$Buf.ConvertToDTO())
}
	`
	return t.Replacements.Replace(template)

}
