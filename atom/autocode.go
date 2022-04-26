package atom

import (
	"fmt"
	"github.com/jtao539/autocode/db"
	"os"
	"strings"
)

type Atom struct {
	Name    string
	TblName string
	Path    string
	ModName string
}

type Model struct {
	columnName    string `db:"column_name"`
	columnComment string `db:"column_comment"`
	dataType      string `db:"data_type"`
}

func (a *Atom) CreateModel() {
	str := fmt.Sprintf("select column_name,column_comment,data_type from information_schema.columns where table_name='%s' and table_schema='%s' ORDER BY ORDINAL_POSITION", a.TblName, db.Name)
	var list []Model
	// err := db.GDB.DB.Select(&list, str)
	query, err := db.GDB.DB.Query(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	for query.Next() {
		m := Model{}
		query.Scan(&m.columnName, &m.columnComment, &m.dataType)
		list = append(list, m)
	}
	code := fmt.Sprintf("package model\n\nimport (\n\t\"database/sql\"\n\t\"%s/model/request\"\n)\n\ntype %s struct {\n", a.ModName, a.Name)
	for i := 0; i < len(list); i++ {
		m := list[i]
		tag := fmt.Sprintf("`db:\"%s\" json:\"%s\"`", m.columnName, m.columnName)
		switch m.dataType {
		case "int", "tinyint":
			code += "\t" + Case2Camel(m.columnName) + " sql.NullInt32 " + tag + " // " + m.columnComment + "\n"
		case "varchar":
			code += "\t" + Case2Camel(m.columnName) + " sql.NullString " + tag + " // " + m.columnComment + "\n"
		case "decimal":
			code += "\t" + Case2Camel(m.columnName) + " sql.NullFloat64 " + tag + " // " + m.columnComment + "\n"
		}
	}
	line := fmt.Sprintf("}\n\ntype %sDTO struct {\n", a.Name)
	code += line
	for i := 0; i < len(list); i++ {
		m := list[i]
		tag := fmt.Sprintf("`db:\"%s\" json:\"%s\"`", m.columnName, m.columnName)
		if m.columnName == "create_time" || m.columnName == "create_user" || m.columnName == "create_user_id" {
			continue
		}
		switch m.dataType {
		case "int", "tinyint":
			code += "\t" + Case2Camel(m.columnName) + " int " + tag + " // " + m.columnComment + "\n"
		case "varchar":
			code += "\t" + Case2Camel(m.columnName) + " string " + tag + " // " + m.columnComment + "\n"
		case "decimal":
			code += "\t" + Case2Camel(m.columnName) + " float64 " + tag + " // " + m.columnComment + "\n"
		}
	}
	lastLine := fmt.Sprintf("}\n\ntype %sReq struct {\n\trequest.PageInfo\n\t%sDTO\n}\n\nfunc (%s) TableName() string {\n\treturn \"%s\"\n}", a.Name, a.Name, a.Name, a.TblName)
	code += lastLine
	fileName := LowFirst(a.Name)
	filePath := fmt.Sprintf("%s/model/%s.go", a.Path, fileName)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Write([]byte(code))
	defer f.Close()
	fmt.Println(filePath, "Model 完成")
}

func (a *Atom) createDB() {
	str := fmt.Sprintf("package db\n\nimport (\n\t\"%s/model\"\n\t\"%s/util\"\n\t\"database/sql\"\n\t\"fmt\"\n\t\"github.com/jmoiron/sqlx\"\n\t\"time\"\n)\n\ntype DepartmentDB struct {\n\tentity model.Department\n}\n\nfunc (d *DepartmentDB) GetDepartmentList(department model.DepartmentReq) (err error, list []model.Department) {\n\tstr := commonSelect(department, d.entity.TableName())\n\tfmt.Println(str)\n\terr = GDB.DB.Select(&list, str)\n\treturn\n}\n\nfunc (d *DepartmentDB) GetDepartmentNameById(id int) (err error, department string) {\n\tstr := fmt.Sprintf(\"select name from %s where id=%d\", d.entity.TableName(), id)\n\terr = GDB.DB.Get(&department, str)\n\treturn\n}\n\nfunc (d *DepartmentDB) GetDepartmentById(id int) (err error, department model.Department) {\n\tstr := fmt.Sprintf(\"select * from %s where id=%d\", d.entity.TableName(), id)\n\terr = GDB.DB.Get(&department, str)\n\treturn\n}\n\nfunc (d *DepartmentDB) AddDepartment(department model.Department, tx ...*sqlx.Tx) error {\n\tdepartment.CreateTime = util.IntToNullInt32(int(time.Now().Unix()))\n\tvar err error\n\tstr := commonInsert(department, d.entity.TableName())\n\tif len(tx) > 0 {\n\t\t_, err = tx[0].NamedExec(str, department)\n\t} else {\n\t\t_, err = GDB.DB.NamedExec(str, department)\n\t}\n\treturn err\n}\n\nfunc (d *DepartmentDB) DeleteDepartmentById(id int, tx ...*sqlx.Tx) (e error, affectedNum int) {\n\tvar err error\n\tvar rows sql.Result\n\tstr := fmt.Sprintf(\"delete from %s where id = %d\", d.entity.TableName(), id)\n\tif len(tx) > 0 {\n\t\trows, err = tx[0].Exec(str)\n\t} else {\n\t\trows, err = GDB.DB.Exec(str)\n\t}\n\taffected, err := rows.RowsAffected()\n\tif err != nil {\n\t\treturn err, 0\n\t}\n\treturn err, int(affected)\n}\n\nfunc (d *DepartmentDB) UpdateDepartment(department model.DepartmentDTO, tx ...*sqlx.Tx) error {\n\tvar err error\n\tstr := commonUpdate(department, d.entity, d.entity.TableName())\n\tif len(tx) > 0 {\n\t\t_, err = tx[0].Exec(str)\n\t} else {\n\t\t_, err = GDB.DB.Exec(str)\n\t}\n\treturn err\n}\n", a.ModName, a.ModName)
	code := strings.ReplaceAll(str, "Department", a.Name)
	fileName := LowFirst(a.Name)
	code = strings.ReplaceAll(code, "department", fileName)
	filePath := fmt.Sprintf("%s/db/%s.go", a.Path, fileName)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Write([]byte(code))
	defer f.Close()
	fmt.Println(filePath, "DB 完成")
}

func (a *Atom) createService() {
	str := "package service\n\nimport (\n\t\"errors\"\n\t\"autocode/db\"\n\t\"autocode/model\"\n\t\"autocode/util\"\n\t\"time\"\n)\n\ntype DepartmentService struct {\n\trepos db.DepartmentDB\n}\n\nfunc (d *DepartmentService) GetDepartmentList(departmentReq model.DepartmentReq) (error error, result interface{}) {\n\terr, list := d.repos.GetDepartmentList(departmentReq)\n\tif err != nil {\n\t\terror = err\n\t\treturn\n\t}\n\tdtoList := make([]model.DepartmentDTO, len(list))\n\tfor i := 0; i < len(list); i++ {\n\t\tutil.Entity2DTO(list[i], &dtoList[i])\n\t}\n\tresult = dtoList\n\treturn\n}\n\nfunc (d *DepartmentService) GetDepartmentById(id int) (err error, result interface{}) {\n\terr, m := d.repos.GetDepartmentById(id)\n\tvar department model.DepartmentDTO\n\tutil.Entity2DTO(m, &department)\n\tresult = department\n\treturn\n}\n\nfunc (d *DepartmentService) AddDepartment(departmentReq model.DepartmentReq) error {\n\tvar department model.Department\n\tutil.DTO2Entity(departmentReq.DepartmentDTO, &department)\n\tif departmentReq.UserId != 0 {\n\t\tdepartment.CreateUserId = util.IntToNullInt32(departmentReq.UserId)\n\t}\n\tdepartment.CreateTime = util.IntToNullInt32(int(time.Now().Unix()))\n\treturn d.repos.AddDepartment(department)\n}\n\nfunc (d *DepartmentService) DeleteDepartmentById(departmentReq model.DepartmentReq) (err error, effected int) {\n\tdto := departmentReq.DepartmentDTO\n\treturn d.repos.DeleteDepartmentById(dto.Id)\n}\n\nfunc (d *DepartmentService) UpdateDepartment(departmentReq model.DepartmentReq) error {\n\tdto := departmentReq.DepartmentDTO\n\tif dto.Id == 0 {\n\t\treturn errors.New(\"非法的数据更新\")\n\t}\n\treturn d.repos.UpdateDepartment(dto)\n}\n"
	fmt.Sprintf("package service\n\nimport (\n\t\"%s/db\"\n\t\"%s/model\"\n\t\"%s/model/commonError\"\n\t\"%s/util\"\n)\n\ntype DepartmentService struct {\n\trepos db.DepartmentDB\n}\n\nfunc (d *DepartmentService) GetDepartmentList(departmentReq model.DepartmentReq) (error error, result interface{}) {\n\terr, list := d.repos.GetDepartmentList(departmentReq)\n\tif err != nil {\n\t\terror = err\n\t\treturn\n\t}\n\tdtoList := make([]model.DepartmentDTO, len(list))\n\tfor i := 0; i < len(list); i++ {\n\t\tutil.Entity2DTO(list[i], &dtoList[i])\n\t}\n\tresult = dtoList\n\treturn\n}\n\nfunc (d *DepartmentService) GetDepartmentById(id int) (err error, result interface{}) {\n\terr, m := d.repos.GetDepartmentById(id)\n\tvar department model.DepartmentDTO\n\tutil.Entity2DTO(m, &department)\n\tresult = department\n\treturn\n}\n\nfunc (d *DepartmentService) AddDepartment(departmentReq model.DepartmentReq) error {\n\tvar department model.Department\n\tutil.DTO2Entity(departmentReq.DepartmentDTO, &department)\n\tif departmentReq.UserId != 0 {\n\t\tdepartment.CreateUserId = util.IntToNullInt32(departmentReq.UserId)\n\t}\n\treturn d.repos.AddDepartment(department)\n}\n\nfunc (d *DepartmentService) DeleteDepartmentById(departmentReq model.DepartmentReq) (err error, effected int) {\n\tdto := departmentReq.DepartmentDTO\n\treturn d.repos.DeleteDepartmentById(dto.Id)\n}\n\nfunc (d *DepartmentService) UpdateDepartment(departmentReq model.DepartmentReq) error {\n\tdto := departmentReq.DepartmentDTO\n\tif dto.Id == 0 {\n\t\treturn commonError.InValidUpdateError\n\t}\n\treturn d.repos.UpdateDepartment(dto)\n}\n", a.ModName, a.ModName, a.ModName, a.ModName)
	code := strings.ReplaceAll(str, "Department", a.Name)
	fileName := LowFirst(a.Name)
	code = strings.ReplaceAll(code, "department", fileName)
	filePath := fmt.Sprintf("%s/service/%s.go", a.Path, fileName)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Write([]byte(code))
	defer f.Close()
	fmt.Println(filePath, "Service 完成")
}

func (a *Atom) createApi() {
	str := fmt.Sprintf("package api\n\nimport (\n\t\"%s/model\"\n\t\"%s/model/commonError\"\n\t\"%s/model/response\"\n\t\"%s/service\"\n\t\"%s/util\"\n\t\"fmt\"\n\t\"github.com/gin-gonic/gin\"\n\t\"log\"\n\t\"strconv\"\n)\n\ntype DepartmentApi struct {\n\tserv service.DepartmentService\n}\n\nfunc (d *DepartmentApi) GetDepartmentList(c *gin.Context) {\n\tvar json model.DepartmentReq\n\tif err := c.ShouldBindJSON(&json); err != nil {\n\t\tlog.Println(err)\n\t\tresponse.FailByJsonError(c)\n\t\treturn\n\t}\n\terr, list := d.serv.GetDepartmentList(json)\n\tif err != nil {\n\t\tif commonError.Contain(err) {\n\t\t\tresponse.FailWithMessage(err.Error(), c)\n\t\t\treturn\n\t\t} else {\n\t\t\tlog.Println(err)\n\t\t\tresponse.Fail(c)\n\t\t\treturn\n\t\t}\n\t} else {\n\t\tjson.Page = 0\n\t\t_, allList := d.serv.GetDepartmentList(json)\n\t\tresponse.OkWithListData(util.GetLength(allList), list, c)\n\t}\n}\n\nfunc (d *DepartmentApi) GetDepartmentById(c *gin.Context) {\n\tids := c.Param(\"id\")\n\tid, _ := strconv.Atoi(ids)\n\terr, department := d.serv.GetDepartmentById(id)\n\tif err != nil {\n\t\tif commonError.Contain(err) {\n\t\t\tresponse.FailWithMessage(err.Error(), c)\n\t\t\treturn\n\t\t} else {\n\t\t\tlog.Println(err)\n\t\t\tresponse.Fail(c)\n\t\t\treturn\n\t\t}\n\t} else {\n\t\tresponse.OkWithData(department, c)\n\t}\n}\n\nfunc (d *DepartmentApi) AddDepartment(c *gin.Context) {\n\tvar json model.DepartmentReq\n\tif err := c.ShouldBindJSON(&json); err != nil {\n\t\tlog.Println(err)\n\t\tresponse.FailByJsonError(c)\n\t\treturn\n\t}\n\terr := d.serv.AddDepartment(json)\n\tif err != nil {\n\t\tif commonError.Contain(err) {\n\t\t\tresponse.FailWithMessage(err.Error(), c)\n\t\t\treturn\n\t\t} else {\n\t\t\tlog.Println(err)\n\t\t\tresponse.Fail(c)\n\t\t\treturn\n\t\t}\n\t} else {\n\t\tresponse.Ok(c)\n\t}\n}\n\nfunc (d *DepartmentApi) AddDepartmentForm(c *gin.Context) {\n\tvar form model.DepartmentReq\n\tMForm, err := c.MultipartForm()\n\tif err != nil {\n\t\tresponse.FailByFormError(c)\n\t\treturn\n\t}\n\tdecoder.SetAliasTag(\"json\")\n\terr = decoder.Decode(&form, MForm.Value)\n\tif err != nil {\n\t\tresponse.FailByFormError(c)\n\t\treturn\n\t}\n\terr = d.serv.AddDepartment(form)\n\tif err != nil {\n\t\tif commonError.Contain(err) {\n\t\t\tresponse.FailWithMessage(err.Error(), c)\n\t\t\treturn\n\t\t} else {\n\t\t\tfmt.Println(err)\n\t\t\tresponse.Fail(c)\n\t\t\treturn\n\t\t}\n\t} else {\n\t\tresponse.Ok(c)\n\t}\n}\n\nfunc (d *DepartmentApi) DeleteDepartment(c *gin.Context) {\n\tvar json model.DepartmentReq\n\tif err := c.ShouldBindJSON(&json); err != nil {\n\t\tlog.Println(err)\n\t\tresponse.FailByJsonError(c)\n\t\treturn\n\t}\n\terr, effectNum := d.serv.DeleteDepartmentById(json)\n\tif err != nil || effectNum == 0 {\n\t\tif commonError.Contain(err) {\n\t\t\tresponse.FailWithMessage(err.Error(), c)\n\t\t\treturn\n\t\t} else {\n\t\t\tlog.Println(err)\n\t\t\tresponse.Fail(c)\n\t\t\treturn\n\t\t}\n\t} else {\n\t\tresponse.Ok(c)\n\t}\n}\n\nfunc (d *DepartmentApi) UpdateDepartment(c *gin.Context) {\n\tvar json model.DepartmentReq\n\tif err := c.ShouldBindJSON(&json); err != nil {\n\t\tlog.Println(err)\n\t\tresponse.FailByJsonError(c)\n\t\treturn\n\t}\n\terr := d.serv.UpdateDepartment(json)\n\tif err != nil {\n\t\tif commonError.Contain(err) {\n\t\t\tresponse.FailWithMessage(err.Error(), c)\n\t\t\treturn\n\t\t} else {\n\t\t\tlog.Println(err)\n\t\t\tresponse.Fail(c)\n\t\t\treturn\n\t\t}\n\t} else {\n\t\tresponse.Ok(c)\n\t}\n}\n", a.ModName, a.ModName, a.ModName, a.ModName, a.ModName)
	code := strings.ReplaceAll(str, "Department", a.Name)
	fileName := LowFirst(a.Name)
	code = strings.ReplaceAll(code, "department", fileName)
	filePath := fmt.Sprintf("%s/api/%s.go", a.Path, fileName)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Write([]byte(code))
	defer f.Close()
	fmt.Println(filePath, "api 完成")
}

func (a *Atom) createRouter() {
	str := fmt.Sprintf("package router\n\nimport (\n\t\"%s/api\"\n\t\"%s/router/middleWare\"\n\t\"github.com/gin-gonic/gin\"\n)\n\ntype DepartmentRouter struct {\n\twebApi api.DepartmentApi\n}\n\nfunc (d *DepartmentRouter) InitDepartmentRouter(g *gin.Engine) {\n\tdeRouter := g.Group(\"department\")\n\t{\n\t\tdeRouter.POST(\"list\", d.webApi.GetDepartmentList)\n\t\tdeRouter.GET(\":id\", middleWare.IdFilter, d.webApi.GetDepartmentById)\n\t\tdeRouter.POST(\"add\", d.webApi.AddDepartment)\n\t\tdeRouter.POST(\"delete\", d.webApi.DeleteDepartment)\n\t\tdeRouter.POST(\"update\", d.webApi.UpdateDepartment)\n\t}\n}\n", a.ModName, a.ModName)
	code := strings.ReplaceAll(str, "Department", a.Name)
	fileName := LowFirst(a.Name)
	code = strings.ReplaceAll(code, "department", fileName)
	filePath := fmt.Sprintf("%s/router/%s.go", a.Path, fileName)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Write([]byte(code))
	defer f.Close()
	fmt.Println(filePath, "router 完成")
}

func (a *Atom) Clear() {
	fileName := LowFirst(a.Name)
	pathArray := []string{"model", "db", "service", "api", "router"}
	for i := 0; i < len(pathArray); i++ {
		filePath := a.Path + "/" + pathArray[i] + "/" + fileName + ".go"
		if flag, _ := PathExists(filePath); flag {
			err := os.Remove(filePath)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(filePath, "清空成功")
		}
	}
	fileNameS := strings.ToLower(a.Name)
	pathArrayS := []string{"model", "db", "service", "api", "router"}
	for i := 0; i < len(pathArray); i++ {
		filePath := a.Path + "/" + pathArrayS[i] + "/" + fileNameS + ".go"
		if flag, _ := PathExists(filePath); flag {
			err := os.Remove(filePath)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(filePath, "清空成功")
		}
	}
}

func (a *Atom) Mkdir() {
	pathArray := []string{"model", "db", "service", "api", "router"}
	for i := 0; i < len(pathArray); i++ {
		filePath := a.Path + "/" + pathArray[i] + "/"
		if flag, _ := PathExists(filePath); !flag {
			err := os.Mkdir(filePath, 777)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(filePath, "文件夹创建成功")
		}
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	//当为空文件或文件夹存在
	if err == nil {
		return true, nil
	}
	//os.IsNotExist(err)为true，文件或文件夹不存在
	if os.IsNotExist(err) {
		return false, nil
	}
	//其它类型，不确定是否存在
	return false, err
}

func (a *Atom) GeneralAutoCode() {
	a.Clear()
	a.Mkdir()
	a.CreateModel()
	a.createDB()
	a.createService()
	a.createApi()
	a.createRouter()
}

func (a *Atom) CreateApiFile() {
	str := fmt.Sprintf("select column_name,column_comment,data_type from information_schema.columns where table_name='%s' and table_schema='%s' ORDER BY ORDINAL_POSITION", a.TblName, db.Name)
	var list []Model
	// err := db.GDB.DB.Select(&list, str)
	query, err := db.GDB.DB.Query(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	for query.Next() {
		m := Model{}
		query.Scan(&m.columnName, &m.columnComment, &m.dataType)
		list = append(list, m)
	}
	// "id": 1,                            // 合同订单id
	for i := 0; i < len(list); i++ {
		m := list[i]
		switch m.dataType {
		case "int", "tinyint", "decimal":
			fmt.Println("\"" + m.columnName + "\":1," + " // " + m.columnComment)
		case "varchar":
			fmt.Println("\"" + m.columnName + "\":\"\"," + " // " + m.columnComment)
		}
	}
	fmt.Println(a.Name, "API 完成")
}
