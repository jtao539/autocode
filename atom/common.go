package atom

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func (a *Atom) CreateError() {
	commonPath := a.Path + "/" + "common"
	filePath := fmt.Sprintf("%s/definiteError/%s.go", commonPath, "common")
	if flag, _ := PathExists(filePath); flag {
		return
	}
	tempFilePath := fmt.Sprintf("%s/common/definiteError/common.go", template)
	var str string
	if bytes, err := ioutil.ReadFile(tempFilePath); err != nil {
		log.Fatal("Failed to read file: " + tempFilePath)
	} else {
		str = string(bytes)
	}
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Write([]byte(str))
	defer f.Close()
	fmt.Println(filePath, "definiteError 完成")
}

func (a *Atom) CreateRequest() {
	commonPath := a.Path + "/" + "common"
	filePath := fmt.Sprintf("%s/request/%s.go", commonPath, "common")
	if flag, _ := PathExists(filePath); flag {
		return
	}
	tempFilePath := fmt.Sprintf("%s/common/request/common.go", template)
	var str string
	if bytes, err := ioutil.ReadFile(tempFilePath); err != nil {
		log.Fatal("Failed to read file: " + tempFilePath)
	} else {
		str = string(bytes)
	}
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
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
	tempFilePath := fmt.Sprintf("%s/common/response/common.go", template)
	var str string
	if bytes, err := ioutil.ReadFile(tempFilePath); err != nil {
		log.Fatal("Failed to read file: " + tempFilePath)
	} else {
		str = string(bytes)
	}
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
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
	tempFilePath := fmt.Sprintf("%s/common/response/response.go", template)
	var str string
	if bytes, err := ioutil.ReadFile(tempFilePath); err != nil {
		log.Fatal("Failed to read file: " + tempFilePath)
	} else {
		str = string(bytes)
	}
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Write([]byte(str))
	defer f.Close()
	fmt.Println(filePath, "response 完成")
}
