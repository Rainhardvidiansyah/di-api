package entity

type User struct {
	Id uint64 `gorm:"primary_key; auto_increment" json:"id"`
}

type Article struct {
	Id     uint64 `gorm:"primary_key; auto_increment" json:"id"`
	person User   `json:"person"`
}
