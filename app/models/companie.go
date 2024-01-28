package models

import (
	"github.com/goravel/framework/database/orm"
)

type Companie struct {
	orm.Model
	Name     string
	Username string
	Password string
	CNPJ     string
	Email    string
}
