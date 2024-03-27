package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name       string `gorm:"column:name"`
	NickName   string `gorm:"column:nickName"`
	Avatar     string `gorm:"column:avatar"`
	Password   string `gorm:"column:password"`
	Email      string `gorm:"column:email"`
	Mobile     string `gorm:"column:mobile"`
	DelStatus  int    `gorm:"column:delStatus"`
	CreateTime int64  `gorm:"column:createTime"`
}

func (User) TableName() string {
	return "users" // 指定表名为 "users"
}
