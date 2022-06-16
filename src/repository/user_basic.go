package repository

import (
	"api-merca/src/database"
	"api-merca/src/model"
)

type UserBasic struct {
}

func (ub UserBasic) FindByEmail(email string) (model.User, error) {
	var model model.User
	database.Connection.With("POSTGRES").First(&model, "email = ?", email)
	return model, nil
}
