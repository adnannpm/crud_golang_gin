package gto

type UserRegister struct {
	Fullname 	string		`form:"fullname" json:"fullname" xml:"fullname"  binding:"required"`
	Email 		string		`form:"email" json:"email" xml:"email"  binding:"required"`
	Password 	string		`form:"password" json:"password" xml:"password"  binding:"required"`
}

type UserLogin struct{
	Indentifier string		`form:"identifier" json:"identifier" binding:"required"`
	Password 	string		`form:"password" json:"password" xml:"password"  binding:"required"`
}