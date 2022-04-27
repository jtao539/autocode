package sqlPro

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/jtao539/autocode/common/definiteError"
	"reflect"
)

type SqlPro struct {
	DB *sqlx.DB
}

// Select 列表查找, dest为要查找的数据类型数组, request为查询条件结构体, table 为表名称, tags为需要跳过的字段
func (s *SqlPro) Select(dest interface{}, request interface{}, table string, tags ...string) error {
	array := anythingToSlice(dest)
	str, params := SafeSelect(array, table, tags...)
	err := s.DB.Select(&dest, str, params...)
	return err
}

// SelectWithFactor 可手动介入查询条件的列表查找, dest为要查找的数据类型数组, request为查询条件结构体, table 为表名称, factors 为sql条件, tags为需要跳过的字段
func (s *SqlPro) SelectWithFactor(dest interface{}, request interface{}, table string, factors []string, tags ...string) error {
	str, params := SafeSelectWithFactor(request, table, factors, tags...)
	err := s.DB.Select(&dest, str, params...)
	return err
}

// Update 数据更新, request为新数据的结构体, entity为SQLNULL实体 table 为表名称， 通过对比request和entity获取跳过的字段, tx 为事务支持
func (s *SqlPro) Update(request interface{}, entity interface{}, table string, tx ...*sqlx.Tx) error {
	var err error
	var rows sql.Result
	str, params := SafeUpdate(request, entity, table)
	if len(tx) > 0 {
		rows, err = tx[0].Exec(str, params...)
	} else {
		rows, err = s.DB.Exec(str, params...)
	}
	if err != nil {
		return err
	}
	AffectedNum, _ := rows.RowsAffected()
	if AffectedNum == 0 {
		return definiteError.UpdateError
	}
	return err
}

// UpdateP 数据更新, request为新数据的结构体, entity为SQLNULL实体 table 为表名称， fields为需要跳过更新的字段, 通过对比request和entity获取跳过的字段, tx 为事务支持
func (s *SqlPro) UpdateP(request interface{}, entity interface{}, table string, fields []string, tx ...*sqlx.Tx) error {
	var err error
	var rows sql.Result
	str, params := SafeUpdateP(request, entity, table, fields...)
	if len(tx) > 0 {
		rows, err = tx[0].Exec(str, params...)
	} else {
		rows, err = s.DB.Exec(str, params...)
	}
	if err != nil {
		return err
	}
	AffectedNum, _ := rows.RowsAffected()
	if AffectedNum == 0 {
		return definiteError.UpdateError
	}
	return err
}

func (s *SqlPro) GetOneById(one interface{}, table string, id int) error {
	str := fmt.Sprintf("select * from %s where id=?", table)
	return s.DB.Get(&one, str, id)
}

func (s *SqlPro) InsertOne(one interface{}, table string, tx ...*sqlx.Tx) error {
	var err error
	var rows sql.Result
	str := CommonInsert(one, table)
	if len(tx) > 0 {
		rows, err = tx[0].NamedExec(str, one)
	} else {
		rows, err = s.DB.NamedExec(str, one)
	}
	if err != nil {
		return err
	}
	AffectedNum, _ := rows.RowsAffected()
	if AffectedNum == 0 {
		return definiteError.InsertError
	}
	return err
}

func anythingToSlice(a interface{}) []interface{} {
	v := reflect.ValueOf(a)
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		result := make([]interface{}, v.Len())
		for i := 0; i < v.Len(); i++ {
			result[i] = v.Index(i).Interface()
			t := reflect.TypeOf(result[i])
			fmt.Println("t = ", t)
		}
		return result
	default:
		panic("not supported")
	}
}
