package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/xing-you-ji/go-container-micro/category/domain/model"
)

type ICategoryRepository interface {
	InitTable() error
	FindCategoryByID(int64) (*model.Category, error)
	CreateCategory(*model.Category) (int64, error)
	DeleteCategoryByID(int64) error
	UpdateCategory(*model.Category) error
	FindAll() ([]model.Category, error)
	FindCategoryByName(string) (*model.Category, error)
	FindCategoryByLevel(uint32) ([]model.Category, error)
	FindCategoryByParent(int64) ([]model.Category, error)
}

// NewCategoryRepository 创建categoryRepository
func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &CategoryRepository{mysqlDb: db}
}

type CategoryRepository struct {
	mysqlDb *gorm.DB
}

// InitTable 初始化表
func (u *CategoryRepository) InitTable() error {
	return u.mysqlDb.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
		CreateTable(&model.Category{}).Error
}

// FindCategoryByID 根据ID查找Category信息
func (u *CategoryRepository) FindCategoryByID(categoryID int64) (category *model.Category, err error) {
	category = &model.Category{}
	return category, u.mysqlDb.Model(&model.Category{}).First(category, categoryID).Error
}

// CreateCategory 创建Category信息
func (u *CategoryRepository) CreateCategory(category *model.Category) (categoryID int64, err error) {
	return category.ID, u.mysqlDb.Create(category).Error
}

// DeleteCategoryByID 根据ID删除Category信息
func (u *CategoryRepository) DeleteCategoryByID(categoryID int64) error {
	return u.mysqlDb.Where("ID = ?", categoryID).Delete(&model.Category{}).Error
}

// UpdateCategory 更新Category信息
func (u *CategoryRepository) UpdateCategory(category *model.Category) (err error) {
	return u.mysqlDb.Model(&category).Update(category).Error
}

// FindAll 获取结果集
func (u *CategoryRepository) FindAll() (categoryAll []model.Category, err error) {
	return categoryAll, u.mysqlDb.Find(&categoryAll).Error
}

// FindCategoryByName 通过Name获取结果
func (u *CategoryRepository) FindCategoryByName(categoryName string) (category *model.Category, err error) {
	category = new(model.Category)
	return category, u.mysqlDb.Where("category_name = ?",
		categoryName).Find(category).Error
}

func (u *CategoryRepository) FindCategoryByLevel(level uint32) (categorySlice []model.Category, err error) {
	return categorySlice, u.mysqlDb.Where("category_level = ?",
		level).Find(categorySlice).Error
}

func (u *CategoryRepository) FindCategoryByParent(parent int64) (categorySlice []model.Category, err error) {
	return categorySlice, u.mysqlDb.Where("category_parent = ?",
		parent).Find(categorySlice).Error
}
