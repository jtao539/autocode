package sqlPro

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const Tag = "db"

func CommonInsert(o interface{}, tbl string) string {
	sql := "INSERT INTO "
	sql += tbl + "("
	t := reflect.TypeOf(o)
	for i := 0; i < t.NumField(); i++ {
		tagName := t.Field(i).Tag.Get(Tag)
		sql += tagName + ","
	}
	sql = sql[:strings.LastIndex(sql, ",")]
	sql += ") VALUES ("
	for i := 0; i < t.NumField(); i++ {
		tagName := t.Field(i).Tag.Get(Tag)
		sql += ":" + tagName + ","
	}
	sql = sql[:strings.LastIndex(sql, ",")]
	sql += ")"
	return sql
}

// CommonUpdateO 通用更新语句生成，o 为DTO, tbl 为表名称， args 为需要跳过的字段
func CommonUpdateO(o interface{}, tbl string, args ...string) string {
	sql := "UPDATE "
	sql += tbl + " SET "
	t := reflect.TypeOf(o)
	v := reflect.ValueOf(o)
	for i := 0; i < t.NumField(); i++ {
		tagName := t.Field(i).Tag.Get(Tag)
		if tagName == "id" || containArray(tagName, args) {
			continue
		}
		sql += tagName + "="
		switch v.Field(i).Kind() {
		case reflect.Int:
			if v.Field(i).Int() == 0 {
				sql += "NULL" + ","
			} else {
				sql += strconv.FormatInt(v.Field(i).Int(), 10) + ","
			}
		case reflect.String:
			if v.Field(i).String() == "" {
				sql += "NULL" + ","
			} else {
				sql += "'" + v.Field(i).String() + "'" + ","
			}
		}
	}
	sql = sql[:strings.LastIndex(sql, ",")]
	sql += " where id=" + strconv.FormatInt(v.FieldByName("Id").Int(), 10)
	return sql
}

// CommonUpdate 通用更新语句生成，o 为DTO, a 为entity tbl 为表名称， 通过对比o和a获取跳过的字段
func CommonUpdate(o interface{}, a interface{}, tbl string) string {
	return localUpdate(o, a, tbl, func(tagName string) bool {
		return contain(tagName, "id", "code", "status")
	})
}

// CommonUpdateP 通用更新语句生成，o 为DTO, a 为entity tbl 为表名称， 通过对比o和a获取跳过的字段，args 为需要跳过的字段
func CommonUpdateP(o interface{}, a interface{}, tbl string, args ...string) string {
	return localUpdate(o, a, tbl, func(tagName string) bool {
		return containArray(tagName, args)
	}, func(tagName string) bool {
		return contain(tagName, "id", "code", "status")
	})
}

// chooseUpdate 选择某些字段更新 args 为选择的字段
func chooseUpdate(o interface{}, a interface{}, tbl string, args ...string) string {
	return localUpdate(o, a, tbl, func(tagName string) bool {
		return !containArray(tagName, args)
	})
}

// localUpdate 通用更新语句生成，o 为DTO, a 为entity tbl 为表名称， 通过对比o和a获取跳过的字段， fs为所有满足条件需要跳过的函数
func localUpdate(o interface{}, a interface{}, tbl string, fs ...func(tagName string) bool) string {
	sql := "UPDATE "
	sql += tbl + " SET "
	t := reflect.TypeOf(o)
	v := reflect.ValueOf(o)
	for i := 0; i < t.NumField(); i++ {
		tagName := t.Field(i).Tag.Get(Tag)
		if noContain(tagName, a) {
			continue
		}
		flag := false
		for i := 0; i < len(fs); i++ {
			if fs[i](tagName) {
				flag = true
			}
		}
		if flag {
			continue
		}
		sql += tagName + "="
		switch v.Field(i).Kind() {
		case reflect.Int:
			if v.Field(i).Int() == 0 {
				sql += "NULL" + ","
			} else {
				sql += strconv.FormatInt(v.Field(i).Int(), 10) + ","
			}
		case reflect.String:
			if v.Field(i).String() == "" {
				sql += "NULL" + ","
			} else {
				sql += "'" + v.Field(i).String() + "'" + ","
			}
		case reflect.Float64:
			if v.Field(i).Float() == 0 {
				sql += "NULL" + ","
			} else {
				value := strconv.FormatFloat(v.Field(i).Float(), 'g', 15, 64)
				sql += "'" + value + "'" + ","
			}
		case reflect.Array, reflect.Slice:
			v := v.Field(i)
			var result string
			for i := 0; i < v.Len(); i++ {
				result += v.Index(i).String() + ","
			}
			result = result[:strings.LastIndex(result, ",")]
			sql += "'" + result + "'" + ","
		}
	}
	sql = sql[:strings.LastIndex(sql, ",")]
	sql += " where id=" + strconv.FormatInt(v.FieldByName("Id").Int(), 10)
	fmt.Println(sql)
	return sql
}

// CommonSelect 条件查询语句生成, o 为DTO, a 为entity tbl 为表名称, tags为跳过的查找字段
// str := SELECT * FROM T_Employee WHERE FNumber BETWEEN 'DEV001' AND 'DEV008' AND FSalary BETWEEN 3000 AND 6000
func CommonSelect(o interface{}, tbl string, tags ...string) string {
	sql := "SELECT * FROM " + tbl + " WHERE "
	ov := reflect.ValueOf(o)
	ot := reflect.TypeOf(o)
	var DTO reflect.Value
	for i := 0; i < ot.NumField(); i++ {
		if ot.Field(i).Name != "PageInfo" {
			DTO = ov.Field(i)
		}
	}
	PageInfo := ov.FieldByName("PageInfo")
	dt := DTO.Type()
	for i := 0; i < dt.NumField(); i++ {
		tagName := dt.Field(i).Tag.Get(Tag)
		if tagName == "id" || containArray(tagName, tags) {
			continue
		}
		switch DTO.Field(i).Kind() {
		case reflect.Int:
			if DTO.Field(i).Int() != 0 {
				sql += tagName + "=" + strconv.FormatInt(DTO.Field(i).Int(), 10) + " AND "
				// sql += tagName + " < 1 AND "
			}
		case reflect.String:
			if DTO.Field(i).String() != "" {
				sql += tagName + " like " + "'%" + DTO.Field(i).String() + "%'" + " AND "
			}
		case reflect.Float64:
			if DTO.Field(i).Float() != 0 {
				value := strconv.FormatFloat(DTO.Field(i).Float(), 'g', 15, 64)
				sql += tagName + "=" + value + " AND "
			}
		}
	}
	if PageInfo.FieldByName("CreateUserId").Int() != 0 {
		sql += "create_user_id" + "=" + strconv.FormatInt(PageInfo.FieldByName("CreateUserId").Int(), 10) + " AND "
	}
	if strings.Contains(sql, "AND") {
		sql = sql[:strings.LastIndex(sql, "AND")]
	} else {
		sql += "1=1"
	}
	page := PageInfo.FieldByName("Page").Int()
	pageSize := PageInfo.FieldByName("PageSize").Int()
	if page > 0 && pageSize > 0 {
		sql += " limit " + strconv.FormatInt((page-1)*pageSize, 10) + " , " + strconv.FormatInt(pageSize, 10)
	}
	return sql
}

func CommonSelectWithFactor(o interface{}, tbl string, factors []string, tags ...string) string {
	sql := "SELECT * FROM " + tbl + " WHERE "
	ov := reflect.ValueOf(o)
	ot := reflect.TypeOf(o)
	var DTO reflect.Value
	for i := 0; i < ot.NumField(); i++ {
		if ot.Field(i).Name != "PageInfo" {
			DTO = ov.Field(i)
		}
	}
	PageInfo := ov.FieldByName("PageInfo")
	dt := DTO.Type()
	for i := 0; i < dt.NumField(); i++ {
		tagName := dt.Field(i).Tag.Get(Tag)
		if tagName == "id" || containArray(tagName, tags) {
			continue
		}
		switch DTO.Field(i).Kind() {
		case reflect.Int:
			if DTO.Field(i).Int() != 0 {
				sql += tagName + "=" + strconv.FormatInt(DTO.Field(i).Int(), 10) + " AND "
				// sql += tagName + " < 1 AND "
			}
		case reflect.String:
			if DTO.Field(i).String() != "" {
				sql += tagName + " like " + "'%" + DTO.Field(i).String() + "%'" + " AND "
			}
		case reflect.Float64:
			if DTO.Field(i).Float() != 0 {
				value := strconv.FormatFloat(DTO.Field(i).Float(), 'g', 15, 64)
				sql += tagName + "=" + value + " AND "
			}
		}
	}
	if PageInfo.FieldByName("CreateUserId").Int() != 0 {
		sql += "create_user_id" + "=" + strconv.FormatInt(PageInfo.FieldByName("CreateUserId").Int(), 10) + " AND "
	}
	for i := 0; i < len(factors); i++ {
		sql += fmt.Sprintf(" %s AND ", factors[i])
	}
	if strings.Contains(sql, "AND") {
		sql = sql[:strings.LastIndex(sql, "AND")]
	} else {
		sql += "1=1"
	}
	page := PageInfo.FieldByName("Page").Int()
	pageSize := PageInfo.FieldByName("PageSize").Int()
	if page > 0 && pageSize > 0 {
		sql += " limit " + strconv.FormatInt((page-1)*pageSize, 10) + " , " + strconv.FormatInt(pageSize, 10)
	}
	return sql
}

func CommonSelectP(o interface{}, tbl string, factors ...string) string {
	sql := "SELECT * FROM " + tbl + " WHERE "
	ov := reflect.ValueOf(o)
	ot := reflect.TypeOf(o)
	var DTO reflect.Value
	for i := 0; i < ot.NumField(); i++ {
		if ot.Field(i).Name != "PageInfo" {
			DTO = ov.Field(i)
		}
	}
	dt := DTO.Type()
	for i := 0; i < dt.NumField(); i++ {
		tagName := dt.Field(i).Tag.Get(Tag)
		if tagName == "id" {
			continue
		}
		switch DTO.Field(i).Kind() {
		case reflect.Int:
			if DTO.Field(i).Int() != 0 {
				sql += tagName + "=" + strconv.FormatInt(DTO.Field(i).Int(), 10) + " AND "
				// sql += tagName + " < 1 AND "
			}
		case reflect.String:
			if DTO.Field(i).String() != "" {
				sql += tagName + "=" + "'" + DTO.Field(i).String() + "'" + " AND "
			}
		case reflect.Float64:
			if DTO.Field(i).Float() != 0 {
				value := strconv.FormatFloat(DTO.Field(i).Float(), 'g', 15, 64)
				sql += tagName + "=" + value + " AND "
			}
		}
	}
	if strings.Contains(sql, "AND") {
		sql = sql[:strings.LastIndex(sql, "AND")]
	} else {
		sql += "1=1"
	}
	PageInfo := ov.FieldByName("PageInfo")
	page := PageInfo.FieldByName("Page").Int()
	pageSize := PageInfo.FieldByName("PageSize").Int()
	if page > 0 && pageSize > 0 {
		sql += " limit " + strconv.FormatInt((page-1)*pageSize, 10) + " , " + strconv.FormatInt(pageSize, 10)
	}
	return sql
}

func contain(tagName string, args ...string) bool {
	for i := 0; i < len(args); i++ {
		if tagName == args[i] {
			return true
		}
	}
	return false
}

func containArray(tagName string, args []string) bool {
	if len(args) == 0 {
		return false
	}
	for i := 0; i < len(args); i++ {
		if tagName == args[i] {
			return true
		}
	}
	return false
}

func noContain(tagName string, o interface{}) bool {
	t := reflect.TypeOf(o)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get(Tag) == tagName {
			return false
		}
	}
	return true
}

func NowDate() string {
	return time.Now().Format("20060102")
}
