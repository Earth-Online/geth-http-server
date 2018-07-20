package model

type User struct {
	Address string `gorm:"primary_key"`
	Password string

}



