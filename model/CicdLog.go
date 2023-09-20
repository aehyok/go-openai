package model

import "time"

type CicdLog struct {
	Id         string    `gorm:"column:id;type:varchar(255);primary_key;" json:"id" binding:"required"`
	Project    string    `gorm:"column:project;type:varchar(255); not null" json:"project" binding:"required"`
	Content    string    `gorm:"column:content;type:varchar(255); not null" json:"content" binding:"required"`
	Version    string    `gorm:"column:version;type:varchar(255); not null" json:"version" binding:"required"`
	CreateTime time.Time `gorm:"column:createTime;type:datetime; not null" json:"createTime" binding:"required"`
	Type       string    `gorm:"column:type;type:varchar(255); not null" json:"type" binding:"required"`
}
