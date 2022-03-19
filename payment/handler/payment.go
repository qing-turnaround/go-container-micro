package handler

import (
	"context"

	"github.com/xing-you-ji/go-container-micro/common"
	"github.com/xing-you-ji/go-container-micro/payment/domain/model"
	"github.com/xing-you-ji/go-container-micro/payment/domain/service"
	. "github.com/xing-you-ji/go-container-micro/payment/proto/payment"
	"go.uber.org/zap"
)

type Payment struct {
	PaymentDataService service.IPaymentDataService
}

func (e *Payment) AddPayment(ctx context.Context, request *PaymentInfo, response *PaymentID) error {
	payment := &model.Payment{}
	paymentID, err := e.PaymentDataService.AddPayment(payment)
	if err != nil {
		zap.L().Error("handler.AddPayment error: ", zap.Error(err))
		return err
	}
	response.PaymentId = paymentID
	return nil
}

func (e *Payment) UpdatePayment(ctx context.Context, request *PaymentInfo, response *Response) error {
	payment := &model.Payment{}
	if err := common.SwapTo(request, payment); err != nil {
		zap.L().Error("handler.UpdatePayment error", zap.Error(err))
		return err
	}
	return e.PaymentDataService.UpdatePayment(payment)
}
func (e *Payment) DeletePaymentByID(ctx context.Context, request *PaymentID, response *Response) error {
	return e.PaymentDataService.DeletePayment(request.PaymentId)
}
func (e *Payment) FindPaymentByID(ctx context.Context, request *PaymentID, response *PaymentInfo) error {
	payment, err := e.PaymentDataService.FindPaymentByID(request.PaymentId)
	if err != nil {
		zap.L().Error("handler.FindPaymentByID error: ", zap.Error(err))
		return err
	}
	return common.SwapTo(payment, response)
}

func (e *Payment) FindAllPayment(ctx context.Context, request *All, response *PaymentAll) error {
	paymentAll, err := e.PaymentDataService.FindAllPayment()
	if err != nil {
		zap.L().Error("handler.FindAllPayment error: ", zap.Error(err))
	}
	for _, v := range paymentAll {
		paymentInfo := &PaymentInfo{}
		if err := common.SwapTo(v, paymentInfo); err != nil {
			zap.L().Error("handler.FindAllPayment error: ", zap.Error(err))
			return err
		}
		response.PaymentInfo = append(response.PaymentInfo, paymentInfo)
	}
	return nil
}
