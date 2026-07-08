package controller

import (
	"net/http"
	"strconv"
	"tugas3/entity"
	"tugas3/gto"
	"tugas3/repository"

	"github.com/gin-gonic/gin"
)

type CategoriController struct{
	repository repository.CategoryRepository
}

func NewCategoriController(repo repository.CategoryRepository) *CategoriController{
	return &CategoriController{
		repository: repo,
	}
}

func (r *CategoriController) Add(c *gin.Context){
	var request gto.Categori
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gto.Response{
			Success: false,
			Error: &gto.ErrorInfo{Code: "VALIDATION_ERROR", Message: err.Error()},
		})
		return
	}

	data := entity.Categori{
		Name: request.Name,
	}

	categori, err := r.repository.Add(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gto.Response{
			Success: false,
			Error: &gto.ErrorInfo{Code: "CREATE_CATEGORI_ERROR", Message: err.Error()},
		})
		return
	}
	c.JSON(http.StatusOK, gto.Response{
		Success: true,
		Data: categori,
	})
}

func (r *CategoriController) Update(c *gin.Context){
	var request gto.Categori
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gto.Response{
			Success: false,
			Error: &gto.ErrorInfo{Code: "VALIDATION_ERROR", Message: err.Error()},
		})
		return
	}

	data := entity.Categori{
		ID: uint(idInt),
		Name: request.Name,
	}

	categori, err := r.repository.Update(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gto.Response{
			Success: false,
			Error: &gto.ErrorInfo{Code: "UPDATE_CATEGORI_ERROR", Message: err.Error()},
		})
		return
	}
	c.JSON(http.StatusOK, gto.Response{
		Success: true,
		Data: categori,
	})
}

func (r *CategoriController) Delete(c *gin.Context){
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	err := r.repository.Delete(int32(idInt))
	if err != nil {
		c.JSON(http.StatusBadRequest, gto.Response{
			Success: false,
			Error: &gto.ErrorInfo{Code: "DELETE_CATEGORI_ERROR", Message: err.Error()},
		})
		return
	}
	c.JSON(http.StatusOK, gto.Response{
		Success: true,
		Data: "Delete Data Susscesfuly",
	})
}

func (r *CategoriController) GetAllData(c *gin.Context){
	categories, err := r.repository.GetAllData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gto.Response{
			Success: false,
			Error: &gto.ErrorInfo{Code: "CATEGORIES_NOT_FOUND", Message: err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, gto.Response{
		Success: true,
		Data: categories,
	})
}