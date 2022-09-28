package autocode

import (
	"github.com/jtao539/autocode/atom"
	"github.com/jtao539/autocode/util"
	"github.com/jtao539/autocode/util/database"
	"os/exec"
)

const Version = "v1.0.14"

const Api = "api"

const Router = "router"

const Model = "model"

const Service = "service"

const Db = "db"

const Json = "json"

type ProBasic struct {
	Name    string
	TblName string
	Path    string
	ModName string
	a       atom.Atom
}

func (p *ProBasic) init() {
	if flag := checkProBasic(p); flag {
		panic("基础数据 不能为空")
	}
	if p.Path == "" {
		p.Path = "."
	}
	p.a = atom.Atom{Name: p.Name, TblName: p.TblName, Path: p.Path, ModName: p.ModName, Version: Version}
}

func (p *ProBasic) Start() {
	if err := database.GDB.DB.Ping(); err != nil {
		panic("数据库连接失败!")
	}
	p.init()
	p.a.GeneralAutoCode()
	var cmd *exec.Cmd
	cmd = exec.Command("go fmt")
	cmd.Start()
}

// StartFunc 自定义生成参数可为 Model、APi 等
func (p *ProBasic) StartFunc(args ...string) {
	if err := database.GDB.DB.Ping(); err != nil {
		panic("数据库连接失败!")
	}
	p.init()
	p.a.InitTemplate()
	p.a.Clear()
	p.a.MkSomeDir()
	p.a.CreateError()
	p.a.CreateResponse()
	p.a.CreateRResponse()
	p.a.CreateRequest()
	if containArray(Model, args) {
		p.a.CreateModel()
	} else if containArray(Db, args) {
		p.a.CreateDB()
	} else if containArray(Service, args) {
		p.a.CreateService()
	} else if containArray(Api, args) {
		p.a.CreateApi()
	} else if containArray(Router, args) {
		p.a.CreateRouter()
	} else if containArray(Json, args) {
		p.a.CreateJson()
	}
	var cmd *exec.Cmd
	cmd = exec.Command("go fmt")
	cmd.Start()
}

func checkProBasic(p *ProBasic) bool {
	return util.CheckStringNULL(p.ModName, p.TblName, p.Name)
}

func InitDB(userName, password, host, port, name string) {
	database.Init(userName, password, host, port, name)
}

func containArray(name string, args []string) bool {
	for i := 0; i < len(args); i++ {
		if name == args[i] {
			return true
		}
	}
	return false
}
