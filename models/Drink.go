package models

import "time"

type Drink struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Name      string     `gorm:"size:255" json:"name,omitempty"`
	Price     uint16     `json:"price,omitempty"`
	Stock     uint16     `json:"stock,omitempty"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Drinks []*Drink

// func (d *Drink) getAll(drinks *[]Drink) error {
// 	return Db.Find(&drinks)
// }
