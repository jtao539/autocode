package autocode

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestDemo(t *testing.T) {
	InitDB("root", "123456", "localhost", "3306", "sale")
	a := ProBasic{Name: "Product", TblName: "tbl_product", Path: ".", ModName: "testimpl"}
	a.Start()
}

func TestChangeSuffix(t *testing.T) {
	dir := `./template`
	// findDir(dir, ".go")
	findDir(dir, ".tpl")
}

// 遍历的文件夹
func findDir(dir, suffix string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	path := dir + `/`
	// 遍历这个文件夹
	for _, f := range files {
		// 判断是不是目录
		if f.IsDir() {
			findDir(path+f.Name(), suffix)
		} else {
			oldPath := path + f.Name()
			newPath := path + f.Name()[:strings.LastIndex(f.Name(), ".")] + suffix
			err := os.Rename(oldPath, newPath)
			if err != nil {
				println(err)
				break
			}
		}
	}
}
