package service

import (
	"category/domain/model"
	"category/domain/repository"
)

type ICategoryDataService interface {
	AddCategory(*model.Category) (int64 , error)
	DeleteCategory(int64) error
	UpdateCategory(*model.Category) error
	FindCategoryByID(int64) (*model.Category, error)
}


//创建
func NewCategoryDataService(categoryRepository repository.ICategoryRepository) ICategoryDataService{
	return &CategoryService{ categoryRepository }
}

type CategoryService struct {
	CategoryRepository repository.ICategoryRepository
}


//插入
func (u *CategoryService) AddCategory(category *model.Category) (categoryID int64 ,err error) {
	 return u.CategoryRepository.CreateCategory(category)
}

//删除
func (u *CategoryService) DeleteCategory(categoryID int64) (err error) {
	return u.CategoryRepository.DeleteCategoryByID(categoryID)
}

//更新
func (u *CategoryService) UpdateCategory(category *model.Category) (err error) {
	return u.CategoryRepository.UpdateCategory(category)
}

//查找
func (u *CategoryService) FindCategoryByID(categoryID int64) (category *model.Category,err error) {
	return u.CategoryRepository.FindCategoryByID(categoryID)
}

