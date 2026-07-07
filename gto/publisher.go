package gto

type Publisher struct{
	Name	string		`form:"name" json:"name" xml:"name"  binding:"required"`
	Address string		`form:"address" json:"address" xml:"address"  binding:"required"`
	Phone	string		`form:"phone" json:"phone" xml:"phone"  binding:"required"`
	Email	string		`form:"email" json:"email" xml:"naemailme"  binding:"required"`
}