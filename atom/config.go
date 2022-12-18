package atom

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func (a *Atom) CreateConfig() {
	m := make(map[string]string)
	m["config.yml"] = "config.yml"
	m["cors.go"] = "cors.tpl"
	m["rbac_models.conf"] = "rbac_models.conf"
	for k, v := range m {
		a.createConfig(k, v)
	}
}

func (a *Atom) createConfig(newName, oldName string) {
	commonPath := a.Path + "/" + "config"
	filePath := fmt.Sprintf("%s/%s", commonPath, newName)
	if flag, _ := PathExists(filePath); flag {
		return
	}
	tempFilePath := fmt.Sprintf("%s/config/%s", template, oldName)
	var str string
	if bytes, err := ioutil.ReadFile(tempFilePath); err != nil {
		log.Fatal("Failed to read file: " + tempFilePath)
	} else {
		str = string(bytes)
		str = strings.ReplaceAll(str, MODName, a.ModName)
	}
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Write([]byte(str))
	defer f.Close()
	fmt.Println(filePath, newName, "完成")
}
