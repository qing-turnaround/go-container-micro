package service

import (
	"cart/domain/model"
	"cart/domain/repository"
)

type ICartDataService interface {
	AddCart(*model.Cart) (int64 , error)
	DeleteCart(int64) error
	UpdateCart(*model.Cart) error
	FindCartByID(int64) (*model.Cart, error)
}


//创建
func NewCartDataService(cartRepository repository.ICartRepository) ICartDataService{
	return &CartService{ cartRepository }
}

type CartService struct {
	CartRepository repository.ICartRepository
}


//插入
func (u *CartService) AddCart(cart *model.Cart) (cartID int64 ,err error) {
	 return u.CartRepository.CreateCart(cart)
}

//删除
func (u *CartService) DeleteCart(cartID int64) (err error) {
	return u.CartRepository.DeleteCartByID(cartID)
}

//更新
func (u *CartService) UpdateCart(cart *model.Cart) (err error) {
	return u.CartRepository.UpdateCart(cart)
}

//查找
func (u *CartService) FindCartByID(cartID int64) (cart *model.Cart,err error) {
	return u.CartRepository.FindCartByID(cartID)
}

