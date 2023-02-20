package model

import "gorm.io/gorm"

type BasicUser struct {
	gorm.Model        // 主键
	Account    string `gorm:"type:varchar(200); not null" json:"account" binding:"required"`
	Password   string `gorm:"type:varchar(200); not null" json:"password" binding:"required"`
	Name       string `gorm:"type:varchar(200); not null" json:"name" binding:"required"`
	State      string `gorm:"type:varchar(20); not null" json:"state" binding:"required"`
	Phone      string `gorm:"type:varchar(20); not null" json:"phone" binding:"required"`
	Email      string `gorm:"type:varchar(40); not null" json:"email" binding:"required"`
	Address    string `gorm:"type:varchar(200); not null" json:"address" binding:"required"`
}

type GeekProduct struct {
	gorm.Model        // 主键
	Title      string `gorm:"type:varchar(20); not null" json:"title" binding:"required"`
	Remark     string `gorm:"type:varchar(20); not null" json:"remark" binding:"required"`
}
