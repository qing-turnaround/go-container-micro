package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	log "github.com/micro/go-micro/v2/logger"
	cart "github.com/xing-you-ji/go-container-micro/cart/proto/cart"
	cartApi "github.com/xing-you-ji/go-container-micro/cartApi/proto/cartApi"
)

type CartApi struct {
	CartService cart.CartService
}

func (e *CartApi) FindAll(ctx context.Context, request *cartApi.Request, response *cartApi.Response) error {
	log.Info("接收到 /cartApi/findAll 访问请求")
	if _, ok := request.Get["user_id"]; !ok {
		return errors.New("参数错误")
	}
	userIdString := request.Get["user_id"].Value[0]
	fmt.Println("userID: ", userIdString)
	userID, err := strconv.ParseInt(userIdString, 10, 64)
	if err != nil {
		return err
	}

	// 获取购物车的所有商品
	cartAll, err := e.CartService.GetAll(context.TODO(), &cart.CartFindAll{UserId: userID})

	// 序列化
	bytes, err := json.Marshal(cartAll)
	if err != nil {
		return err
	}
	response.StatusCode = 200
	response.Body = string(bytes)
	return nil
}
