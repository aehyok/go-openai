package store

import "errors"

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
