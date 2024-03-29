package atom

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func (a *Atom) CreateError() {
	commonPath := a.Path + "/" + "common"
	filePath := fmt.Sprintf("%s/syserror/%s.go", commonPath, "common")
	if flag, _ := PathExists(filePath); flag {
		return
	}
	tempFilePath := fmt.Sprintf("%s/common/syserror/common%s", template, TPL)
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
	fmt.Println(filePath, "syserror 完成")
}

func (a *Atom) CreateRequest() {
	commonPath := a.Path + "/" + "common"
	filePath := fmt.Sprintf("%s/request/%s.go", commonPath, "common")
	if flag, _ := PathExists(filePath); flag {
		return
	}
	tempFilePath := fmt.Sprintf("%s/common/request/common%s", template, TPL)
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
	fmt.Println(filePath, "request 完成")
}

func (a *Atom) CreateResponse() {
	commonPath := a.Path + "/" + "common"
	filePath := fmt.Sprintf("%s/response/%s.go", commonPath, "common")
	if flag, _ := PathExists(filePath); flag {
		return
	}
	tempFilePath := fmt.Sprintf("%s/common/response/common%s", template, TPL)
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
	fmt.Println(filePath, "response common 完成")
}

func (a *Atom) CreateRResponse() {
	commonPath := a.Path + "/" + "common"
	filePath := fmt.Sprintf("%s/response/%s.go", commonPath, "response")
	if flag, _ := PathExists(filePath); flag {
		return
	}
	tempFilePath := fmt.Sprintf("%s/common/response/response%s", template, TPL)
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
	fmt.Println(filePath, "response 完成")
}

func (a *Atom) CreateUtil() {
	commonPath := a.Path + "/" + "common"
	filePath := fmt.Sprintf("%s/util/%s.go", commonPath, "common")
	if flag, _ := PathExists(filePath); flag {
		return
	}
	tempFilePath := fmt.Sprintf("%s/common/util/common%s", template, TPL)
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
	fmt.Println(filePath, "util 完成")
}

func (a *Atom) CreateZLog() {
	commonPath := a.Path + "/" + "common"
	filePath := fmt.Sprintf("%s/zlog/%s.go", commonPath, "common")
	if flag, _ := PathExists(filePath); flag {
		return
	}
	tempFilePath := fmt.Sprintf("%s/common/zlog/common%s", template, TPL)
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
	fmt.Println(filePath, "zlog 完成")
}
