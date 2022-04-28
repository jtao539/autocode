package response

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/schema"
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
