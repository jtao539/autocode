package service

import (
	"github.com/jtao539/autocode/template/common/response"
	"github.com/jtao539/autocode/template/common/syserror"
	"github.com/jtao539/autocode/template/db"
	"github.com/jtao539/autocode/template/model"
	"github.com/jtao539/sqlxp"
)

type DepartmentService struct {
	repos db.DepartmentDB
}

func (d *DepartmentService) GetDepartmentList(departmentReq model.DepartmentReq) (error error, result interface{}) {
	err, list, total := d.repos.GetList(departmentReq)
	if err != nil {
		return err, nil
	}
	dtoList := make([]model.DepartmentDTO, len(list))
	sqlxp.N2BList(list, dtoList)
	result = response.ListData{List: dtoList, Total: total}
	return
}

func (d *DepartmentService) GetDepartmentById(id int) (error error, result interface{}) {
	err, m := d.repos.GetOne(id)
	if err != nil {
		return err, nil
	}
	var department model.DepartmentDTO
	sqlxp.N2B(m, &department)
	result = department
	return
}

func (d *DepartmentService) AddDepartment(departmentReq model.DepartmentReq) error {
	var department model.Department
	sqlxp.B2N(departmentReq.DepartmentDTO, &department)
	return d.repos.Add(department)
}

func (d *DepartmentService) DeleteDepartmentById(departmentReq model.DepartmentReq) error {
	dto := departmentReq.DepartmentDTO
	return d.repos.DeleteById(dto.Id)
}

func (d *DepartmentService) UpdateDepartment(departmentReq model.DepartmentReq) error {
	dto := departmentReq.DepartmentDTO
	if dto.Id == 0 {
		return syserror.InValidUpdateError
	}
	return d.repos.Update(dto)
}
