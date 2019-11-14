package page

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

//默认分页数量
const defaultPerPage = 50

//Page 分页
type Page struct {
	PerPage     int //每页数量,默认50
	CurrentPage int //当前所在页码
	Total       int //总数量
	Skip        int //跨过
	TotalPage   int //总页数
}

// New 创建Page对象
func New(obj map[string]int) *Page {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("BuildPagepage err--------", err)
		}
	}()

	perPage := obj["perPage"]
	if perPage == 0 {
		perPage = defaultPerPage
	}
	currentPage := obj["currentPage"]
	if currentPage == 0 {
		currentPage = 1
	}

	skip := (currentPage - 1) * perPage
	// fmt.Println("currentPage", currentPage, "perPage", perPage, "skip", skip)

	return &Page{PerPage: perPage, CurrentPage: currentPage, Skip: skip}
}

//GetInfo 客户端分页信息
func (page *Page) GetInfo() map[string]int {
	info := make(map[string]int)

	defer func() {
		if err := recover(); err != nil {
			logrus.Error("[utils.page.GetInfo]", err)
		}
	}()

	info["perPage"] = page.PerPage
	info["total"] = page.Total
	info["totalPage"] = page.TotalPage
	info["currentPage"] = page.CurrentPage

	return info
}

//SetTotal 设置总页码
func (page *Page) SetTotal(total int) {
	page.Total = total
	remainder := total % page.PerPage
	page.TotalPage = total / page.PerPage

	if remainder != 0 {
		page.TotalPage++
	}
}
