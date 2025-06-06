package model

import (
	"errors"
	"fmt"
	"gin-vue-blog/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Category Category `gorm:"foreignkey:Cid"`
	Tag      Tag      `gorm:"foreignkey:Tid"`
	Title    string   `json:"title" gorm:"type:varchar(100);not null"`
	Cid      int      `json:"cid" gorm:"type:int;not null;"`
	Tid      int      `json:"tid" gorm:"type:int;not null;"`
	Uid      int      `json:"uid" gorm:"type:int;not null;"`
	Desc     string   `json:"desc" gorm:"type:varchar(200)"`
	Content  string   `json:"content" gorm:"type:longtext;not null"`
	HTML     string   `json:"html" gorm:"type:longtext;not null"`
	Img      string   `json:"img" gorm:"type:varchar(200)"`
	View     int      `json:"view" gorm:"type:int;"`
}

//查询分类下的所有文章

func GetArticleByCategory(pageSize int, pageNum int, cid int) ([]Article, int, int64) {
	var cate Category
	var total int64
	result := db.Where("id = ?", cid).First(&cate)
	if result.RowsAffected == 0 {
		return []Article{}, errmsg.ERROR_CATEGORY_NOT_EXIST, 0
	}
	var cateArtList []Article
	if cate.ID == 0 {
		db.Order("created_at DESC").Preload("Tag").Preload("Category").Where("NOT ID = 1").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cateArtList)
		db.Model(&cateArtList).Where("NOT ID = 1").Count(&total)
		return cateArtList, errmsg.SUCCESS, total
	}
	result = db.Order("created_at DESC").Preload("Tag").Preload("Category").Where("NOT ID = 1").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("Cid = ?", cid).Find(&cateArtList)
	db.Model(&cateArtList).Where("NOT ID = 1").Where("Cid = ?", cid).Count(&total)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return []Article{}, errmsg.ERROR_ARTICLE_NOT_EXIST, total
		}
		return []Article{}, errmsg.ERROR, total
	}
	return cateArtList, errmsg.SUCCESS, total
}

//查询标签下的所有文章

func GetArticleByTag(pageSize int, pageNum int, tid int) ([]Article, int, int64) {
	var tag Tag
	var total int64
	result := db.Where("id = ?", tid).First(&tag)
	if result.RowsAffected == 0 {
		return []Article{}, errmsg.ERROR_TAG_NOT_EXIST, 0
	}
	var tagArtList []Article
	result = db.Order("created_at DESC").Preload("Tag").Preload("Category").Where("NOT ID = 1").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("Tid = ?", tid).Find(&tagArtList)
	db.Model(&tagArtList).Where("NOT ID = 1").Where("Tid = ?", tid).Count(&total)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return []Article{}, errmsg.ERROR_ARTICLE_NOT_EXIST, total
		}
		return []Article{}, errmsg.ERROR, total
	}
	return tagArtList, errmsg.SUCCESS, total
}

//查询单个文章

func SearchArticleByCid(id int) (Article, int) {
	var article Article
	err := db.Preload("Category").Preload("Tag").Where("id=?", id).First(&article).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return Article{}, errmsg.ERROR
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Article{}, errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	return article, errmsg.SUCCESS
}

//查询文章列表

func GetArticle(title string, pageSize int, pageNum int) ([]Article, int, int64) {
	var art []Article
	var total int64
	var err error
	if title != "" {
		err = db.Order("created_at DESC").Preload("Category").Preload("Tag").Where("title LIKE ? && NOT id = 1", title+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&art).Error
		db.Model(&art).Where("title LIKE ? AND NOT id = 1", title+"%").Count(&total)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("文章查询出错,%v", err)
			return nil, errmsg.ERROR, 0
		}
		return art, errmsg.SUCCESS, total
	} else {
		err = db.Order("created_at DESC").Preload("Category").Preload("Tag").Where("NOT id = 1").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&art).Error
		db.Model(&art).Where("NOT id = 1").Count(&total)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("文章查询出错,%v", err)
			return nil, errmsg.ERROR, 0
		}
		return art, errmsg.SUCCESS, total
	}
}

//添加文章

func CreateArticle(data *Article) int {
	data.View = 0
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//编辑文章

func EditArticle(id int, data *Article) int {
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["tid"] = data.Tid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	maps["html"] = data.HTML
	result := db.Model(&Article{}).Where("id = ?", id).Updates(maps)
	if result.Error != nil {
		return errmsg.ERROR
	}
	if result.RowsAffected == 0 {
		return errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	fmt.Println(result.RowsAffected)
	return errmsg.SUCCESS
}

//更新点击量

func UpdateView(id int, data *Article) int {
	var maps = make(map[string]interface{})
	maps["view"] = data.View
	result := db.Model(&Article{}).Where("id = ?", id).Updates(maps)
	if result.Error != nil {
		return errmsg.ERROR
	}
	if result.RowsAffected == 0 {
		return errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	fmt.Println(result.RowsAffected)
	return errmsg.SUCCESS
}

//删除文章

func DeleteArticle(id int) int {
	var art Article
	result := db.Delete(&art, id)
	if result.Error != nil {
		return errmsg.ERROR
	}
	if result.RowsAffected == 0 {
		return errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	return errmsg.SUCCESS
}
