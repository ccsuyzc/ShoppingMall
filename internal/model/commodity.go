package model

import (
	"ShoppingMall/pkg/errcode"
	"fmt"
	"github.com/jinzhu/gorm"
)

// Commodity 商品信息表结构
type Commodity struct {
	ID                  string  `json:"id" gorm:"primary_key;column:ID"`                                   // 主键
	Name                string  `json:"name" gorm:"not null;column:Name"`                                  // 商品名称
	Price               float64 `json:"price" gorm:"not null;column:Price"`                                // 商品价格
	Stock               int     `json:"stock" gorm:"not null;column:Stock"`                                // 商品库存
	CategoryID          string  `json:"category_id" gorm:"not null;column:CategoryID"`                     // 商品分类
	Status              int     `json:"status" gorm:"not null;column:Status"`                              // 状态 0 表示正常，1 表示禁用
	Image1              string  `json:"image1" gorm:"column:Image1"`                                       // 商品图片1
	Image2              string  `json:"image2" gorm:"column:Image2"`                                       // 商品图片2`                                                            // 商品图片2
	Image3              string  `json:"image3" gorm:"column:Image3"`                                       // 商品图片3
	Desc                string  `json:"desc" gorm:"column:Desc"`                                           // 商品描述
	HomePageOnRecommend int     `json:"home_page_on_recommend" gorm:"not null;column:HomePageOnRecommend"` // 首页是否推荐 0 表示不推荐，1 表示推荐
	OnTerRecommend      int     `json:"on_ter_recommend" gorm:"not null;column:OnTerRecommend"`            // 此字段待定 0,1待定
	IsRemove            bool    `json:"is_remove" gorm:"not null;column:IsRemove"`                         // 是否删除 false 表示未删除，true 表示删除
}

// GetCommodity 得到全部商品信息,其中IsRemove字段为false的商品信息
func (c *Commodity) GetCommodity(db *gorm.DB) ([]Commodity, error) {
	var commodities []Commodity
	if err := db.Where("IsRemove =?", false).Find(&commodities).Error; err != nil {
		return nil, err
	}
	return commodities, nil
}

// GetCommodityByID 根据商品ID得到商品信息
func (c *Commodity) GetCommodityByID(db *gorm.DB, id string) (*Commodity, error) {
	var commodity Commodity
	err := db.Where("id =?", id).First(&commodity).Error
	if err != nil {
		return nil, err
	}
	return &commodity, nil
}

// PutCommodity 修改商品信息
func (c *Commodity) PutCommodity(db *gorm.DB, commodity *Commodity) error {
	fmt.Printf("this:%+v\n", commodity)
	if flag := c.CheckCommodity(db, commodity.ID); flag {
		if err := db.LogMode(true).Table("commodity").Where("id =?", commodity.ID).Updates(map[string]interface{}{
			"Name":                commodity.Name,
			"Price":               commodity.Price,
			"Stock":               commodity.Stock,
			"CategoryID":          commodity.CategoryID,
			"Status":              commodity.Status,
			"Image1":              commodity.Image1,
			"Image2":              commodity.Image2,
			"Image3":              commodity.Image3,
			"Desc":                commodity.Desc,
			"HomePageOnRecommend": commodity.HomePageOnRecommend,
			"OnTerRecommend":      commodity.OnTerRecommend,
			"IsRemove":            commodity.IsRemove,
		}).Error; err != nil {
			return err
		}
		return nil
	}
	// 返回一个不存在该商品的错误
	return errcode.NotFound
}

// PutCommodityImage 修改商品的图片信息
func (c *Commodity) PutCommodityImage(db *gorm.DB, id, image string, val int) error {
	if flag := c.CheckCommodity(db, id); flag {
		if val == 1 {
			if err := db.Model(&Commodity{}).Where("id =?", id).Updates(map[string]interface{}{"Image1": image}).Error; err != nil {
				return err
			}
		} else if val == 2 {
			if err := db.Model(&Commodity{}).Where("id =?", id).Updates(map[string]interface{}{"Image2": image}).Error; err != nil {
				return err
			}
		} else if val == 3 {
			if err := db.Model(&Commodity{}).Where("id =?", id).Updates(map[string]interface{}{"Image3": image}).Error; err != nil {
			}
		}

		return nil
	}
	return errcode.NotFound
}

// DeleteCommodity 删除商品信息
func (c *Commodity) DeleteCommodity(db *gorm.DB, id string) error {
	if err := db.Model(&Commodity{}).Where("id =?", id).Update("IsRemove", true).Error; err != nil {
		return err
	}
	return nil
}

// PostCommodity 添加商品信息
func (c *Commodity) PostCommodity(db *gorm.DB, commodity *Commodity) error {
	//先检查是否存在该商品
	if err := db.Where("Name =?", commodity.Name).First(&Commodity{}).Error; err == nil {
		return nil
	} else {
		if err := db.Create(&commodity).Error; err != nil {
			return err
		}
	}
	return nil
}

// CheckCommodity 检查是否存在该商品
func (c *Commodity) CheckCommodity(db *gorm.DB, id string) bool {
	var commodity Commodity
	if err := db.Where("ID =?", id).First(&commodity).Error; err != nil {
		return false
	}
	return true
}

// CheckCommodityName 检查是否存在该商品
func (c *Commodity) CheckCommodityName(db *gorm.DB, name string) bool {
	var commodity Commodity
	if err := db.Where("Name =?", name).First(&commodity).Error; err != nil {
		return false
	}
	return true
}

// CheckCommodityType 检查是否存在该类商品
func CheckCommodityType(db *gorm.DB, categoryID string) bool {
	var commodity []Commodity
	if err := db.Where("CategoryID =?", categoryID).Find(&commodity).Error; err != nil {
		return false
	}
	if len(commodity) == 0 {
		return false
	}
	return true
}

// GetCommodityByCategoryID 根据商品种类得到那一类商品信息
func (c *Commodity) GetCommodityByCategoryID(db *gorm.DB, categoryID string) ([]Commodity, error) {
	var commodities []Commodity
	if err := db.Where("CategoryID =? and IsRemove =? and Status = ?", categoryID, false, 0).Find(&commodities).Error; err != nil {
		return nil, err
	}
	return commodities, nil
}

// GetHomePageOnRecommend 得到所有HomePageOnRecommend字段为1 并且没有删除的商品
func (c *Commodity) GetHomePageOnRecommend(db *gorm.DB) ([]Commodity, error) {
	var commodities []Commodity
	if err := db.Where("HomePageOnRecommend =? and IsRemove =? and Status = ?", 1, false, 0).Find(&commodities).Error; err != nil {
		return nil, err
	}
	return commodities, nil
}

// GetCommodityByBatch 查询传入的批次的商品的信息
func (c *Commodity) GetCommodityByBatch(db *gorm.DB, val int) ([]Commodity, error) {
	arr, err := c.GetCommodity(db)
	if err != nil {
		return nil, err
	}
	val2 := val * 4
	// 4 8
	// len = 7
	if val2 >= len(arr) {
		return arr[val2-4:], nil
	}
	if val2 < len(arr) {
		return arr[val2-4 : val2], nil
	}

	return arr, nil
}
