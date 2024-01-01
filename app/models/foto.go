
package models

import (
	"fmt"

	"github.com/k23dev/tango/pkg/tango_errors"
	"gorm.io/gorm"
)

type Foto struct {
	gorm.Model
	// add backticks in tags here
	Name string json:"name" param:"name" query:"name" form:"name"

}

type FotoDTO struct {
	// add backticks in tags here
	Name string json:"name" param:"name" query:"name" form:"name"
}

type FotoCounter struct {
	Total int
}

func NewFoto() *Foto {
	return &Foto{}
}

func (c *Foto) Count(db *gorm.DB) (int, error) {
	counter := &FotoCounter{}
	db.Model(&Foto{}).Select("count(ID) as total").Where("delete = ? ", "").Find(&counter)
	return counter.Total, nil
}

func (c *Foto) FindOne(db *gorm.DB, id int) (*Foto, error) {
	var foto Foto
	db.First(&foto, id)
	if foto.ID == 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Foto",
			Code:      0,
			Message:   tango_errors.MsgIDNotFound(id),
		}
	}
	return &foto, nil
}

func (c *Foto) FindAll(db *gorm.DB) ([]Foto, error) {
	var fotos []Foto
	db.Order("created_at ASC").Find(&fotos)
	if len(fotos) <= 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Foto",
			Code:      0,
			Message:   tango_errors.MsgZeroRecordsFound(),
		}
	}
	return fotos, nil
}

func (c *Foto) FindAllPagination(db *gorm.DB, itemsPerPage, currentPage int) (*[]Foto, error) {
	fotos := []Foto{}

	db.Order("created_at ASC").Limit(itemsPerPage).Offset(itemsPerPage * currentPage).Find(&fotos)
	if len(fotos) <= 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Foto",
			Code:      0,
			Message:   tango_errors.MsgZeroRecordsFound(),
		}
	}
	return &fotos, nil
}

func (c *Foto) Create(db *gorm.DB, name string) (*Foto, error) {
	foto := Foto{
		Name: name,
	}
	db.Create(&foto)
	return &foto, nil
}

func (c *Foto) Update(db *gorm.DB, id int, name string) (*Foto, error) {
	db.Model(&Foto{}).Where("ID =?", id).Update("name", name)
	return c, nil
}

func (c *Foto) Delete(db *gorm.DB, id int) (*Foto, error) {
	foto, err := c.FindOne(db, id)
	if err != nil {
		return nil, err
	}
	db.Delete(&foto)
	return foto, nil
}

func (c *Foto) GetIDAsString() string {
	return fmt.Sprintf("%d", c.ID)
}	
	