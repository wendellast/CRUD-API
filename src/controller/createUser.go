package controller

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wendellast/CRUD-API/src/configuration/validation"
	"github.com/wendellast/CRUD-API/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var userResquest request.UserResquest

	if err := c.ShouldBindJSON(&userResquest); err != nil {
		log.Printf("Error trying to marshal object, error %s ", err.Error())
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)

		return
	}

	fmt.Print(userResquest)

}
