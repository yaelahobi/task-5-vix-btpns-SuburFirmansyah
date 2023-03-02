package models

type Photo struct {
	ID       uint `gorm:"primary_key;auto_increment"`
	Title    string
	Caption  string
	PhotoUrl string
	UserID   uint
}
