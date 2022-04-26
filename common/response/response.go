package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type ListData struct {
	Total int         `json:"total"`
	List  interface{} `json:"list"`
}

const (
	ParamError  = 1
	NoDataError = 2
	ERROR       = 3
	SUCCESS     = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkWithListData(total int, data interface{}, c *gin.Context) {
	Result(SUCCESS, ListData{Total: total, List: data}, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailByImageUpload(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "图片上传失败", c)
}

func FailByFileUpload(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "文件上传失败", c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

func FailByJsonError(c *gin.Context) {
	Result(ParamError, map[string]interface{}{}, "JSON格式错误或数据类型错误", c)
}

func FailByNoPermission(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "没有权限", c)
}

func FailByFormError(c *gin.Context) {
	Result(ParamError, map[string]interface{}{}, "form-data格式错误或数据类型错误", c)
}

func FailByNoData(c *gin.Context) {
	Result(NoDataError, map[string]interface{}{}, "数据不存在", c)
}

func FailByIdValid(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "路径未找到", c)
}
