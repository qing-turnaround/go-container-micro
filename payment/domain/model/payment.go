package model

type Payment struct{
	ID int64 `gorm:"primary_key;not_null;auto_increment"`
}

