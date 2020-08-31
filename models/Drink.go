package models

type Drink struct {
	Model
	Name  string `gorm:"size:255" json:"name,omitempty"`
	Price uint16 `json:"price,omitempty"`
	Stock uint16 `json:"stock,omitempty"`
}

func getAll(drinks *[]Drink) error {
	Db.Find(&drinks)
}
