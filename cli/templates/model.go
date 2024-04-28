package templates

func (t *Templates) Model() string {

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
	template += "	Name string `json:\"name\" param:\"name\" query:\"name\" form:\"name\"`"
	template += `
}

type $SC$Counter struct {
	Total int
}

func New$SC$() *$SC$ {
	return &$SC${}
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
		return nil, &tango_errors.ModelError{
			ModelName: "$SC$",
			Code:      0,
			Message:   tango_errors.MsgIDNotFound(id),
		}
	}
	return &$SL$, nil
}

func ($FL$ *$SC$) FindAll(db *gorm.DB) ([]$SC$, error) {
	var $PL$ []$SC$
	db.Order("created_at ASC").Find(&$PL$)
	if len($PL$) <= 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "$SC$",
			Code:      0,
			Message:   tango_errors.MsgZeroRecordsFound(),
		}
	}
	return $PL$, nil
}

func ($FL$ *$SC$) FindAllPagination(db *gorm.DB, itemsPerPage, currentPage int) (*[]$SC$, error) {
	$PL$ := []$SC${}

	db.Order("created_at ASC").Limit(itemsPerPage).Offset(itemsPerPage * currentPage).Find(&$PL$)
	if len($PL$) <= 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "$SC$",
			Code:      0,
			Message:   tango_errors.MsgZeroRecordsFound(),
		}
	}
	return &$PL$, nil
}

func ($FL$ *$SC$) Create(db *gorm.DB, dto $SC$DTO) (*$SC$, error) {
	$FL$.SatinizeDTOCreate(&dto)
	$SL$ := $SC${
		Name: dto.Name,
	}
	db.Create(&$SL$)
	return &$SL$, nil
}

func ($FL$ *$SC$) Update(db *gorm.DB, id int, dto $SC$DTO) (*$SC$, error) {
	$FL$.SatinizeDTOUpdate(&dto)
	db.Model(&$SC${}).Where("ID =?", id).Update("name", dto.Name)
	return $FL$, nil
}

func ($FL$ *$SC$) Delete(db *gorm.DB, id int) (*$SC$, error) {
	$SL$, err := $FL$.FindOne(db, id)
	if err != nil {
		return nil, err
	}
	db.Delete(&$SL$)
	return $SL$, nil
}

func ($FL$ *$SC$) GetIDAsString() string {
	return fmt.Sprintf("%d", $FL$.ID)
}	

func ($FL$ *$SC$) SatinizeDTOCreate(dto *$SC$DTO) error {
	// TODO
	return nil
}

func ($FL$ *$SC$) SatinizeDTOUpdate(dto *$SC$DTO) error {
	// TODO
	return nil
}
	`
	return t.Replacements.Replace(template)

}
