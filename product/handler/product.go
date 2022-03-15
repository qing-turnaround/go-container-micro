package handler

import (
	"context"

	"github.com/xing-you-ji/go-container-micro/common"

	"github.com/xing-you-ji/go-container-micro/product/domain/model"

	"github.com/xing-you-ji/go-container-micro/product/domain/service"
	. "github.com/xing-you-ji/go-container-micro/product/proto/product"
)

type Product struct {
	ProductDataService service.IProductDataService
}

// AddProduct 添加商品服务
func (h *Product) AddProduct(ctx context.Context, request *ProductInfo,
	response *ResponseProduct) error {
	product := &model.Product{}
	if err := common.SwapTo(request, product); err != nil {
		return err
	}
	productID, err := h.ProductDataService.AddProduct(product)
	if err != nil {
		return err
	}
	response.ProductId = productID
	return nil
}

// FindProductByID 通过ID查找商品服务
func (h *Product) FindProductByID(ctx context.Context, request *RequestID, response *ProductInfo) error {
	product, err := h.ProductDataService.FindProductByID(request.ProductId)
	if err != nil {
		return err
	}
	if err = common.SwapTo(product, response); err != nil {
		return err
	}
	return nil
}
func (h *Product) UpdateProduct(ctx context.Context, request *ProductInfo, response *Response) error {
	product := &model.Product{}
	if err := common.SwapTo(request, product); err != nil {
		return err
	}
	if err := h.ProductDataService.UpdateProduct(product); err != nil {
		return err
	}
	response.Msg = "更新成功！"
	return nil
}
func (h *Product) DeleteProductByID(ctx context.Context, request *RequestID, response *Response) error {
	if err := h.ProductDataService.DeleteProduct(request.ProductId); err != nil {
		return err
	}
	response.Msg = "删除成功！"

	return nil
}
func (h *Product) FindAllProduct(ctx context.Context, request *RequestAll, response *AllProduct) error {
	productAll, err := h.ProductDataService.FindAllProduct()
	if err != nil {
		return err
	}
	return common.SwapProductTo(productAll, response)
}
