package gto

type Categori struct{
	Name	string		`form:"name" json:"name" xml:"name"  binding:"required"`
}