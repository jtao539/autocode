package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jtao539/autocode/template/common/request"
	"github.com/jtao539/autocode/template/common/response"
	"github.com/jtao539/autocode/template/model"
	"github.com/jtao539/autocode/template/service"
	"strconv"
)

type DepartmentApi struct {
	serv service.DepartmentService
}

func (d *DepartmentApi) GetDepartmentList(c *gin.Context) {
	var req model.DepartmentReq
	if err := request.BindJson(&req, c); err != nil {
		return
	}
	err, list := d.serv.GetDepartmentList(req)
	response.Pack(err, c, list)
}

func (d *DepartmentApi) GetDepartmentById(c *gin.Context) {
	ids := c.Param("id")
	id, _ := strconv.Atoi(ids)
	err, department := d.serv.GetDepartmentById(id)
	response.Pack(err, c, department)
}

func (d *DepartmentApi) AddDepartment(c *gin.Context) {
	var req model.DepartmentReq
	if err := request.BindJson(&req, c); err != nil {
		return
	}
	err := d.serv.AddDepartment(req)
	response.Pack(err, c)
}

func (d *DepartmentApi) AddDepartmentForm(c *gin.Context) {
	var req model.DepartmentReq
	if err := request.BindForm(&req, c); err != nil {
		return
	}
	err := d.serv.AddDepartment(req)
	response.Pack(err, c)
}

func (d *DepartmentApi) DeleteDepartment(c *gin.Context) {
	var req model.DepartmentReq
	if err := request.BindJson(&req, c); err != nil {
		return
	}
	err := d.serv.DeleteDepartmentById(req)
	response.Pack(err, c)
}

func (d *DepartmentApi) UpdateDepartment(c *gin.Context) {
	var req model.DepartmentReq
	if err := request.BindJson(&req, c); err != nil {
		return
	}
	err := d.serv.UpdateDepartment(req)
	response.Pack(err, c)
}
