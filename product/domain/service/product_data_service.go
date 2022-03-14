package service

import (
	"product/domain/model"
	"product/domain/repository"
)

type IProductDataService interface {
	AddProduct(*model.Product) (int64 , error)
	DeleteProduct(int64) error
	UpdateProduct(*model.Product) error
	FindProductByID(int64) (*model.Product, error)
}


//创建
func NewProductDataService(productRepository repository.IProductRepository) IProductDataService{
	return &ProductService{ productRepository }
}

type ProductService struct {
	ProductRepository repository.IProductRepository
}


//插入
func (u *ProductService) AddProduct(product *model.Product) (productID int64 ,err error) {
	 return u.ProductRepository.CreateProduct(product)
}

//删除
func (u *ProductService) DeleteProduct(productID int64) (err error) {
	return u.ProductRepository.DeleteProductByID(productID)
}

//更新
func (u *ProductService) UpdateProduct(product *model.Product) (err error) {
	return u.ProductRepository.UpdateProduct(product)
}

//查找
func (u *ProductService) FindProductByID(productID int64) (product *model.Product,err error) {
	return u.ProductRepository.FindProductByID(productID)
}

