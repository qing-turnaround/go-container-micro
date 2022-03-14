package model

type Product struct{
	ID int64 `gorm:"primary_key;not_null;auto_increment"`
}

