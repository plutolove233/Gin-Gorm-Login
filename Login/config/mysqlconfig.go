package config

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

type Configinf struct{
	Mysql MySQLconfig
}

var (
	config *Configinf
	once sync.Once
)

type MySQLconfig struct{
	Host string
	Ports int
	Username string
	Password string
	Database string
	Charset string
}

func (c *Configinf)InitConfig(){
	viper.SetConfigFile("config/config.toml")
	if err := viper.ReadInConfig();err!=nil{
		fmt.Println("导入配置文件失败，错误",err)
	}
	if err:=viper.Unmarshal(c);err!=nil{
		fmt.Println("配置文件绑定失败，错误",err)
	}
}

func GetConfigInf()(string,string){
	once.Do(func(){
		config = &Configinf{}
		config.InitConfig()
	})
	return "mysql",fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",config.Mysql.Username,
		config.Mysql.Password, config.Mysql.Host, config.Mysql.Ports, config.Mysql.Database)
}