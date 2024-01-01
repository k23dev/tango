
package models

import (
	"fmt"

	"github.com/k23dev/tango/pkg/tango_errors"
	"gorm.io/gorm"
)

type Gato struct {
	gorm.Model
	// add backticks in tags here
	Name string json:"name" param:"name" query:"name" form:"name"

}

type GatoDTO struct {
	// add backticks in tags here
	Name string json:"name" param:"name" query:"name" form:"name"
}

type GatoCounter struct {
	Total int
}

func NewGato() *Gato {
	return &Gato{}
}

func (c *Gato) Count(db *gorm.DB) (int, error) {
	counter := &GatoCounter{}
	db.Model(&Gato{}).Select("count(ID) as total").Where("delete = ? ", "").Find(&counter)
	return counter.Total, nil
}

func (c *Gato) FindOne(db *gorm.DB, id int) (*Gato, error) {
	var gato Gato
	db.First(&gato, id)
	if gato.ID == 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Gato",
			Code:      0,
			Message:   tango_errors.MsgIDNotFound(id),
		}
	}
	return &gato, nil
}

func (c *Gato) FindAll(db *gorm.DB) ([]Gato, error) {
	var gatos []Gato
	db.Order("created_at ASC").Find(&gatos)
	if len(gatos) <= 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Gato",
			Code:      0,
			Message:   tango_errors.MsgZeroRecordsFound(),
		}
	}
	return gatos, nil
}

func (c *Gato) FindAllPagination(db *gorm.DB, itemsPerPage, currentPage int) (*[]Gato, error) {
	gatos := []Gato{}

	db.Order("created_at ASC").Limit(itemsPerPage).Offset(itemsPerPage * currentPage).Find(&gatos)
	if len(gatos) <= 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Gato",
			Code:      0,
			Message:   tango_errors.MsgZeroRecordsFound(),
		}
	}
	return &gatos, nil
}

func (c *Gato) Create(db *gorm.DB, name string) (*Gato, error) {
	gato := Gato{
		Name: name,
	}
	db.Create(&gato)
	return &gato, nil
}

func (c *Gato) Update(db *gorm.DB, id int, name string) (*Gato, error) {
	db.Model(&Gato{}).Where("ID =?", id).Update("name", name)
	return c, nil
}

func (c *Gato) Delete(db *gorm.DB, id int) (*Gato, error) {
	gato, err := c.FindOne(db, id)
	if err != nil {
		return nil, err
	}
	db.Delete(&gato)
	return gato, nil
}

func (c *Gato) GetIDAsString() string {
	return fmt.Sprintf("%d", c.ID)
}	
	