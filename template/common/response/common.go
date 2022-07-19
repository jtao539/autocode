package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/schema"
	"github.com/jtao539/autocode/template/common/commonError"
	"go.uber.org/zap"
	"reflect"
	"strconv"
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

func GetLength(a interface{}) int {
	v := reflect.ValueOf(a)
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		return v.Len()
	default:
		panic("not supported")
	}
}

func IdFilter(c *gin.Context) {
	ids := c.Param("id")
	_, err := strconv.Atoi(ids)
	if err == nil {
		c.Next()
		return
	}
	FailByIdValid(c)
	c.Abort()
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
