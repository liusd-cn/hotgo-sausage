// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package sys

import (
	"context"
	"hotgo/api/admin/book"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

var (
	Book = cBook{}
)

type cBook struct{}

// List 查看书籍列表
func (c *cBook) List(ctx context.Context, req *book.ListReq) (res *book.ListRes, err error) {
	list, totalCount, err := service.SysBook().List(ctx, &req.BookListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*sysin.BookListModel{}
	}

	res = new(book.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出书籍列表
func (c *cBook) Export(ctx context.Context, req *book.ExportReq) (res *book.ExportRes, err error) {
	err = service.SysBook().Export(ctx, &req.BookListInp)
	return
}

// Edit 更新书籍
func (c *cBook) Edit(ctx context.Context, req *book.EditReq) (res *book.EditRes, err error) {
	err = service.SysBook().Edit(ctx, &req.BookEditInp)
	return
}

// View 获取指定书籍信息
func (c *cBook) View(ctx context.Context, req *book.ViewReq) (res *book.ViewRes, err error) {
	data, err := service.SysBook().View(ctx, &req.BookViewInp)
	if err != nil {
		return
	}

	res = new(book.ViewRes)
	res.BookViewModel = data
	return
}

// Delete 删除书籍
func (c *cBook) Delete(ctx context.Context, req *book.DeleteReq) (res *book.DeleteRes, err error) {
	err = service.SysBook().Delete(ctx, &req.BookDeleteInp)
	return
}