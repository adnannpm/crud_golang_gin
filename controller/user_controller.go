package controller

import (
	"net/http"
	"tugas3/entity"
	"tugas3/gto"
	"tugas3/helper"
	"tugas3/repository"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct{
	Repository repository.UserRepository
}

func NewUserController(repo repository.UserRepository) *UserController {
	return &UserController{
		Repository: repo,
	}
}

func (r *UserController) Register(c *gin.Context){
	var user gto.UserRegister
	if err := c.ShouldBindBodyWithJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest, gto.Response{
			Success: false,
			Error: &gto.ErrorInfo{Code: "VALIDATION_ERROR", Message: err.Error()},
		})
		return
	}

	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gto.Response{
			Success: false,
			Error: &gto.ErrorInfo{Code: "USER_HASH_PASSWORD_ERROR", Message: err.Error()},
		})
		return
	}

	userQuery := entity.User{
		Fullname: user.Fullname,
		Email: user.Email,
		Password: hashedPassword,
	} 

	result, err := r.Repository.Register(userQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gto.Response{
			Success: false,
			Error: &gto.ErrorInfo{Code: "REGISTER_ERROR", Message: err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, gto.Response{
		Success: true,
		Data: result,
	})
}

func (r *UserController) Login(c *gin.Context) {
	var user gto.UserLogin
	if err := c.ShouldBindBodyWithJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest, gto.Response{
			Success: false,
			Error: &gto.ErrorInfo{Code: "VALIDATION_ERROR", Message: err.Error()},
		})
		return
	}

	result, err := r.Repository.Login(user.Indentifier)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gto.Response{
			Success: false,
			Error: &gto.ErrorInfo{Code: "UNAUTHORIZED_ERROR", Message: err.Error()},
		})
		return
	}

	err = bcrypt.CompareHashAndPassword(
        []byte(result.Password),
        []byte(user.Password),
    )

	if err != nil {
		c.JSON(http.StatusUnauthorized, gto.Response{
			Success: false,
			Error: &gto.ErrorInfo{Code: "UNAUTHORIZED_ERROR", Message: err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, gto.Response{Success: true, Data: "awfjaifawf"})
}

