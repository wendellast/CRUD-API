package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/wendellast/CRUD-API/src/configuration/rest_err"
	"github.com/wendellast/CRUD-API/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var userResquest request.UserResquest

	if err := c.ShouldBindJSON(&userResquest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect filds, error=%s", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Print(userResquest)

}
