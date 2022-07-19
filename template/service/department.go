package service

import (
	"github.com/jtao539/autocode/template/common/commonError"
	"github.com/jtao539/autocode/template/common/response"
	"github.com/jtao539/autocode/template/db"
	"github.com/jtao539/autocode/template/model"
	"github.com/jtao539/autocode/util"
	"time"
)

type DepartmentService struct {
	repos db.DepartmentDB
}

func (d *DepartmentService) GetDepartmentList(departmentReq model.DepartmentReq) (error error, result interface{}) {
	err, list, total := d.repos.GetDepartmentList(departmentReq)
	if err != nil {
		error = err
		return
	}
	dtoList := make([]model.DepartmentDTO, len(list))
	for i := 0; i < len(list); i++ {
		util.Entity2DTO(list[i], &dtoList[i])
	}
	result = response.ListData{List: dtoList, Total: total}
	return
}

func (d *DepartmentService) GetDepartmentById(id int) (err error, result interface{}) {
	err, m := d.repos.GetDepartmentById(id)
	var department model.DepartmentDTO
	util.Entity2DTO(m, &department)
	result = department
	return
}

func (d *DepartmentService) AddDepartment(departmentReq model.DepartmentReq) error {
	var department model.Department
	util.DTO2Entity(departmentReq.DepartmentDTO, &department)
	if departmentReq.UserId != 0 {
		department.CreateUserId = util.IntToNullInt32(departmentReq.UserId)
	}
	department.CreateTime = util.IntToNullInt32(int(time.Now().Unix()))
	return d.repos.AddDepartment(department)
}

func (d *DepartmentService) DeleteDepartmentById(departmentReq model.DepartmentReq) error {
	dto := departmentReq.DepartmentDTO
	return d.repos.DeleteDepartmentById(dto.Id)
}

func (d *DepartmentService) UpdateDepartment(departmentReq model.DepartmentReq) error {
	dto := departmentReq.DepartmentDTO
	if dto.Id == 0 {
		return commonError.InValidUpdateError
	}
	return d.repos.UpdateDepartment(dto)
}
