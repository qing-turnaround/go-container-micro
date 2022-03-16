package service

import (
	"order/domain/model"
	"order/domain/repository"
)

type IOrderDataService interface {
	AddOrder(*model.Order) (int64 , error)
	DeleteOrder(int64) error
	UpdateOrder(*model.Order) error
	FindOrderByID(int64) (*model.Order, error)
}


//创建
func NewOrderDataService(orderRepository repository.IOrderRepository) IOrderDataService{
	return &OrderService{ orderRepository }
}

type OrderService struct {
	OrderRepository repository.IOrderRepository
}


//插入
func (u *OrderService) AddOrder(order *model.Order) (orderID int64 ,err error) {
	 return u.OrderRepository.CreateOrder(order)
}

//删除
func (u *OrderService) DeleteOrder(orderID int64) (err error) {
	return u.OrderRepository.DeleteOrderByID(orderID)
}

//更新
func (u *OrderService) UpdateOrder(order *model.Order) (err error) {
	return u.OrderRepository.UpdateOrder(order)
}

//查找
func (u *OrderService) FindOrderByID(orderID int64) (order *model.Order,err error) {
	return u.OrderRepository.FindOrderByID(orderID)
}

