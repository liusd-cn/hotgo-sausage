// Package book
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package book

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询书籍列表
type ListReq struct {
	g.Meta `path:"/book/list" method:"get" tags:"书籍" summary:"获取书籍列表"`
	sysin.BookListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.BookListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出书籍列表
type ExportReq struct {
	g.Meta `path:"/book/export" method:"get" tags:"书籍" summary:"导出书籍列表"`
	sysin.BookListInp
}

type ExportRes struct{}

// ViewReq 获取书籍指定信息
type ViewReq struct {
	g.Meta `path:"/book/view" method:"get" tags:"书籍" summary:"获取书籍指定信息"`
	sysin.BookViewInp
}

type ViewRes struct {
	*sysin.BookViewModel
}

// EditReq 修改/新增书籍
type EditReq struct {
	g.Meta `path:"/book/edit" method:"post" tags:"书籍" summary:"修改/新增书籍"`
	sysin.BookEditInp
}

type EditRes struct{}

// DeleteReq 删除书籍
type DeleteReq struct {
	g.Meta `path:"/book/delete" method:"post" tags:"书籍" summary:"删除书籍"`
	sysin.BookDeleteInp
}

type DeleteRes struct{}