package model

import (
	"github.com/jtao539/autocode/template/common/request"
)

type Department struct {
}

type DepartmentDTO struct {
}

type DepartmentReq struct {
	request.PageInfo
	DepartmentDTO
}

func (Department) TableName() string {
	return "tbl_department"
}
