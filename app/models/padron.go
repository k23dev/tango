package models

import (
	"fmt"

	"github.com/k23dev/tango/pkg/tango_errors"
	"gorm.io/gorm"
)

type Padron struct {
	gorm.Model
	Name string `json:"name" param:"name" query:"name" form:"name"`
}

type PadronDTO struct {
	Name string `json:"name" param:"name" query:"name" form:"name"`
}

type PadronCounter struct {
	Total int
}

func NewPadron() *Padron {
	return &Padron{}
}

func (c *Padron) Count(db *gorm.DB) (int, error) {
	counter := &PadronCounter{}
	db.Model(&Padron{}).Select("count(ID) as total").Where("delete = ? ", "").Find(&counter)
	return counter.Total, nil
}

func (c *Padron) FindOne(db *gorm.DB, id int) (*Padron, error) {
	var padron Padron
	db.First(&padron, id)
	if padron.ID == 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Padron",
			Code:      0,
			Message:   tango_errors.MsgIDNotFound(id),
		}
	}
	return &padron, nil
}

func (c *Padron) FindAll(db *gorm.DB) ([]Padron, error) {
	var padrons []Padron
	db.Order("created_at ASC").Find(&padrons)
	if len(padrons) <= 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Padron",
			Code:      0,
			Message:   tango_errors.MsgZeroRecordsFound(),
		}
	}
	return padrons, nil
}

func (c *Padron) FindAllPagination(db *gorm.DB, itemsPerPage, currentPage int) (*[]Padron, error) {
	padrons := []Padron{}

	db.Order("created_at ASC").Limit(itemsPerPage).Offset(itemsPerPage * currentPage).Find(&padrons)
	if len(padrons) <= 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Padron",
			Code:      0,
			Message:   tango_errors.MsgZeroRecordsFound(),
		}
	}
	return &padrons, nil
}

func (c *Padron) Create(db *gorm.DB, name string) (*Padron, error) {
	padron := Padron{
		Name: name,
	}
	db.Create(&padron)
	return &padron, nil
}

func (c *Padron) Update(db *gorm.DB, id int, name string) (*Padron, error) {
	db.Model(&Padron{}).Where("ID =?", id).Update("name", name)
	return c, nil
}

func (c *Padron) Delete(db *gorm.DB, id int) (*Padron, error) {
	padron, err := c.FindOne(db, id)
	if err != nil {
		return nil, err
	}
	db.Delete(&padron)
	return padron, nil
}

func (c *Padron) GetIDAsString() string {
	return fmt.Sprintf("%d", c.ID)
}
