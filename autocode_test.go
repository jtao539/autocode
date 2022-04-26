package autocode

import "testing"

func TestDemo(t *testing.T) {
	InitDB("root", "123456", "localhost", "3306", "sale")
	a := ProBasic{Name: "Product", TblName: "tbl_product", Path: ".", ModName: "testimpl"}
	a.Start()
}
