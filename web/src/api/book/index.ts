import { http, jumpExport } from '@/utils/http/axios';

// 获取书籍列表
export function List(params) {
  return http.request({
    url: '/book/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除书籍
export function Delete(params) {
  return http.request({
    url: '/book/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑书籍
export function Edit(params) {
  return http.request({
    url: '/book/edit',
    method: 'POST',
    params,
  });
}

// 获取书籍指定详情
export function View(params) {
  return http.request({
    url: '/book/view',
    method: 'GET',
    params,
  });
}

// 导出书籍
export function Export(params) {
  jumpExport('/book/export', params);
}