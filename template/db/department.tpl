package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/jtao539/autocode/template/common/zlog"
	"github.com/jtao539/autocode/template/model"
)

type DepartmentDB struct {
	entity model.Department
}

func (d *DepartmentDB) GetList(req model.DepartmentReq) (err error, list []model.Department, total int) {
	err, total = Pro.Select(&list, req, d.entity.TableName())
	if err != nil {
		zlog.SugarLogger.Error(err)
	}
	return
}

func (d *DepartmentDB) GetOne(id int) (err error, department model.Department) {
	err = Pro.GetOneById(&department, d.entity.TableName(), id)
	if err != nil {
		zlog.SugarLogger.Error(err)
	}
	return
}

func (d *DepartmentDB) Add(department model.Department, tx ...*sqlx.Tx) error {
	err := Pro.InsertOne(department, d.entity.TableName(), tx...)
	if err != nil {
		zlog.SugarLogger.Error(err)
	}
	return err
}

func (d *DepartmentDB) Update(department model.DepartmentDTO, tx ...*sqlx.Tx) error {
	err := Pro.Update(department, d.entity, d.entity.TableName(), tx...)
	if err != nil {
		zlog.SugarLogger.Error(err)
	}
	return err
}

func (d *DepartmentDB) DeleteById(id int, tx ...*sqlx.Tx) error {
	err := Pro.DeleteOneById(d.entity.TableName(), id, tx...)
	if err != nil {
		zlog.SugarLogger.Error(err)
	}
	return err
}
