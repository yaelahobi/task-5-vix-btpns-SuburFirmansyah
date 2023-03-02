package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint `gorm:"primary_key;auto_increment"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Photos    []Photo `gorm:"OnDelete:CASCADE;"`
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	err := db.Create(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) BeforeSave(db *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) GetUser(db *gorm.DB, uid uint) (*User, error) {
	err := db.Where("id = ?", uid).First(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) UpdateUser(db *gorm.DB, uid uint) (*User, error) {
	err := db.Model(&u).Where("id = ?", uid).Updates(User{Name: u.Name, Email: u.Email}).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) UpdatePassword(db *gorm.DB, uid uint) (*User, error) {
	err := db.Model(&u).Where("id = ?", uid).Update("password", u.Password).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) DeleteUser(db *gorm.DB, uid uint) (int64, error) {
	db = db.Where("id = ?", uid).Delete(&User{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
