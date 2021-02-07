package dao

import (
	"Login/config"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func StartMySQL()(err error){
	DB,err = gorm.Open(config.GetConfigInf())
	return err
}

type DBmodel interface{
	Query() bool
}
