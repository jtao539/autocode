package model

import "github.com/jtao539/sqlxp"

type Department struct {
}

type DepartmentDTO struct {
	Id int `json:"id"`
}

type DepartmentReq struct {
	sqlxp.PageInfo
	DepartmentDTO
}

func (Department) TableName() string {
	return "tbl_department"
}
