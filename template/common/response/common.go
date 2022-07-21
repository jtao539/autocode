package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/schema"
	"github.com/jtao539/autocode/template/common/commonError"
	"go.uber.org/zap"
)

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

var Decoder = schema.NewDecoder()

func init() {
	Decoder.SetAliasTag("json")
}

// Pack data 中存储要返回的数据
func Pack(err error, c *gin.Context, data ...interface{}) {
	if err != nil {
		if commonError.Contain(&err) {
			FailWithMessage(err.Error(), c)
			return
		} else {
			fmt.Println(err)
			zap.L().Error(err.Error())
			Fail(c)
			return
		}
	} else {
		if len(data) == 0 {
			Ok(c)
		} else if len(data) > 0 {
			OkWithData(data[0], c)
		}
	}
}
