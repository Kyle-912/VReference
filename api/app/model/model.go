package model

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Headset struct {
	gorm.Model
	Name string `gorm:"unique" json:"name"`
	Wireless bool `json:"wireless"`
	Price int `json:"price"`
	Description bool `json:"description"`
	Status bool `json:"status"`
}

func (e *Headset) Disable() {
	e.Status = false
}

func (p *Headset) Enable() {
	p.Status = true
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Headset{})
	return db
}