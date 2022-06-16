package model

import "gorm.io/gorm"

type Tenant struct {
	gorm.Model
	Entidade string `json:"entidade"`
	Database string `json:"database"`
}
