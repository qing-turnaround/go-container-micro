package repository

import (
	"github.com/xing-you-ji/go-container-micro/product/domain/model"

	"github.com/jinzhu/gorm"
)

type IProductRepository interface {
	InitTable() error
	FindProductByID(int64) (*model.Product, error)
	CreateProduct(*model.Product) (int64, error)
	DeleteProductByID(int64) error
	UpdateProduct(*model.Product) error
	FindAll() ([]model.Product, error)
}

// NewProductRepository 创建productRepository
func NewProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{mysqlDb: db}
}

type ProductRepository struct {
	mysqlDb *gorm.DB
}

// InitTable 初始化表
func (u *ProductRepository) InitTable() error {
	// 创建四张表（并设置字符集）
	return u.mysqlDb.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
		CreateTable(&model.Product{}, &model.ProductImage{},
			&model.ProductSeo{}, &model.ProductSize{}).Error
}

// FindProductByID 根据ID查找Product信息
func (u *ProductRepository) FindProductByID(productID int64) (product *model.Product, err error) {
	product = &model.Product{}
	// 因为有些信息在不同的表中，所以需哟使用Preload来加载其他的表
	return product, u.mysqlDb.Preload("ProductImage'").
		Preload("ProductSeo").Preload("ProductSize").First(product, productID).Error
}

// CreateProduct 创建Product信息
func (u *ProductRepository) CreateProduct(product *model.Product) (productID int64, err error) {
	return product.ID, u.mysqlDb.Create(product).Error
}

// DeleteProductByID 根据ID删除Product信息
func (u *ProductRepository) DeleteProductByID(productID int64) error {
	// 开启事务
	tx := u.mysqlDb.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // 如果错误就回滚
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}
	// 删除
	if err := tx.Unscoped().Where("id = ?", productID).Delete(&model.Product{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Unscoped().Where("images_product_id = ?", productID).Delete(&model.ProductImage{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Unscoped().Where("size_product_id = ?", productID).Delete(&model.ProductSize{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Unscoped().Where("seo_product_id = ?", productID).Delete(&model.ProductSeo{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// UpdateProduct 更新Product信息
func (u *ProductRepository) UpdateProduct(product *model.Product) (err error) {
	return u.mysqlDb.Model(&product).Update(product).Error
}

// FindAll 获取结果集
func (u *ProductRepository) FindAll() (productAll []model.Product, err error) {
	// 同样需要用Preload进行关联
	return productAll, u.mysqlDb.Preload("ProductImage'").
		Preload("ProductSeo").Preload("ProductSize").Find(&productAll).Error
}
