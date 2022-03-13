package common

import (
	"encoding/json"
	"log"

	"github.com/xing-you-ji/go-container-micro/category/domain/model"
	category "github.com/xing-you-ji/go-container-micro/category/proto/category"
)

// SwapTo 通过json tag进行反序列化
func SwapTo(request, category interface{}) (err error) {
	dataBytes, err := json.Marshal(request)
	if err != nil {
		return err
	}
	return json.Unmarshal(dataBytes, category)
}

//
func SwapSliceTo(categorySlice []model.Category, response *category.FindAllResponse) (err error) {
	for _, cg := range categorySlice {
		cr := &category.CategoryResponse{}
		if SwapTo(cg, cr) != nil {
			log.Println(err)
			return err
		}
		response.Category = append(response.Category, cr)
	}
	return nil
}
