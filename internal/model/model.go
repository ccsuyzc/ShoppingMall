package model

import (
	"ShoppingMall/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//type Model struct {
//	ID         uint32 `gorm:"primary_key" json:"id"`
//	CreatedBy  string `json:"created_by"`
//	ModifiedBy string `json:"modified_by"`
//	CreatedOn  uint32 `json:"created_on"`
//	ModifiedOn uint32 `json:"modified_on"`
//	DeletedOn  uint32 `json:"deleted_on"`
//	IsDel      uint8  `json:"is_del"`
//}
//
//type Article struct {
//	*Model
//	Title         string `json:"title"`
//	Desc          string `json:"desc"`
//	Content       string `json:"content"`
//	CoverImageUrl string `json:"cover_image_url"`
//	State         uint8  `json:"state"`
//}
//
//func (a Article) TableName() string {
//	return "blog_article"
//}
//
//type ArticleTag struct {
//	*Model
//	TagID     uint32 `json:"tag_id"`
//	ArticleID uint32 `json:"article_id"`
//}

//func (a ArticleTag) TableName() string {
//	return "blog_article_tag"
//}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf(s,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))
	if err != nil {
		return nil, err
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}
