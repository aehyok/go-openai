package model

import "gorm.io/gorm"

type GeekCourseType struct {
	gorm.Model          // 主键
	TypeId       string `gorm:"column:typeId;type:varchar(255); not null" json:"typeId" binding:"required"`
	TypeName     string `gorm:"column:typeName;type:varchar(255); not null" json:"typeName" binding:"required"`
	DisplayOrder string `gorm:"column:displayOrder;type:int(11); not null" json:"displayOrder" binding:"required"`
	IsDeleted    string `gorm:"column:isDeleted;type:int(1); not null" json:"isDeleted" binding:"required"`
}
