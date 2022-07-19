package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/jtao539/autocode/template/model"
	"runtime"
)

type DepartmentDB struct {
	entity model.Department
}

func (d *DepartmentDB) GetDepartmentList(department model.DepartmentReq) (err error, list []model.Department, total int) {
	err, total = Pro.Select(&list, department, d.entity.TableName())
	return
}

func (d *DepartmentDB) GetDepartmentNameById(id int) (err error, department string) {
	str := fmt.Sprintf("select name from %s where id=?", d.entity.TableName())
	err = Pro.DB.Get(&department, str, id)
	return
}

func (d *DepartmentDB) GetDepartmentById(id int) (err error, department model.Department) {
	err = Pro.GetOneById(&department, d.entity.TableName(), id)
	return
}

func (d *DepartmentDB) AddDepartment(department model.Department, tx ...*sqlx.Tx) error {
	return Pro.InsertOne(department, d.entity.TableName(), tx...)
}

func (d *DepartmentDB) DeleteDepartmentById(id int, tx ...*sqlx.Tx) error {
	return Pro.DeleteOneById(d.entity.TableName(), id, tx...)
}

func (d *DepartmentDB) UpdateDepartment(department model.DepartmentDTO, tx ...*sqlx.Tx) error {
	return Pro.Update(department, d.entity, d.entity.TableName(), tx...)
}

// 获取正在运行的函数名
func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}
