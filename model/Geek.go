package model

import "gorm.io/gorm"

type GeekCourseType struct {
	gorm.Model          // 主键
	TypeId       string `gorm:"type:varchar(255); not null" json:"type_id" binding:"required"`
	TypeName     string `gorm:"type:varchar(255); not null" json:"type_name" binding:"required"`
	DisplayOrder string `gorm:"type:int(11); not null" json:"displayOrder" binding:"required"`
	IsDeleted    string `gorm:"type:int(1); not null" json:"isDeleted" binding:"required"`
}
