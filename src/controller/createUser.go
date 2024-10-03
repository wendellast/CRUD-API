package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wendellast/CRUD-API/src/configuration/logger"
	"github.com/wendellast/CRUD-API/src/configuration/validation"
	"github.com/wendellast/CRUD-API/src/controller/model/request"
	"github.com/wendellast/CRUD-API/src/model"
	"github.com/wendellast/CRUD-API/src/model/service"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser Controller", zap.String("journey", "createuser"))
	var userResquest request.UserResquest

	if err := c.ShouldBindJSON(&userResquest); err != nil {
		logger.Error("Error trying to validate user info ", err)
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)

		return
	}

	domain := model.NewUserDomain(
		userResquest.Email,
		userResquest.Password,
		userResquest.Name,
		userResquest.Age,
	)
	service := service.NewUserDomainService()
	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	// response := response.UseResponse{
	// 	Id:    "test",
	// 	Email: userResquest.Email,
	// 	Name:  userResquest.Name,
	// 	Age:   userResquest.Age,
	// }

	logger.Info("User created successfully", zap.String("journey", "createuser"))

	c.String(http.StatusOK, "")
}
