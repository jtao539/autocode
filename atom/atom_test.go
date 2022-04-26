package atom

import (
	"github.com/jtao539/autocode/db"
	"testing"
)

func TestAuto(t *testing.T) {
	db.Init("root", "123456", "localhost", "3306", "sale")
	a := Atom{Name: "Product", TblName: "tbl_product", Path: ".", ModName: "autocode"}
	a.GeneralAutoCode()
}
