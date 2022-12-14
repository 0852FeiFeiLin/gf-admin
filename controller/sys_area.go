package sysController

import (
	"context"
	"github.com/SupenBysz/gf-admin/api_v1/sysapi"
	"github.com/SupenBysz/gf-admin/model"
	"github.com/SupenBysz/gf-admin/service"
)

// SysArea 地区
var SysArea = cSysArea{}

type cSysArea struct{}

// GetAreaListByParentId 获取属于父级ID的地区列表
func (c *cSysArea) GetAreaListByParentId(ctx context.Context, req *sysapi.GetAreaListByParentIdReq) (*model.AreaListRes, error) {
	return service.Area().GetAreaListByParentId(ctx, req.ParentId)
}
