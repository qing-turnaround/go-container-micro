package service

import (
	"payment/domain/model"
	"payment/domain/repository"
)

type IPaymentDataService interface {
	AddPayment(*model.Payment) (int64 , error)
	DeletePayment(int64) error
	UpdatePayment(*model.Payment) error
	FindPaymentByID(int64) (*model.Payment, error)
}


//创建
func NewPaymentDataService(paymentRepository repository.IPaymentRepository) IPaymentDataService{
	return &PaymentService{ paymentRepository }
}

type PaymentService struct {
	PaymentRepository repository.IPaymentRepository
}


//插入
func (u *PaymentService) AddPayment(payment *model.Payment) (paymentID int64 ,err error) {
	 return u.PaymentRepository.CreatePayment(payment)
}

//删除
func (u *PaymentService) DeletePayment(paymentID int64) (err error) {
	return u.PaymentRepository.DeletePaymentByID(paymentID)
}

//更新
func (u *PaymentService) UpdatePayment(payment *model.Payment) (err error) {
	return u.PaymentRepository.UpdatePayment(payment)
}

//查找
func (u *PaymentService) FindPaymentByID(paymentID int64) (payment *model.Payment,err error) {
	return u.PaymentRepository.FindPaymentByID(paymentID)
}

