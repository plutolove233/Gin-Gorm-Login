package models

import (
	"Login/dao"
)

type UserInf struct{
	Id int
	UserName string
	Password string
	Email string
	Phone string
}

type UserInfmodel interface {
	dao.DBmodel
}

func Query(N string,P string,E string,Ph string)(ok bool){
	var user UserInf
	dao.DB.Where("user_name=?",N).Find(&user)
	if user.Password==P&&user.Email==E&&user.Phone==Ph{
		ok = true
	}else{
		ok = false
	}
	return
}