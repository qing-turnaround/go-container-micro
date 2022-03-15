package common

import (
	"encoding/json"
	"log"

	"github.com/xing-you-ji/go-container-micro/product/domain/model"
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

// SwapSliceTo 切片的SwapTo
func SwapSliceTo(productAll []model.Product, response *product.AllProduct) (err error) {
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
