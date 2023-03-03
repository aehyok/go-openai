package model

import "gorm.io/gorm"

type GeekCourseType struct {
	gorm.Model          // 主键
	TypeId       string `gorm:"column:typeId;type:varchar(255); not null" json:"typeId" binding:"required"`
	TypeName     string `gorm:"column:typeName;type:varchar(255); not null" json:"typeName" binding:"required"`
	DisplayOrder string `gorm:"column:displayOrder;type:int(11); not null" json:"displayOrder" binding:"required"`
	IsDeleted    string `gorm:"column:isDeleted;type:int(1); not null" json:"isDeleted" binding:"required"`
}

type GeekCourse struct {
	gorm.Model          // 主键
	TypeId       string `gorm:"column:typeId;type:varchar(255); not null" json:"typeId" binding:"required"`
	Title        string `gorm:"column:title;type:varchar(255); not null" json:"title" binding:"required"`
	DisplayOrder string `gorm:"column:displayOrder;type:int(11); not null" json:"displayOrder" binding:"required"`
	Json         string `gorm:"column:json;type:int(1); not null" json:"json" binding:"required"`
}
