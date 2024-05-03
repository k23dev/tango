package templates

func (t *Templates) Models() string {

	t.setReplacements()

	var template string

	template = `
package models

import (
	"fmt"

	"github.com/k23dev/tango/pkg/tango_errors"
	"gorm.io/gorm"
)

type $SC$ struct {
	gorm.Model
	Name string
}
type $SC$DTO struct {
`
	template += "	ID uint `json:\"id\" param:\"id\" query:\"id\" form:\"id\"` \n"
	template += "	Name string `json:\"name\" param:\"name\" query:\"name\" form:\"name\"`"
	template += `
}

type $SC$Counter struct {
	Total int
}

func New$SC$() *$SC$ {
	return &$SC${}
}

func ($FL$ *$SC$) ConvertToDTO() *$SC$DTO {
	return &$SC$DTO{
		ID:   $FL$.ID,
		Name: $FL$.Name,
	}
}

func ($FL$ *$SC$) ConvertFromDTO(dto *$SC$DTO) *$SC$ {
	return &$SC${
		ID:   $FL$.ID,
		Name: $FL$.Name,
	}
}

func ($FL$ *$SC$) Count(db *gorm.DB) (int, error) {
	counter := &$SC$Counter{}
	db.Model(&$SC${}).Select("count(ID) as total").Find(&counter)
	return counter.Total, nil
}

func ($FL$ *$SC$) FindOne(db *gorm.DB, id int) (*$SC$, error) {
	var $SL$ $SC$
	db.First(&$SL$, id)
	if $SL$.ID == 0 {
		return nil, tango_errors.ReturnModel("$SC$", tango_errors.MsgNotFound(), 0)
	}
	return &$SL$, nil
}

func ($FL$ *$SC$) FindAll(db *gorm.DB) (*[]$SC$, error) {
	var $PL$ []$SC$
	db.Order("created_at ASC").Find(&$PL$)
	if len($PL$) <= 0 {
		return nil, tango_errors.ReturnModel("$SC$", tango_errors.MsgZeroRecordsFound(), 0)
	}
	return &$PL$, nil
}

func ($FL$ *$SC$) FindAllPagination(db *gorm.DB, itemsPerPage, currentPage int) (*[]$SC$, error) {
	$PL$ := []$SC${}

	db.Order("created_at ASC").Limit(itemsPerPage).Offset(itemsPerPage * currentPage).Find(&$PL$)
	if len($PL$) <= 0 {
		return nil, tango_errors.ReturnModel("$SC$", tango_errors.MsgZeroRecordsFound(), 0)
	}
	return &$PL$, nil
}

func ($FL$ *$SC$) Create(db *gorm.DB, dto $SC$DTO) (*$SC$, error) {
	$FL$.satinizeDTOCreate(&dto)
	$SL$ := $SC${
		Name: dto.Name,
	}
	result := db.Create(&$SL$)
	if result.Error != nil {
		return &$SC${}, result.Error
	}
	return &$SL$, nil
}

func ($FL$ *$SC$) Update(db *gorm.DB, id int, dto $SC$DTO) (*$SC$, error) {
	$FL$.satinizeDTOUpdate(&dto)

	$SL$ := &$SC${}
	db.First($SL$, "id=?", id)
	if $SL$.ID == 0 {
		return $SL$, tango_errors.ReturnModel("$SC$", tango_errors.MsgIDNotFound(id), 0)
	}

	// changes
	$SL$.Name = dto.Name
	
	db.Save($SL$)
	return $SL$, nil
}

func ($FL$ *$SC$) Delete(db *gorm.DB, id int) (*$SC$, error) {
	$SL$, err := $FL$.FindOne(db, id)
	if err != nil {
		return nil, err
	}
	db.Delete(&$SL$)
	return $SL$, nil
}

func ($FL$ *$SC$) satinizeDTOCreate(dto *$SC$DTO) error {
	// TODO
	dto.Name = strings.TrimSpace(dto.Name)
	return nil
}

func ($FL$ *$SC$) satinizeDTOUpdate(dto *$SC$DTO) error {
	// TODO
	dto.Name = strings.TrimSpace(dto.Name)
	return nil
}
	`
	return t.Replacements.Replace(template)

}
