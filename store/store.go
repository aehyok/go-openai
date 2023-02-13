package store

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrNotFound = errors.New("not found")
	ErrExist    = errors.New("exist")
)

type Book struct {
	Id      string   `json:"id"`      // 图书ISBN ID
	Name    string   `json:"name"`    // 图书名称
	Authors []string `json:"authors"` // 图书作者
	Press   string   `json:"press"`   // 出版社
}

type Store interface {
	Create(*Book) error       // 创建一个新图书条目
	Update(*Book) error       // 更新某图书条目
	Get(string) (Book, error) // 获取某图书信息
	GetAll() ([]Book, error)  // 获取所有图书信息
	Delete(string) error      // 删除某图书条目
}

// 结构体
type List struct {
	gorm.Model        // 主键
	Name       string `gorm:"type:varchar(20); not null" json:"name" binding:"required"`
	State      string `gorm:"type:varchar(20); not null" json:"state" binding:"required"`
	Phone      string `gorm:"type:varchar(20); not null" json:"phone" binding:"required"`
	Email      string `gorm:"type:varchar(40); not null" json:"email" binding:"required"`
	Address    string `gorm:"type:varchar(200); not null" json:"address" binding:"required"`
}
