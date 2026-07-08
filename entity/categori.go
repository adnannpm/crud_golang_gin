package entity

type Categori struct{
	ID 		uint		`gorm:"primary_key;column:id"`
	Name 	string		`gorm:"column:name"`
}

func (u *Categori) TableName() string{
	return "categories"
}