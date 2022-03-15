package handler

import (
	"context"

	"github.com/xing-you-ji/go-container-micro/cart/domain/model"
	"github.com/xing-you-ji/go-container-micro/cart/domain/service"
	. "github.com/xing-you-ji/go-container-micro/cart/proto/cart"
	"github.com/xing-you-ji/go-container-micro/common"
)

type Cart struct {
	CartDataService service.ICartDataService
}

// AddCart 添加购物车服务
func (h *Cart) AddCart(ctx context.Context, request *CartInfo, response *ResponseAdd) (err error) {
	cart := &model.Cart{}
	common.SwapTo(request, cart)
	response.CartId, err = h.CartDataService.AddCart(cart)
	return err
}

// CleanCart 清空购物车服务
func (h *Cart) CleanCart(ctx context.Context, request *Clean, response *Response) error {
	if err := h.CartDataService.CleanCart(request.UserId); err != nil {
		return err
	}

	response.Meg = "清空购物车成功"
	return nil
}

// Incr 添加购物车服务
func (h *Cart) Incr(ctx context.Context, request *Item, response *Response) error {
	if err := h.CartDataService.IncrNum(request.Id, request.ChangeNum); err != nil {
		return err
	}
	response.Meg = "添加购物车成功"
	return nil
}

// Decr 减少购物车服务
func (h *Cart) Decr(ctx context.Context, request *Item, response *Response) error {
	if err := h.CartDataService.DecrNum(request.Id, request.ChangeNum); err != nil {
		return err
	}
	response.Meg = "减少购物车成功"
	return nil
}

// DeleteItemByID 通过购物车ID删除服务
func (h *Cart) DeleteItemByID(ctx context.Context, request *CartID, response *Response) error {
	if err := h.CartDataService.DeleteCart(request.Id); err != nil {
		return err
	}
	response.Meg = "购物车删除成功"
	return nil
}

// GetAll 获取购物车所有
func (h *Cart) GetAll(ctx context.Context, request *CartFindAll, response *CartAll) error {
	cartAll, err := h.CartDataService.FindAllCart(request.UserId)
	if err != nil {
		return err
	}
	return common.SwapCartTo(cartAll, response)
}
