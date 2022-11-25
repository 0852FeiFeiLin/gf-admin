// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/SupenBysz/gf-admin/model/dao/internal"
)

// internalSysUserDao is internal type for wrapping internal DAO implements.
type internalSysUserDao = *internal.SysUserDao

// sysUserDao is the data access object for table sys_user.
// You can define custom methods on it to extend its functionality as you wish.
type sysUserDao struct {
	internalSysUserDao
}

var (
	// SysUser is globally public accessible object for table sys_user operations.
	SysUser = sysUserDao{
		internal.NewSysUserDao(),
	}
)

// Fill with you ideas below.
