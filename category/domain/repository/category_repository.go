package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/xing-you-ji/go-container-micro/category/domain/model"
)
type ICategoryRepository interface{
    InitTable() error
    FindCategoryByID(int64) (*model.Category, error)
	CreateCategory(*model.Category) (int64, error)
	DeleteCategoryByID(int64) error
	UpdateCategory(*model.Category) error
	FindAll()([]model.Category,error)

}
//创建categoryRepository
func NewCategoryRepository(db *gorm.DB) ICategoryRepository  {
	return &CategoryRepository{mysqlDb:db}
}

type CategoryRepository struct {
	mysqlDb *gorm.DB
}

//初始化表
func (u *CategoryRepository)InitTable() error  {
	return u.mysqlDb.CreateTable(&model.Category{}).Error
}

//根据ID查找Category信息
func (u *CategoryRepository)FindCategoryByID(categoryID int64) (category *model.Category,err error) {
	category = &model.Category{}
	return category, u.mysqlDb.Model(&model.Category{}).First(category,categoryID)
}

//创建Category信息
func (u *CategoryRepository) CreateCategory(category *model.Category) (categoryID int64,err error) {
	return category.ID, u.mysqlDb.Create(category).Error
}

//根据ID删除Category信息
func (u *CategoryRepository) DeleteCategoryByID(categoryID int64) err error {
	return u.mysqlDb.Where("ID = ?",categoryID).Delete(&model.Category{}).Error
}

//更新Category信息
func (u *CategoryRepository) UpdateCategory(category *model.Category) (err error) {
	return u.mysqlDb.Model(&category).Update(category).Error
}

//获取结果集
func (u *CategoryRepository) FindAll()(categoryAll []model.Category,err error) {
	return categoryAll, u.mysqlDb.Find(&categoryAll).Error
}

