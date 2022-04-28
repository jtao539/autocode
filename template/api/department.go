package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/jtao539/autocode/template/common/definiteError"
	"github.com/jtao539/autocode/template/common/response"
	"github.com/jtao539/autocode/template/model"
	"github.com/jtao539/autocode/template/service"
	"log"
	"strconv"
)

type DepartmentApi struct {
	serv service.DepartmentService
}

func (d *DepartmentApi) GetDepartmentList(c *gin.Context) {
	var json model.DepartmentReq
	if err := c.ShouldBindJSON(&json); err != nil {
		log.Println(err)
		response.FailByJsonError(c)
		return
	}
	err, list := d.serv.GetDepartmentList(json)
	if err != nil {
		if Contain(err) {
			response.FailWithMessage(err.Error(), c)
			return
		} else {
			log.Println(err)
			response.Fail(c)
			return
		}
	} else {
		json.Page = 0
		_, allList := d.serv.GetDepartmentList(json)
		response.OkWithListData(response.GetLength(allList), list, c)
	}
}

func (d *DepartmentApi) GetDepartmentById(c *gin.Context) {
	ids := c.Param("id")
	id, _ := strconv.Atoi(ids)
	err, department := d.serv.GetDepartmentById(id)
	if err != nil {
		if Contain(err) {
			response.FailWithMessage(err.Error(), c)
			return
		} else {
			log.Println(err)
			response.Fail(c)
			return
		}
	} else {
		response.OkWithData(department, c)
	}
}

func (d *DepartmentApi) AddDepartment(c *gin.Context) {
	var json model.DepartmentReq
	if err := c.ShouldBindJSON(&json); err != nil {
		log.Println(err)
		response.FailByJsonError(c)
		return
	}
	err := d.serv.AddDepartment(json)
	if err != nil {
		if Contain(err) {
			response.FailWithMessage(err.Error(), c)
			return
		} else {
			log.Println(err)
			response.Fail(c)
			return
		}
	} else {
		response.Ok(c)
	}
}

func (d *DepartmentApi) AddDepartmentForm(c *gin.Context) {
	var form model.DepartmentReq
	MForm, err := c.MultipartForm()
	if err != nil {
		response.FailByFormError(c)
		return
	}
	err = response.Decoder.Decode(&form, MForm.Value)
	if err != nil {
		response.FailByFormError(c)
		return
	}
	err = d.serv.AddDepartment(form)
	if err != nil {
		if Contain(err) {
			response.FailWithMessage(err.Error(), c)
			return
		} else {
			fmt.Println(err)
			response.Fail(c)
			return
		}
	} else {
		response.Ok(c)
	}
}

func (d *DepartmentApi) DeleteDepartment(c *gin.Context) {
	var json model.DepartmentReq
	if err := c.ShouldBindJSON(&json); err != nil {
		log.Println(err)
		response.FailByJsonError(c)
		return
	}
	err := d.serv.DeleteDepartmentById(json)
	if err != nil {
		if Contain(err) {
			response.FailWithMessage(err.Error(), c)
			return
		} else {
			log.Println(err)
			response.Fail(c)
			return
		}
	} else {
		response.Ok(c)
	}
}

func (d *DepartmentApi) UpdateDepartment(c *gin.Context) {
	var json model.DepartmentReq
	if err := c.ShouldBindJSON(&json); err != nil {
		log.Println(err)
		response.FailByJsonError(c)
		return
	}
	err := d.serv.UpdateDepartment(json)
	if err != nil {
		if Contain(err) {
			response.FailWithMessage(err.Error(), c)
			return
		} else {
			log.Println(err)
			response.Fail(c)
			return
		}
	} else {
		response.Ok(c)
	}
}
