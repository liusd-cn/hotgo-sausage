// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Book is the golang structure for table book.
type Book struct {
	Id       int         `json:"id"       orm:"id"        description:""`
	Name     string      `json:"name"     orm:"name"      description:""`
	Type     int         `json:"type"     orm:"type"      description:""`
	SaleTime *gtime.Time `json:"saleTime" orm:"sale_time" description:""`
}
