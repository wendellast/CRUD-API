package model

import (
	"fmt"

	"github.com/wendellast/CRUD-API/src/configuration/logger"
	"github.com/wendellast/CRUD-API/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomain) CreateUser(userDomain) *rest_err.RestErr {

	logger.Info("Init createUser model: ", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	fmt.Println(ud)

	return nil
}
