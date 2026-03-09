package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(30);comment:用户名" json:"username"`
	Password string `gorm:"type:varchar(32);comment:密码" json:"password"`
}

func (u *User) GetUserById(db *gorm.DB, uid int64) error {
	return db.Where("id = ?", uid).Limit(1).Find(&u).Error
}
