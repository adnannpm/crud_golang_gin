package controller

import (
	"net/http"
	"strconv"
	"tugas3/entity"
	"tugas3/gto"
	"tugas3/repository"

	"github.com/gin-gonic/gin"
)

type PublisherController struct{
	Repository repository.PublisherRepository
}

func NewPublisherController(repo repository.PublisherRepository) *PublisherController {
	return &PublisherController{
		Repository: repo,
	}
}

func (r *PublisherController) AddPublisher(c *gin.Context){
	var request gto.Publisher
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gto.Response{
			Success: false,
			Error: &gto.ErrorInfo{Code: "VALIDATION_ERROR", Message: err.Error()},
		})
		return
	}

	publisherData := entity.Publisher{
		Name: request.Name,
		Address: request.Address,
		Email: request.Email,
		Phone: request.Phone,
	}

	publisher, err := r.Repository.Add(publisherData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gto.Response{
			Success: false,
			Error: &gto.ErrorInfo{Code: "CREATE_PUBLISHER_ERROR", Message: err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, gto.Response{
		Success: true,
		Data: publisher,
	})
}

func (r *PublisherController) Update(c *gin.Context){
	var request gto.Publisher
	id := c.Param("id")

	idInt, _ := strconv.Atoi(id)

	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gto.Response{
			Success: false,
			Error: &gto.ErrorInfo{Code: "VALIDATION_ERROR", Message: err.Error()},
		})
		return
	}

	publisherData := entity.Publisher{
		ID: uint(idInt),
		Name: request.Name,
		Address: request.Address,
		Email: request.Email,
		Phone: request.Phone,
	}

	publisher, err := r.Repository.Update(publisherData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gto.Response{
			Success: false,
			Error: &gto.ErrorInfo{Code: "UPDATE_PUBLISHER_ERROR", Message: err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, gto.Response{
		Success: true,
		Data: publisher,
	})
}

func (r *PublisherController) Delete(c *gin.Context){
	id := c.Param("id")

	idInt, _ := strconv.Atoi(id)

	if err := r.Repository.Delete(int32(idInt)); err != nil {
		c.JSON(http.StatusBadRequest, gto.Response{
			Success: false,
			Error: &gto.ErrorInfo{
				Code: "PUBLISHER_NOT_FOUND",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gto.Response{
		Success: true,
		Data: "Delete Data Successfuly",
	}) 
}

func (r *PublisherController) GetData(c *gin.Context){
	id := c.Param("id")

	idInt, _ := strconv.Atoi(id)

	publisher, err := r.Repository.GetData(int32(idInt))
	if err != nil {
		c.JSON(http.StatusBadRequest, gto.Response{
			Success: false,
			Error: &gto.ErrorInfo{
				Code: "PUBLISHER_NOT_FOUND",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gto.Response{
		Success: true,
		Data: publisher,
	}) 
}

func (r *PublisherController) GetAllData(c *gin.Context){
	publishers, err := r.Repository.GetAllData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gto.Response{
			Success: false,
			Error: &gto.ErrorInfo{
				Code: "PUBLISHERS_NOT_FOUND",
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gto.Response{
		Success: true,
		Data: publishers,
	}) 
}