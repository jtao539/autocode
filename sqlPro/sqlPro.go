package sqlPro

import (
	"github.com/jmoiron/sqlx"
)

type SqlPro struct {
	DB *sqlx.DB
}

// Select 列表查找, dest为要查找的数据类型数组, request为查询条件结构体, table 为表名称, tags为需要跳过的字段
func (s *SqlPro) Select(dest interface{}, request interface{}, table string, tags ...string) error {
	str, params := safeSelect(request, table, tags...)
	err := s.DB.Select(&dest, str, params...)
	return err
}

// SelectWithFactor 可手动介入查询条件的列表查找, dest为要查找的数据类型数组, request为查询条件结构体, table 为表名称, factors 为sql条件, tags为需要跳过的字段
func (s *SqlPro) SelectWithFactor(dest interface{}, request interface{}, table string, factors []string, tags ...string) error {
	str, params := safeSelectWithFactor(request, table, factors, tags...)
	err := s.DB.Select(&dest, str, params...)
	return err
}

// Update 数据更新, request为新数据的结构体, entity为SQLNULL实体 table 为表名称， 通过对比request和entity获取跳过的字段, tx 为事务支持
func (s *SqlPro) Update(request interface{}, entity interface{}, table string, tx ...*sqlx.Tx) error {
	var err error
	str, params := safeUpdate(request, entity, table)
	if len(tx) > 0 {
		_, err = tx[0].Exec(str, params)
	} else {
		_, err = s.DB.Exec(str, params)
	}
	return err
}

// UpdateP 数据更新, request为新数据的结构体, entity为SQLNULL实体 table 为表名称， fields为需要跳过更新的字段, 通过对比request和entity获取跳过的字段, tx 为事务支持
func (s *SqlPro) UpdateP(request interface{}, entity interface{}, table string, fields []string, tx ...*sqlx.Tx) error {
	var err error
	str, params := safeUpdateP(request, entity, table, fields...)
	if len(tx) > 0 {
		_, err = tx[0].Exec(str, params)
	} else {
		_, err = s.DB.Exec(str, params)
	}
	return err
}
