package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/xing-you-ji/go-container-micro/payment/domain/model"
)

type IPaymentRepository interface{
    InitTable() error
    FindPaymentByID(int64) (*model.Payment, error)
	CreatePayment(*model.Payment) (int64, error)
	DeletePaymentByID(int64) error
	UpdatePayment(*model.Payment) error
	FindAll()([]model.Payment,error)

}
// NewPaymentRepository 创建paymentRepository
func NewPaymentRepository(db *gorm.DB) IPaymentRepository  {
	return &PaymentRepository{mysqlDb:db}
}

type PaymentRepository struct {
	mysqlDb *gorm.DB
}

// InitTable 初始化表
func (u *PaymentRepository)InitTable() error  {
	return u.mysqlDb.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
		CreateTable(&model.Payment{}).Error
}

// FindPaymentByID 根据ID查找Payment信息
func (u *PaymentRepository)FindPaymentByID(paymentID int64) (payment *model.Payment,err error) {
	payment = &model.Payment{}
	return payment, u.mysqlDb.Model(&model.Payment{}).First(payment,paymentID)
}

// CreatePayment 创建Payment信息
func (u *PaymentRepository) CreatePayment(payment *model.Payment) (paymentID int64,err error) {
	return payment.ID, u.mysqlDb.Create(payment).Error
}

// DeletePaymentByID 根据ID删除Payment信息
func (u *PaymentRepository) DeletePaymentByID(paymentID int64) err error {
	return u.mysqlDb.Where("ID = ?",paymentID).Delete(&model.Payment{}).Error
}

// UpdatePayment 更新Payment信息
func (u *PaymentRepository) UpdatePayment(payment *model.Payment) (err error) {
	return u.mysqlDb.Model(&payment).Update(payment).Error
}

// FindAll 获取结果集
func (u *PaymentRepository) FindAll()(paymentAll []model.Payment,err error) {
	return paymentAll, u.mysqlDb.Find(&paymentAll).Error
}

