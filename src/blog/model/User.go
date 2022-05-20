package model

type User struct {
	UserId int32 `gorm:"auto_increment;primary;column:UserId"`
}
