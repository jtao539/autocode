package autocode

import (
	"github.com/jtao539/autocode/atom"
	"github.com/jtao539/autocode/util"
	"github.com/jtao539/autocode/util/database"
	"os/exec"
)

const Version = "v0.2.30"

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
	p.a = atom.Atom{Name: p.Name, TblName: p.TblName, Path: p.Path, ModName: p.ModName}
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

func checkProBasic(p *ProBasic) bool {
	return util.CheckStringNULL(p.Path, p.ModName, p.TblName, p.Name)
}

func InitDB(userName, password, host, port, name string) {
	database.Init(userName, password, host, port, name)
}
