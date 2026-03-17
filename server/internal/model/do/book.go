// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Book is the golang structure of table hg_book for DAO operations like Where/Data.
type Book struct {
	g.Meta   `orm:"table:hg_book, do:true"`
	Id       any         //
	Name     any         //
	Type     any         //
	SaleTime *gtime.Time //
}
