// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package sysin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/os/gtime"
)

// BookUpdateFields 修改书籍字段过滤
type BookUpdateFields struct {
	Name     string      `json:"name"     dc:"name"`
	Type     int         `json:"type"     dc:"type"`
	SaleTime *gtime.Time `json:"saleTime" dc:"sale_time"`
}

// BookInsertFields 新增书籍字段过滤
type BookInsertFields struct {
	Name     string      `json:"name"     dc:"name"`
	Type     int         `json:"type"     dc:"type"`
	SaleTime *gtime.Time `json:"saleTime" dc:"sale_time"`
}

// BookEditInp 修改/新增书籍
type BookEditInp struct {
	entity.Book
}

func (in *BookEditInp) Filter(ctx context.Context) (err error) {

	return
}

type BookEditModel struct{}

// BookDeleteInp 删除书籍
type BookDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *BookDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type BookDeleteModel struct{}

// BookViewInp 获取指定书籍信息
type BookViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *BookViewInp) Filter(ctx context.Context) (err error) {
	return
}

type BookViewModel struct {
	entity.Book
}

// BookListInp 获取书籍列表
type BookListInp struct {
	form.PageReq
	Id int `json:"id" dc:"id"`
}

func (in *BookListInp) Filter(ctx context.Context) (err error) {
	return
}

type BookListModel struct {
	Id       int         `json:"id"       dc:"id"`
	Name     string      `json:"name"     dc:"name"`
	Type     int         `json:"type"     dc:"type"`
	SaleTime *gtime.Time `json:"saleTime" dc:"sale_time"`
}

// BookExportModel 导出书籍
type BookExportModel struct {
	Id       int         `json:"id"       dc:"id"`
	Name     string      `json:"name"     dc:"name"`
	Type     int         `json:"type"     dc:"type"`
	SaleTime *gtime.Time `json:"saleTime" dc:"sale_time"`
}