package sysController

import (
	"context"
	"github.com/SupenBysz/gf-admin/api_v1"
	"github.com/SupenBysz/gf-admin/api_v1/sysapi"
	"github.com/SupenBysz/gf-admin/model"
	"github.com/SupenBysz/gf-admin/service"
)

// SysOrganization 组织架
var SysOrganization = cSysOrganization{}

type cSysOrganization struct{}

// QueryOrganizationList 获取组织架构|列表
func (c *cSysOrganization) QueryOrganizationList(ctx context.Context, req *sysapi.QueryOrganizationListReq) (*sysapi.OrganizationInfoListRes, error) {
	return service.SysOrganization().QueryOrganizationList(ctx, req.SearchFilter)
}

// GetOrganizationList 根据ID获取下级组织架构|列表
func (c *cSysOrganization) GetOrganizationList(ctx context.Context, req *sysapi.GetOrganizationListReq) (*sysapi.OrganizationInfoListRes, error) {
	data, count, err := service.SysOrganization().GetOrganizationList(ctx, req.Id, req.IsRecursive)
	if err != nil {
		return nil, err
	}

	return &sysapi.OrganizationInfoListRes{
		List: data,
		PaginationRes: model.PaginationRes{
			Pagination: model.Pagination{
				Page:     1,
				PageSize: count,
			},
			PageTotal: 1,
		},
	}, err
}

// GetOrganizationTree 根据ID获取下级组织架构|树
func (c *cSysOrganization) GetOrganizationTree(ctx context.Context, req *sysapi.GetOrganizationTreeReq) (*sysapi.OrganizationInfoTreeListRes, error) {
	result, err := service.SysOrganization().GetOrganizationTree(ctx, req.Id)

	return (*sysapi.OrganizationInfoTreeListRes)(result), err
}

// CreateOrganizationInfo 创建或更新组织架构|信息
func (c *cSysOrganization) CreateOrganizationInfo(ctx context.Context, req *sysapi.CreateOrganizationInfoReq) (*sysapi.OrganizationInfoRes, error) {
	result, err := service.SysOrganization().CreateOrganizationInfo(ctx, req.SysOrganizationInfo)
	return (*sysapi.OrganizationInfoRes)(result), err
}

// UpdateOrganizationInfo 创建或更新组织架构|信息
func (c *cSysOrganization) UpdateOrganizationInfo(ctx context.Context, req *sysapi.UpdateOrganizationInfoReq) (*sysapi.OrganizationInfoRes, error) {
	result, err := service.SysOrganization().UpdateOrganizationInfo(ctx, req.SysOrganizationInfo)
	return (*sysapi.OrganizationInfoRes)(result), err
}

// GetOrganizationInfo 获取组织架构|信息
func (c *cSysOrganization) GetOrganizationInfo(ctx context.Context, req *sysapi.GetOrganizationInfoReq) (*sysapi.OrganizationInfoRes, error) {
	result, err := service.SysOrganization().GetOrganizationInfo(ctx, req.Id)

	return (*sysapi.OrganizationInfoRes)(result), err
}

// DeleteOrganizationInfo 根据ID删除组织架构
func (c *cSysOrganization) DeleteOrganizationInfo(ctx context.Context, req *sysapi.DeleteOrganizationInfoReq) (api_v1.BoolRes, error) {
	result, err := service.SysOrganization().DeleteOrganizationInfo(ctx, req.Id)

	return result == true, err
}
