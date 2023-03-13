package model

type GeekCourseType struct {
	Id           uint   `gorm:"column:id;type:bigint;primary_key;auto_increment" json:"id" binding:"required"`
	TypeId       string `gorm:"column:typeId;type:varchar(255); not null" json:"typeId" binding:"required"`
	TypeName     string `gorm:"column:typeName;type:varchar(255); not null" json:"typeName" binding:"required"`
	DisplayOrder string `gorm:"column:displayOrder;type:int(11); not null" json:"displayOrder" binding:"required"`
	IsDeleted    string `gorm:"column:isDeleted;type:int(1); not null" json:"isDeleted" binding:"required"`
}

type GeekCourse struct {
	Id           uint   `gorm:"column:id;type:bigint;primary_key;auto_increment" json:"id" binding:"required"`
	TypeId       string `gorm:"column:typeId;type:varchar(255); not null" json:"typeId" binding:"required"`
	Title        string `gorm:"column:title;type:varchar(255); not null" json:"title" binding:"required"`
	DisplayOrder string `gorm:"column:displayOrder;type:int(11); not null" json:"displayOrder" binding:"required"`
	Json         string `gorm:"column:json;type:int(1); not null" json:"json" binding:"required"`
}

type GeekArticle struct {
	Id           uint   `gorm:"column:id;type:bigint;primary_key;auto_increment" json:"id" binding:"required"`
	TypeId       string `gorm:"column:typeId;type:varchar(255); not null" json:"typeId" binding:"required"`
	Title        string `gorm:"column:title;type:varchar(255); not null" json:"title" binding:"required"`
	DisplayOrder string `gorm:"column:displayOrder;type:int(11); not null" json:"displayOrder" binding:"required"`
	Json         string `gorm:"column:json;type:int(1); not null" json:"json" binding:"required"`
}

