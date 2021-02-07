package main

import (
	"Login/dao"
	"Login/models"
	"Login/router"
	"fmt"
)

func main() {
	fmt.Println("This is our first project!")
	err := dao.StartMySQL()
	if err!=nil{
		fmt.Println("数据库启动失败,%v",err)
	}
	defer dao.DB.Close()
	dao.DB.AutoMigrate(&models.UserInf{})
	r   := router.StartRouter()
	if err=r.Run(":8080");err!=nil{
		fmt.Println("运行时出错",err)
	}
}
