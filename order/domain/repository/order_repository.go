package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/xing-you-ji/go-container-micro/order/domain/model"
)

type IOrderRepository interface {
	InitTable() error
	FindOrderByID(int64) (*model.Order, error)
	CreateOrder(*model.Order) (int64, error)
	DeleteOrderByID(int64) error
	UpdateOrder(*model.Order) error
	FindAll() ([]model.Order, error)
	UpdateShipStatus(int64, int32) error
	UpdatePayStatus(int64, int32) error
}

// NewOrderRepository 创建orderRepository
func NewOrderRepository(db *gorm.DB) IOrderRepository {
	return &OrderRepository{mysqlDb: db}
}

type OrderRepository struct {
	mysqlDb *gorm.DB
}

// InitTable 初始化表
func (u *OrderRepository) InitTable() error {
	return u.mysqlDb.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
		CreateTable(&model.Order{}, &model.OrderDetail{}).Error
}

// FindOrderByID 根据ID查找Order信息
func (u *OrderRepository) FindOrderByID(orderID int64) (order *model.Order, err error) {
	order = &model.Order{}
	return order, u.mysqlDb.Preload("OrderDetail").First(order, orderID).Error
}

// CreateOrder 创建Order信息
func (u *OrderRepository) CreateOrder(order *model.Order) (orderID int64, err error) {
	return order.ID, u.mysqlDb.Create(order).Error
}

// DeleteOrderByID 根据ID删除Order信息
func (u *OrderRepository) DeleteOrderByID(orderID int64) error {
	// 事务
	tx := u.mysqlDb.Begin()
	// 遇到错误回滚
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Unscoped().Delete("id = ?", orderID).Delete(&model.Order{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Unscoped().Delete("order_id", orderID).Delete(&model.OrderDetail{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 提交事务
	return tx.Commit().Error
}

// UpdateOrder 更新Order信息
func (u *OrderRepository) UpdateOrder(order *model.Order) (err error) {
	return u.mysqlDb.Model(&order).Update(order).Error
}

// FindAll 获取结果集
func (u *OrderRepository) FindAll() (orderAll []model.Order, err error) {
	return orderAll, u.mysqlDb.Find(&orderAll).Error
}

func (u *OrderRepository) UpdateShipStatus(orderID int64, shipStatus int32) error {
	db := u.mysqlDb.Model(&model.Order{}).Where("id = ?", orderID).
		UpdateColumn("ship_status", shipStatus)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("更新失败")
	}
	return nil
}

func (u *OrderRepository) UpdatePayStatus(orderID int64, payStatus int32) error {
	db := u.mysqlDb.Model(&model.Order{}).Where("id = ?", orderID).
		UpdateColumn("pay_status", payStatus)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("更新失败")
	}
	return nil
}
