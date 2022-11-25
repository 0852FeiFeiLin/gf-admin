package userEventState

import "github.com/SupenBysz/gf-admin/model/enum"

var (
	BeforeCreate = kyEnum.New(2, "创建前")
	AfterCreate  = kyEnum.New(4, "创建后")
)
