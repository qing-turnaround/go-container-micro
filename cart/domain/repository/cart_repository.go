package repository

import (
	"errors"

	"github.com/xing-you-ji/go-container-micro/cart/domain/model"

	"github.com/jinzhu/gorm"
)

type ICartRepository interface {
	InitTable() error
	FindCartByID(int64) (*model.Cart, error)
	CreateCart(*model.Cart) (int64, error)
	DeleteCartByID(int64) error
	UpdateCart(*model.Cart) error
	FindAll(int64) ([]model.Cart, error)

	CleanCart(int64) error
	IncrNum(int64, int64) error
	DecrNum(int64, int64) error
}

// NewCartRepository 创建cartRepository
func NewCartRepository(db *gorm.DB) ICartRepository {
	return &CartRepository{mysqlDb: db}
}

type CartRepository struct {
	mysqlDb *gorm.DB
}

// InitTable 初始化表
func (u *CartRepository) InitTable() error {
	return u.mysqlDb.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
		CreateTable(&model.Cart{}).Error
}

// FindCartByID 根据ID查找Cart信息
func (u *CartRepository) FindCartByID(cartID int64) (*model.Cart, error) {
	cart := &model.Cart{}
	return cart, u.mysqlDb.Model(&model.Cart{}).First(cart, cartID).Error
}

// CreateCart 创建Cart信息
func (u *CartRepository) CreateCart(cart *model.Cart) (cartID int64, err error) {
	// 判断某个用户是否已经存在同商品同尺寸的购物车，若存在，则跳过，否则创建
	db := u.mysqlDb.FirstOrCreate(cart, model.Cart{
		ProductID: cart.ProductID,
		SizeID:    cart.SizeID,
		UserID:    cart.UserID,
	})
	if db.Error != nil {
		return 0, db.Error
	}
	if db.RowsAffected == 0 { // 判断插入操作是否有行数影响
		return 0, errors.New("购物车插入失败")
	}

	return cart.ID, nil
}

// DeleteCartByID 根据ID删除Cart信息
func (u *CartRepository) DeleteCartByID(cartID int64) (err error) {
	return u.mysqlDb.Where("ID = ?", cartID).Delete(&model.Cart{}).Error
}

// UpdateCart 更新Cart信息
func (u *CartRepository) UpdateCart(cart *model.Cart) (err error) {
	return u.mysqlDb.Model(&cart).Update(cart).Error
}

// FindAll 获取结果集
func (u *CartRepository) FindAll(userID int64) (cartAll []model.Cart, err error) {
	return cartAll, u.mysqlDb.Where("user_id", userID).Find(&cartAll).Error
}

// CleanCart 清空购物车
func (u *CartRepository) CleanCart(userID int64) error {
	return u.mysqlDb.Where("user_id", userID).Delete(&model.Cart{}).Error
}

// IncrNum 添加购物车某商品的数量
func (u *CartRepository) IncrNum(cartID int64, num int64) error {
	cart := &model.Cart{
		ID: cartID,
	}
	// 使用 SQL 表达式更新
	return u.mysqlDb.Model(cart).UpdateColumn("num", gorm.Expr("num + ?", num)).Error
}

// DecrNum 减少购物车某商品的数量
func (u *CartRepository) DecrNum(cartID int64, num int64) error {
	cart := model.Cart{ID: cartID}
	db := u.mysqlDb.Model(cart).Where("num >= ?", num).UpdateColumn("num",
		gorm.Expr("num - ", num))
	if db.Error != nil {
		return db.Error
	}

	if db.RowsAffected == 0 {
		return errors.New("减少失败")
	}
	return nil
}
