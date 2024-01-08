package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

// User 用户表结构
type User struct {
	ID        string     `json:"id" gorm:"primary_key;column:ID"`
	Username  string     `json:"username" gorm:"not null;column:Username"` // 用户名
	Password  string     `json:"password" gorm:"not null;column:Password"` // 密码
	Email     string     `json:"email" gorm:"column:Email"`                // 邮箱
	Phone     string     `json:"phone" gorm:"column:Phone"`                // 电话
	Role      int        `json:"role" gorm:"column:Role"`                  // 角色  0 表示普通用户，1 表示管理员
	Status    int        `json:"status" gorm:"column:Status"`              // 状态 0 表示正常，1 表示禁用`                                   // 状态 0 表示正常，1 表示禁用
	Address   string     `json:"address" gorm:"column:Address"`            // 地址
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`      // 软删除字段
}

// GetUser 获取用户信息 用于注册
func (u *User) GetUser(db *gorm.DB, username string) (*User, error) {
	var users []User
	result := db.LogMode(true).Where("Username = ?", username).Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &users[0], nil
}

// AddUser 添加一个用户
func (u *User) AddUser(db *gorm.DB, user *User) error {
	result := db.Create(user) // 通过数据的指针来创建
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return result.Error
	}
	return nil
}

// GetUsers 获取用户列表
func (u *User) GetUsers(db *gorm.DB) ([]User, error) {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// GetUserByID 获取单个用户信息
func (u *User) GetUserByID(db *gorm.DB, id string) (*User, error) {
	var user User
	result := db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// DeleteUser 删除单个用户，gorm的软删除
func (u *User) DeleteUser(db *gorm.DB, id string) error {
	result := db.Where("id = ?", id).Delete(&User{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// PutUser 修改单个用户信息
func (u *User) PutUser(db *gorm.DB, user *User) error {
	result := db.Model(&User{}).Where("id =?", user.ID).Updates(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetUserByUsername 通过用户名查询到用户ID
func GetUserByUsername(db *gorm.DB, username string) (string, error) {
	var user User
	var username1 = username
	fmt.Printf("%v\n", username1)
	result := db.Table("user").Where("Username = ?", username).First(&user)
	fmt.Printf("%v", user.ID)
	if result.Error != nil {
		return "", result.Error
	}
	return user.ID, nil
}

// GetUserByEmail 通过邮箱查找用户
func (u *User) GetUserByEmail(db *gorm.DB, email string) (User, error) {
	var user User
	var email1 = email
	fmt.Printf("%v\n", email1)
	result := db.Table("user").Where("Email =? and status =? ", email, 0).First(&user)
	fmt.Printf("%v", user.ID)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

// GetUserByEmailRoot 通过邮箱查找用户
func (u *User) GetUserByEmailRoot(db *gorm.DB, email string) (User, error) {
	var user User
	var email1 = email
	fmt.Printf("%v\n", email1)
	result := db.Table("user").Where("Email =? and role = ?", email, 1).First(&user)
	fmt.Printf("%v", user.ID)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
