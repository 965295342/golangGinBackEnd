package db

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DRIVER = "mysql"

var db *gorm.DB

func GetDb() *gorm.DB {
	return db
}

type conf struct {
	Url      string `yaml:"url"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	Port     string `yaml:"post"`
}

func InitMySql() error {
	var c conf
	var err error
	//获取yaml配置参数
	conf := c.getConf()
	//将yaml配置参数接成连接数据库的url
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.UserName,
		conf.Password,
		conf.Url,
		conf.Port,
		conf.DbName,
	)
	//连接数据库
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	// db, err = db.DB()
	// if err != nil {
	// 	return err
	// }
	// err = db.Ping()
	// return err
	return nil
}

func Close() {
	closeDb, _ := db.DB()
	closeDb.Close()
}

func (c *conf) getConf() *conf {
	//读取resources/application.yaml文件
	yamlFile, err := os.ReadFile("./resources/application.yaml")
	//若出现错误，打印错误提示
	if err != nil {
		fmt.Println(err.Error())
	}
	//将读取的字符串转换成结构体conf
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
