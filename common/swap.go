package common

import (
	"encoding/json"
	"log"

	cartModel "github.com/xing-you-ji/go-container-micro/cart/domain/model"
	cart "github.com/xing-you-ji/go-container-micro/cart/proto/cart"
	categoryModel "github.com/xing-you-ji/go-container-micro/category/domain/model"
	category "github.com/xing-you-ji/go-container-micro/category/proto/category"
	productModel "github.com/xing-you-ji/go-container-micro/product/domain/model"
	product "github.com/xing-you-ji/go-container-micro/product/proto/product"
)

// SwapTo 通过json tag进行反序列化
func SwapTo(request, category interface{}) (err error) {

	dataBytes, err := json.Marshal(request)
	if err != nil {
		return err
	}
	return json.Unmarshal(dataBytes, category)
}

// SwapCategoryTo 切片的SwapTo
func SwapCategoryTo(categorySlice []categoryModel.Category, response *category.FindAllResponse) (err error) {
	for _, v := range categorySlice {
		categoryResponse := &category.CategoryResponse{}
		if SwapTo(v, categoryResponse) != nil {
			log.Println(err)
			return err
		}
		response.Category = append(response.Category, categoryResponse)
	}
	return nil
}

// SwapProductTo 切片的SwapTo
func SwapProductTo(productAll []productModel.Product, response *product.AllProduct) (err error) {
	for _, v := range productAll {
		productInfo := &product.ProductInfo{}
		if SwapTo(v, productInfo) != nil {
			log.Println(err)
			return err
		}
		response.ProductInfo = append(response.ProductInfo, productInfo)
	}
	return nil
}

// SwapCartTo 切片的SwapTo
func SwapCartTo(cartAll []cartModel.Cart, response *cart.CartAll) (err error) {
	for _, v := range cartAll {
		cart := &cart.CartInfo{}
		if SwapTo(v, cart) != nil {
			log.Println(err)
			return err
		}
		response.CartInfo = append(response.CartInfo, cart)
	}
	return nil
}
