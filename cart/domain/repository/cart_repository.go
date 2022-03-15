package repository
import (
	"github.com/jinzhu/gorm"
	"cart/domain/model"
)
type ICartRepository interface{
    InitTable() error
    FindCartByID(int64) (*model.Cart, error)
	CreateCart(*model.Cart) (int64, error)
	DeleteCartByID(int64) error
	UpdateCart(*model.Cart) error
	FindAll()([]model.Cart,error)

}
//创建cartRepository
func NewCartRepository(db *gorm.DB) ICartRepository  {
	return &CartRepository{mysqlDb:db}
}

type CartRepository struct {
	mysqlDb *gorm.DB
}

//初始化表
func (u *CartRepository)InitTable() error  {
	return u.mysqlDb.CreateTable(&model.Cart{}).Error
}

//根据ID查找Cart信息
func (u *CartRepository)FindCartByID(cartID int64) (cart *model.Cart,err error) {
	cart = &model.Cart{}
	return cart, u.mysqlDb.Model(&model.Cart{}).First(cart,cartID)
}

//创建Cart信息
func (u *CartRepository) CreateCart(cart *model.Cart) (cartID int64,err error) {
	return cart.ID, u.mysqlDb.Create(cart).Error
}

//根据ID删除Cart信息
func (u *CartRepository) DeleteCartByID(cartID int64) err error {
	return u.mysqlDb.Where("ID = ?",cartID).Delete(&model.Cart{}).Error
}

//更新Cart信息
func (u *CartRepository) UpdateCart(cart *model.Cart) (err error) {
	return u.mysqlDb.Model(&cart).Update(cart).Error
}

//获取结果集
func (u *CartRepository) FindAll()(cartAll []model.Cart,err error) {
	return cartAll, u.mysqlDb.Find(&cartAll).Error
}

