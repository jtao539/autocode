package atom

import (
	"fmt"
	"github.com/jtao539/autocode/util/database"
	"go/build"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var template string

const MODName = "github.com/jtao539/autocode/template"

const TPL = ".tpl"

type Atom struct {
	Name    string
	TblName string
	Path    string
	ModName string
	Version string
}

type Model struct {
	columnName    string `db:"column_name"`
	columnComment string `db:"column_comment"`
	dataType      string `db:"data_type"`
}

func (a *Atom) CreateModel() {
	str := fmt.Sprintf("select column_name,column_comment,data_type from information_schema.columns where table_name='%s' and table_schema='%s' ORDER BY ORDINAL_POSITION", a.TblName, database.Name)
	var list []Model
	// err := db.GDB.DB.Select(&list, str)
	query, err := database.GDB.DB.Query(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	for query.Next() {
		m := Model{}
		query.Scan(&m.columnName, &m.columnComment, &m.dataType)
		list = append(list, m)
	}
	code := fmt.Sprintf("package model\n\nimport (\n\t\"%s/common/request\"\n\t\"database/sql\"\n)\n\ntype %s struct {\n", a.ModName, a.Name)
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
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Write([]byte(code))
	defer f.Close()
	fmt.Println(filePath, "Model ??????")
}

func (a *Atom) createDB() {
	global := fmt.Sprintf("%s/db/global.go", a.Path)
	tempGlobal := fmt.Sprintf("%s/db/global%s", template, TPL)
	if flag, _ := PathExists(global); !flag {
		var code string
		if bytes, err := ioutil.ReadFile(tempGlobal); err != nil {
			log.Fatal("Failed to read file: " + tempGlobal)
		} else {
			code = string(bytes)
		}
		f, err := os.OpenFile(global, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			fmt.Println(err)
			return
		}
		f.Write([]byte(code))
		defer f.Close()
	}
	fileName := LowFirst(a.Name)
	filePath := fmt.Sprintf("%s/db/%s.go", a.Path, fileName)
	tempFilePath := fmt.Sprintf("%s/db/department%s", template, TPL)
	var str string
	if bytes, err := ioutil.ReadFile(tempFilePath); err != nil {
		log.Fatal("Failed to read file: " + tempFilePath)
	} else {
		str = string(bytes)
	}
	str = strings.ReplaceAll(str, MODName, a.ModName)
	code := strings.ReplaceAll(str, "Department", a.Name)
	code = strings.ReplaceAll(code, "department", fileName)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Write([]byte(code))
	defer f.Close()
	fmt.Println(filePath, "DB ??????")
}

func (a *Atom) createService() {
	fileName := LowFirst(a.Name)
	filePath := fmt.Sprintf("%s/service/%s.go", a.Path, fileName)
	tempFilePath := fmt.Sprintf("%s/service/department%s", template, TPL)
	var str string
	if bytes, err := ioutil.ReadFile(tempFilePath); err != nil {
		log.Fatal("Failed to read file: " + tempFilePath)
	} else {
		str = string(bytes)
	}
	str = strings.ReplaceAll(str, MODName, a.ModName)
	code := strings.ReplaceAll(str, "Department", a.Name)
	code = strings.ReplaceAll(code, "department", fileName)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Write([]byte(code))
	defer f.Close()
	fmt.Println(filePath, "Service ??????")
}

func (a *Atom) createApi() {
	fileName := LowFirst(a.Name)
	filePath := fmt.Sprintf("%s/api/%s.go", a.Path, fileName)
	tempFilePath := fmt.Sprintf("%s/api/department%s", template, TPL)
	var str string
	if bytes, err := ioutil.ReadFile(tempFilePath); err != nil {
		log.Fatal("Failed to read file: " + tempFilePath)
	} else {
		str = string(bytes)
	}
	str = strings.ReplaceAll(str, MODName, a.ModName)
	code := strings.ReplaceAll(str, "Department", a.Name)
	code = strings.ReplaceAll(code, "department", fileName)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Write([]byte(code))
	defer f.Close()
	fmt.Println(filePath, "api ??????")
}

func (a *Atom) createRouter() {
	register := fmt.Sprintf("%s/router/register.go", a.Path)
	tempRegister := fmt.Sprintf("%s/router/register%s", template, TPL)
	if flag, _ := PathExists(register); !flag {
		var code string
		if bytes, err := ioutil.ReadFile(tempRegister); err != nil {
			log.Fatal("Failed to read file: " + tempRegister)
		} else {
			code = string(bytes)
		}
		f, err := os.OpenFile(register, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			fmt.Println(err)
			return
		}
		f.Write([]byte(code))
		defer f.Close()
	}
	fileName := LowFirst(a.Name)
	filePath := fmt.Sprintf("%s/router/%s.go", a.Path, fileName)
	tempFilePath := fmt.Sprintf("%s/router/department%s", template, TPL)
	var str string
	if bytes, err := ioutil.ReadFile(tempFilePath); err != nil {
		log.Fatal("Failed to read file: " + tempFilePath)
	} else {
		str = string(bytes)
	}
	str = strings.ReplaceAll(str, MODName, a.ModName)
	code := strings.ReplaceAll(str, "Department", a.Name)
	code = strings.ReplaceAll(code, "department", fileName)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Write([]byte(code))
	defer f.Close()
	fmt.Println(filePath, "router ??????")
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
			fmt.Println(filePath, "????????????")
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
			fmt.Println(filePath, "????????????")
		}
	}
}

func (a *Atom) MkSomeDir() {
	pathArray := []string{"model", "db", "service", "api", "router", "common"}
	for i := 0; i < len(pathArray); i++ {
		MkDir(a.Path, pathArray[i])
	}
	innerArray := []string{"commonError", "request", "response"}
	for i := 0; i < len(innerArray); i++ {
		MkDir(a.Path+"/common", innerArray[i])
	}
}

func MkDir(path, fileName string) {
	filePath := path + "/" + fileName + "/"
	if flag, _ := PathExists(filePath); !flag {
		err := os.Mkdir(filePath, 0777)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	//?????????????????????????????????
	if err == nil {
		return true, nil
	}
	//os.IsNotExist(err)???true??????????????????????????????
	if os.IsNotExist(err) {
		return false, nil
	}
	//????????????????????????????????????
	return false, err
}

func (a *Atom) GeneralAutoCode() {
	a.initTemplate()
	a.Clear()
	a.MkSomeDir()
	a.CreateError()
	a.CreateResponse()
	a.CreateRResponse()
	a.CreateRequest()
	a.CreateModel()
	a.createDB()
	a.createService()
	a.createApi()
	a.createRouter()
}

func (a *Atom) CreateApiFile() {
	str := fmt.Sprintf("select column_name,column_comment,data_type from information_schema.columns where table_name='%s' and table_schema='%s' ORDER BY ORDINAL_POSITION", a.TblName, database.Name)
	var list []Model
	// err := db.GDB.DB.Select(&list, str)
	query, err := database.GDB.DB.Query(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	for query.Next() {
		m := Model{}
		query.Scan(&m.columnName, &m.columnComment, &m.dataType)
		list = append(list, m)
	}
	// "id": 1,                            // ????????????id
	for i := 0; i < len(list); i++ {
		m := list[i]
		switch m.dataType {
		case "int", "tinyint", "decimal":
			fmt.Println("\"" + m.columnName + "\":1," + " // " + m.columnComment)
		case "varchar":
			fmt.Println("\"" + m.columnName + "\":\"\"," + " // " + m.columnComment)
		}
	}
	fmt.Println(a.Name, "API ??????")
}

func (a *Atom) initTemplate() {
	template = build.Default.GOPATH + string(filepath.Separator) + "pkg" + string(filepath.Separator) + "mod" + string(filepath.Separator) + "github.com" + string(filepath.Separator) + "jtao539" + string(filepath.Separator) + "autocode@" + a.Version + string(filepath.Separator) + "template"
}
