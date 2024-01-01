
package models

import (
	"fmt"

	"github.com/k23dev/tango/pkg/tango_errors"
	"gorm.io/gorm"
)

type Tanga struct {
	gorm.Model
	// add backticks in tags here
	Name string json:"name" param:"name" query:"name" form:"name"

}

type TangaDTO struct {
	// add backticks in tags here
	Name string json:"name" param:"name" query:"name" form:"name"
}

type TangaCounter struct {
	Total int
}

func NewTanga() *Tanga {
	return &Tanga{}
}

func (c *Tanga) Count(db *gorm.DB) (int, error) {
	counter := &TangaCounter{}
	db.Model(&Tanga{}).Select("count(ID) as total").Where("delete = ? ", "").Find(&counter)
	return counter.Total, nil
}

func (c *Tanga) FindOne(db *gorm.DB, id int) (*Tanga, error) {
	var tanga Tanga
	db.First(&tanga, id)
	if tanga.ID == 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Tanga",
			Code:      0,
			Message:   tango_errors.MsgIDNotFound(id),
		}
	}
	return &tanga, nil
}

func (c *Tanga) FindAll(db *gorm.DB) ([]Tanga, error) {
	var tangas []Tanga
	db.Order("created_at ASC").Find(&tangas)
	if len(tangas) <= 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Tanga",
			Code:      0,
			Message:   tango_errors.MsgZeroRecordsFound(),
		}
	}
	return tangas, nil
}

func (c *Tanga) FindAllPagination(db *gorm.DB, itemsPerPage, currentPage int) (*[]Tanga, error) {
	tangas := []Tanga{}

	db.Order("created_at ASC").Limit(itemsPerPage).Offset(itemsPerPage * currentPage).Find(&tangas)
	if len(tangas) <= 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "Tanga",
			Code:      0,
			Message:   tango_errors.MsgZeroRecordsFound(),
		}
	}
	return &tangas, nil
}

func (c *Tanga) Create(db *gorm.DB, name string) (*Tanga, error) {
	tanga := Tanga{
		Name: name,
	}
	db.Create(&tanga)
	return &tanga, nil
}

func (c *Tanga) Update(db *gorm.DB, id int, name string) (*Tanga, error) {
	db.Model(&Tanga{}).Where("ID =?", id).Update("name", name)
	return c, nil
}

func (c *Tanga) Delete(db *gorm.DB, id int) (*Tanga, error) {
	tanga, err := c.FindOne(db, id)
	if err != nil {
		return nil, err
	}
	db.Delete(&tanga)
	return tanga, nil
}

func (c *Tanga) GetIDAsString() string {
	return fmt.Sprintf("%d", c.ID)
}	
	