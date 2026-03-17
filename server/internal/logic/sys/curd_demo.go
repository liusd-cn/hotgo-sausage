// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.6
package sys

import (
	"context"
	"fmt"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/library/hgorm/hook"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysCurdDemo struct{}

func NewSysCurdDemo() *sSysCurdDemo {
	return &sSysCurdDemo{}
}

func init() {
	service.RegisterSysCurdDemo(NewSysCurdDemo())
}

// Model CURD列表ORM模型
func (s *sSysCurdDemo) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SysGenCurdDemo.Ctx(ctx), option...)
}

// List 获取CURD列表列表
func (s *sSysCurdDemo) List(ctx context.Context, in *sysin.CurdDemoListInp) (list []*sysin.CurdDemoListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.FieldsPrefix(dao.SysGenCurdDemo.Table(), sysin.CurdDemoListModel{})
	mod = mod.Fields(hgorm.JoinFields(ctx, sysin.CurdDemoListModel{}, &dao.TestCategory, "testCategory"))

	// 关联表字段
	mod = mod.LeftJoinOnFields(dao.TestCategory.Table(), dao.SysGenCurdDemo.Columns().CategoryId, "=", dao.TestCategory.Columns().Id)

	// 查询ID
	if in.Id > 0 {
		mod = mod.Where(dao.SysGenCurdDemo.Columns().Id, in.Id)
	}

	// 查询标题
	if in.Title != "" {
		mod = mod.WhereLike(dao.SysGenCurdDemo.Columns().Title, "%"+in.Title+"%")
	}

	// 查询描述
	if in.Description != "" {
		mod = mod.WhereLike(dao.SysGenCurdDemo.Columns().Description, "%"+in.Description+"%")
	}

	// 查询创建者
	if in.CreatedBy != "" {
		ids, err := service.AdminMember().GetIdsByKeyword(ctx, in.CreatedBy)
		if err != nil {
			return nil, 0, err
		}
		mod = mod.WhereIn(dao.SysGenCurdDemo.Columns().CreatedBy, ids)
	}

	// 查询删除者
	if in.DeletedBy != "" {
		ids, err := service.AdminMember().GetIdsByKeyword(ctx, in.DeletedBy)
		if err != nil {
			return nil, 0, err
		}
		mod = mod.WhereIn(dao.SysGenCurdDemo.Columns().DeletedBy, ids)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.SysGenCurdDemo.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 查询关联分类
	if in.TestCategoryName != "" {
		mod = mod.WherePrefix(dao.TestCategory.Table(), dao.TestCategory.Columns().Name, in.TestCategoryName)
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderAsc(dao.SysGenCurdDemo.Table() + "." + dao.SysGenCurdDemo.Columns().Sort).OrderDesc(dao.SysGenCurdDemo.Table() + "." + dao.SysGenCurdDemo.Columns().Id)

	// 操作人摘要信息
	mod = mod.Hook(hook.MemberSummary)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取CURD列表列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出CURD列表
func (s *sSysCurdDemo) Export(ctx context.Context, in *sysin.CurdDemoListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.CurdDemoExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出CURD列表-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.CurdDemoExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增CURD列表
func (s *sSysCurdDemo) Edit(ctx context.Context, in *sysin.CurdDemoEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			in.UpdatedBy = contexts.GetUserId(ctx)
			if _, err = s.Model(ctx).
				Fields(sysin.CurdDemoUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改CURD列表失败，请稍后重试！")
			}
			return
		}

		// 新增
		in.CreatedBy = contexts.GetUserId(ctx)
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(sysin.CurdDemoInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增CURD列表失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除CURD列表
func (s *sSysCurdDemo) Delete(ctx context.Context, in *sysin.CurdDemoDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.SysGenCurdDemo.Columns().DeletedBy: contexts.GetUserId(ctx),
		dao.SysGenCurdDemo.Columns().DeletedAt: gtime.Now(),
	}).Unscoped().Update(); err != nil {
		err = gerror.Wrap(err, "删除CURD列表失败，请稍后重试！")
		return
	}
	return
}

// MaxSort 获取CURD列表最大排序
func (s *sSysCurdDemo) MaxSort(ctx context.Context, in *sysin.CurdDemoMaxSortInp) (res *sysin.CurdDemoMaxSortModel, err error) {
	if err = dao.SysGenCurdDemo.Ctx(ctx).Fields(dao.SysGenCurdDemo.Columns().Sort).OrderDesc(dao.SysGenCurdDemo.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取CURD列表最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(sysin.CurdDemoMaxSortModel)
	}

	res.Sort = form.DefaultMaxSort(res.Sort)
	return
}

// View 获取CURD列表指定信息
func (s *sSysCurdDemo) View(ctx context.Context, in *sysin.CurdDemoViewInp) (res *sysin.CurdDemoViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Hook(hook.MemberSummary).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取CURD列表信息，请稍后重试！")
		return
	}
	return
}

// Status 更新CURD列表状态
func (s *sSysCurdDemo) Status(ctx context.Context, in *sysin.CurdDemoStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.SysGenCurdDemo.Columns().Status:    in.Status,
		dao.SysGenCurdDemo.Columns().UpdatedBy: contexts.GetUserId(ctx),
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新CURD列表状态失败，请稍后重试！")
		return
	}
	return
}

// Switch 更新CURD列表开关
func (s *sSysCurdDemo) Switch(ctx context.Context, in *sysin.CurdDemoSwitchInp) (err error) {
	var fields = []string{
		dao.SysGenCurdDemo.Columns().Switch,
		// ...
	}

	if !validate.InSlice(fields, in.Key) {
		err = gerror.New("开关键名不在白名单")
		return
	}

	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		in.Key:                                 in.Value,
		dao.SysGenCurdDemo.Columns().UpdatedBy: contexts.GetUserId(ctx),
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新CURD列表开关失败，请稍后重试！")
		return
	}
	return
}

// ImportTemplate 下载CURD列表导入模板
func (s *sSysCurdDemo) ImportTemplate(ctx context.Context, in *sysin.CurdDemoImportTemplateInp) (err error) {
	// 获取导出模型的字段描述作为模板表头
	tags, err := convert.GetEntityDescTags(sysin.CurdDemoExportModel{})
	if err != nil {
		return gerror.Wrap(err, "获取模板表头失败")
	}

	fileName := "CURD列表导入模板"
	sheetName := "导入模板"

	err = excel.ExportTemplate(ctx, tags, fileName, sheetName)
	if err != nil {
		return gerror.Wrap(err, "导出导入模板失败")
	}
	return
}

// Import 导入CURD列表数据
func (s *sSysCurdDemo) Import(ctx context.Context, in *sysin.CurdDemoImportInp) (res *sysin.CurdDemoImportModel, err error) {
	res = &sysin.CurdDemoImportModel{
		FailDetails: make([]string, 0),
	}

	// 解析Excel文件
	var importData []sysin.CurdDemoInsertFields
	err = excel.ImportFromFile(in.File, "Sheet1", &importData)
	if err != nil {
		return nil, gerror.Wrap(err, "解析Excel文件失败")
	}

	if len(importData) == 0 {
		return nil, gerror.New("Excel文件中没有有效数据")
	}

	// 使用事务批量导入
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		userId := contexts.GetUserId(ctx)

		for i, item := range importData {
			lineNum := i + 2 // Excel行号，从第2行开始是数据

			// 验证数据 - 复用Edit方法中的验证逻辑
			editInp := &sysin.CurdDemoEditInp{
				SysGenCurdDemo: entity.SysGenCurdDemo{
					Title:       item.Title,
					Description: item.Description,
					Content:     item.Content,
					Image:       item.Image,
					Attachfile:  item.Attachfile,
					CityId:      item.CityId,
					Switch:      item.Switch,
					Sort:        item.Sort,
				},
			}

			if err := editInp.Filter(ctx); err != nil {
				res.FailCount++
				res.FailDetails = append(res.FailDetails, fmt.Sprintf("第%d行: %s", lineNum, err.Error()))
				continue
			}

			// 设置创建人
			item.CreatedBy = userId

			// 插入数据
			_, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
				Fields(sysin.CurdDemoInsertFields{}).
				Data(item).OmitEmptyData().Insert()
			
			if err != nil {
				res.FailCount++
				res.FailDetails = append(res.FailDetails, fmt.Sprintf("第%d行: 数据插入失败 - %s", lineNum, err.Error()))
				continue
			}

			res.SuccessCount++
		}

		// 如果有失败的记录，回滚事务
		if res.FailCount > 0 {
			return gerror.New(fmt.Sprintf("有%d条数据导入失败", res.FailCount))
		}

		return nil
	})

	// 事务执行失败，但我们已经收集了错误详情，所以不返回错误
	if err != nil && res.FailCount == 0 {
		return nil, gerror.Wrap(err, "批量导入事务执行失败")
	}

	return res, nil
}