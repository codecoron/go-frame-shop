// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"go-frame-shop/internal/dao/internal"
)

// internalAdminInfoDao is internal type for wrapping internal DAO implements.
type internalAdminInfoDao = *internal.AdminInfoDao

// adminInfoDao is the data access object for table admin_info.
// You can define custom methods on it to extend its functionality as you wish.
type adminInfoDao struct {
	internalAdminInfoDao
}

var (
	// AdminInfo is globally public accessible object for table admin_info operations.
	AdminInfo = adminInfoDao{
		internal.NewAdminInfoDao(),
	}
)

// Fill with you ideas below.
