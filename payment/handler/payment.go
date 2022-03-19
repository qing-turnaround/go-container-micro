package handler

import (
	"context"

	"github.com/xing-you-ji/go-container-micro/payment/domain/model"
	"github.com/xing-you-ji/go-container-micro/payment/domain/service"
	. "github.com/xing-you-ji/go-container-micro/payment/proto/payment"
)

type Payment struct {
	PaymentDataService service.IPaymentDataService
}

func (e *Payment) AddPayment(ctx context.Context, request *PaymentInfo, response *PaymentID) error {
	payment := &model.Payment{}
	paymentID, err := e.PaymentDataService.AddPayment(payment)
	if err != nil {
		return err
	}
	response.PaymentId = paymentID
	return nil
}

func (e *Payment) UpdatePayment(context.Context, *PaymentInfo, *Response) error {

}
func (e *Payment) DeletePaymentByID(context.Context, *PaymentID, *Response) error {

}
func (e *Payment) FindPaymentByID(context.Context, *PaymentID, *PaymentInfo) error {}
func (e *Payment) FindAllPayment(context.Context, *All, *PaymentAll) error         {}
