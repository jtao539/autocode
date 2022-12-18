package request

import (
	"github.com/gin-gonic/gin"
	"github.com/jtao539/autocode/template/common/response"
	"go.uber.org/zap"
)

func BindJson(obj interface{}, c *gin.Context) error {
	err := c.ShouldBindJSON(obj)
	if err != nil {
		zap.L().Error(err.Error())
		response.FailByJsonError(c)
		return err
	}
	return err
}

func BindForm(form interface{}, c *gin.Context) error {
	MForm, err := c.MultipartForm()
	if err != nil {
		zap.L().Error(err.Error())
		response.FailByFormError(c)
		return err
	}
	err = response.Decoder.Decode(form, MForm.Value)
	if err != nil {
		zap.L().Error(err.Error())
		response.FailByFormError(c)
		return err
	}
	return err
}
