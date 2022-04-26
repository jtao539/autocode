package atom

import (
	"github.com/jtao539/autocode/util/database"
	"testing"
)

func TestAuto(t *testing.T) {
	database.Init("root", "admin539", "localhost", "3307", "sale")
	a := Atom{Name: "Department", TblName: "tbl_department", Path: "..", ModName: "testimpl"}
	a.GeneralAutoCode()
}
