package entity

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `json:"product_name" binding: "required"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}
