package request

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jtao539/autocode/template/common/response"
	"go.uber.org/zap"
)

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page"`           // 页码
	PageSize int `json:"page_size" form:"page_size"` // 每页大小
}

func BindJson(obj interface{}, c *gin.Context) error {
	err := c.ShouldBindJSON(obj)
	if err != nil {
		fmt.Println(err)
		zap.L().Error(err.Error())
		response.FailByJsonError(c)
		return err
	}
	return err
}

func BindForm(form interface{}, c *gin.Context) error {
	MForm, err := c.MultipartForm()
	if err != nil {
		response.FailByFormError(c)
		return err
	}
	err = response.Decoder.Decode(form, MForm.Value)
	if err != nil {
		response.FailByFormError(c)
		return err
	}
	return err
}
