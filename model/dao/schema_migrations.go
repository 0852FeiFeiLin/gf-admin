// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/SupenBysz/gf-admin/model/dao/internal"
)

// internalSchemaMigrationsDao is internal type for wrapping internal DAO implements.
type internalSchemaMigrationsDao = *internal.SchemaMigrationsDao

// schemaMigrationsDao is the data access object for table schema_migrations.
// You can define custom methods on it to extend its functionality as you wish.
type schemaMigrationsDao struct {
	internalSchemaMigrationsDao
}

var (
	// SchemaMigrations is globally public accessible object for table schema_migrations operations.
	SchemaMigrations = schemaMigrationsDao{
		internal.NewSchemaMigrationsDao(),
	}
)

// Fill with you ideas below.
