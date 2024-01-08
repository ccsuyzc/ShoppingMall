package model

import (
	"ShoppingMall/pkg/errcode"
	"github.com/jinzhu/gorm"
)

// Category 商品种类表
type Category struct {
	ID       string `json:"id" gorm:"primary_key;column:ID"`           // 主键
	Name     string `json:"name" gorm:"not null;column:Name"`          // 种类名称
	IsRemove bool   `json:"is_remove" gorm:"not null;column:IsRemove"` // 是否删除 false 表示未删除，true 表示删除
}

// GetCategory 获取商品种类列表
func (c *Category) GetCategory(db *gorm.DB) ([]Category, error) {
	var categories []Category
	err := db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// GetCategoryByID 根据ID获取商品种类
func (c *Category) GetCategoryByID(db *gorm.DB, id string) (*Category, error) {
	var category Category
	err := db.Where("id =?", id).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// PutCategory 修改商品种类
func (c *Category) PutCategory(db *gorm.DB, category *Category) error {
	if flag1, _ := c.CheckCategoryID(db, category.ID); flag1 { // 判断该商品种类是否存在
		if flag2 := CheckCommodityType(db, category.Name); flag2 { // 判断该商品种类下是否有商品
			err := db.Model(&category).Updates(category).Error
			if err != nil {
				return err
			}
		}

	}
	return errcode.NotFound
}

// PostCategory 添加商品种类
func (c *Category) PostCategory(db *gorm.DB, category *Category) error {
	if flag := c.CheckCategoryName(db, category.Name); flag {
		return errcode.AlreadyExists // 已存在
	}
	err := db.Create(&category).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteCategory 删除商品种类
func (c *Category) DeleteCategory(db *gorm.DB, id string) error {
	if flag1, category := c.CheckCategoryID(db, id); flag1 { // 判断该商品种类是否存在
		if flag2 := CheckCommodityType(db, category.ID); !flag2 { // 判断该商品种类下是否有商品
			err := db.Where("id =?", id).Delete(&Category{}).Error
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errcode.NotFound
}

// CheckCategoryName 确认该商品种类是否存在
func (c *Category) CheckCategoryName(db *gorm.DB, Name string) bool {
	var category Category
	err := db.Where("Name =?", Name).First(&category).Error
	if err != nil {
		return false
	}
	return true
}

// CheckCategoryID 确认该商品是否存在
func (c *Category) CheckCategoryID(db *gorm.DB, ID string) (bool, Category) {
	var category Category
	err := db.LogMode(true).Where("ID =?", ID).First(&category).Error
	if err != nil {
		return false, category
	}
	return true, category
}
