// Package user 存放用户 Model 的逻辑
package user

import "go-rxy/app/models"

// User 用户模型
type User struct {
	models.BaseModel

	Name     string `gorm:"column:name;not null;" json:"name,omitempty"`
	Email    string `gorm:"column:email;not null;" json:"-"`
	Phone    string `gorm:"column:phone;not null;" json:"-"`
	Password string `gorm:"column:password;not null;" json:"-"`

	models.CommonTimestampsField
}
