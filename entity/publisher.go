package entity

type Publisher struct{
	ID uint	`gorm:"primary_key;column:id"`
	Name string	`gorm:"column:name"`
	Address string	`gorm:"column:address"`
	Phone string	`gorm:"column:phone"`
	Email string	`gorm:"column:email"`
}