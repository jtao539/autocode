package atom

import (
	"fmt"
	"os"
)

func (a *Atom) CreateError() {
	commonPath := a.Path + "/" + "common"
	filePath := fmt.Sprintf("%s/definiteError/%s.go", commonPath, "common")
	if flag, _ := PathExists(filePath); flag {
		return
	}
	str := "package definiteError\n\nimport \"errors\"\n\nvar Array []error\n\nvar ProcessError = errors.New(\"不符合流程或用户无权限\")\n\nvar ReqFieldError = errors.New(\"请求字段不符合逻辑\")\n\nvar RepeatPhoneError = errors.New(\"手机号码已存在\")\n\nvar CodeError = errors.New(\"唯一编码自动生成失败\")\n\nvar InValidUpdateError = errors.New(\"非法的数据更新,更新缺少主键\")\n\nvar InValidSwitchError = errors.New(\"非法的启用/禁用\")\n\nfunc init() {\n\tArray = []error{ProcessError, ReqFieldError, RepeatPhoneError, CodeError, InValidSwitchError, InValidUpdateError}\n}\n\nfunc Contain(err error) bool {\n\tfor i := 0; i < len(Array); i++ {\n\t\tif Array[i] == err {\n\t\t\treturn true\n\t\t}\n\t}\n\treturn false\n}\n"
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Write([]byte(str))
	defer f.Close()
	fmt.Println(filePath, "definiteError 完成")
}

func (a *Atom) CreateRequest() {
	commonPath := a.Path + "/" + "common"
	filePath := fmt.Sprintf("%s/request/%s.go", commonPath, "common")
	if flag, _ := PathExists(filePath); flag {
		return
	}
	str := "package request\n\n// PageInfo Paging common input parameter structure\ntype PageInfo struct {\n\tPage         int `json:\"page\" form:\"page\"`           // 页码\n\tPageSize     int `json:\"page_size\" form:\"page_size\"` // 每页大小\n\tUserId       int `json:\"user_id\" form:\"user_id\"`\n\tFlag         int `json:\"flag\" form:\"flag\"`\n\tCreateUserId int `db:\"create_user_id\" json:\"create_user_id\"` // 创建人用户id（tbl_user_id）\n}\n\n// GetById Find by id structure\ntype GetById struct {\n\tID int `json:\"id\" form:\"id\"` // 主键ID\n}\n\nfunc (r *GetById) Uint() uint {\n\treturn uint(r.ID)\n}\n\ntype IdsReq struct {\n\tIds []int `json:\"ids\" form:\"ids\"`\n}\n\n// GetAuthorityId Get role by id structure\ntype GetAuthorityId struct {\n\tAuthorityId string `json:\"authorityId\" form:\"authorityId\"` // 角色ID\n}\n\ntype Empty struct{}\n"
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Write([]byte(str))
	defer f.Close()
	fmt.Println(filePath, "request 完成")
}

func (a *Atom) CreateResponse() {
	commonPath := a.Path + "/" + "common"
	filePath := fmt.Sprintf("%s/response/%s.go", commonPath, "common")
	if flag, _ := PathExists(filePath); flag {
		return
	}
	str := "package response\n\nimport (\n\t\"github.com/gin-gonic/gin\"\n\t\"github.com/gorilla/schema\"\n\t\"reflect\"\n\t\"strconv\"\n)\n\ntype PageResult struct {\n\tList     interface{} `json:\"list\"`\n\tTotal    int64       `json:\"total\"`\n\tPage     int         `json:\"page\"`\n\tPageSize int         `json:\"pageSize\"`\n}\n\nvar Decoder = schema.NewDecoder()\n\nfunc init() {\n\tDecoder.SetAliasTag(\"json\")\n}\n\nfunc GetLength(a interface{}) int {\n\tv := reflect.ValueOf(a)\n\tswitch v.Kind() {\n\tcase reflect.Slice, reflect.Array:\n\t\treturn v.Len()\n\tdefault:\n\t\tpanic(\"not supported\")\n\t}\n}\n\nfunc IdFilter(c *gin.Context) {\n\tids := c.Param(\"id\")\n\t_, err := strconv.Atoi(ids)\n\tif err == nil {\n\t\tc.Next()\n\t\treturn\n\t}\n\tFailByIdValid(c)\n\tc.Abort()\n}\n"
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Write([]byte(str))
	defer f.Close()
	fmt.Println(filePath, "response common 完成")
}

func (a *Atom) CreateRResponse() {
	commonPath := a.Path + "/" + "common"
	filePath := fmt.Sprintf("%s/response/%s.go", commonPath, "response")
	if flag, _ := PathExists(filePath); flag {
		return
	}
	str := "package response\n\nimport (\n\t\"github.com/gin-gonic/gin\"\n\t\"net/http\"\n)\n\ntype Response struct {\n\tCode int         `json:\"code\"`\n\tData interface{} `json:\"data\"`\n\tMsg  string      `json:\"msg\"`\n}\n\ntype ListData struct {\n\tTotal int         `json:\"total\"`\n\tList  interface{} `json:\"list\"`\n}\n\nconst (\n\tParamError  = 1\n\tNoDataError = 2\n\tERROR       = 3\n\tSUCCESS     = 0\n)\n\nfunc Result(code int, data interface{}, msg string, c *gin.Context) {\n\t// 开始时间\n\tc.JSON(http.StatusOK, Response{\n\t\tcode,\n\t\tdata,\n\t\tmsg,\n\t})\n}\n\nfunc Ok(c *gin.Context) {\n\tResult(SUCCESS, map[string]interface{}{}, \"操作成功\", c)\n}\n\nfunc OkWithMessage(message string, c *gin.Context) {\n\tResult(SUCCESS, map[string]interface{}{}, message, c)\n}\n\nfunc OkWithData(data interface{}, c *gin.Context) {\n\tResult(SUCCESS, data, \"操作成功\", c)\n}\n\nfunc OkWithListData(total int, data interface{}, c *gin.Context) {\n\tResult(SUCCESS, ListData{Total: total, List: data}, \"操作成功\", c)\n}\n\nfunc OkWithDetailed(data interface{}, message string, c *gin.Context) {\n\tResult(SUCCESS, data, message, c)\n}\n\nfunc Fail(c *gin.Context) {\n\tResult(ERROR, map[string]interface{}{}, \"操作失败\", c)\n}\n\nfunc FailWithMessage(message string, c *gin.Context) {\n\tResult(ERROR, map[string]interface{}{}, message, c)\n}\n\nfunc FailByImageUpload(c *gin.Context) {\n\tResult(ERROR, map[string]interface{}{}, \"图片上传失败\", c)\n}\n\nfunc FailByFileUpload(c *gin.Context) {\n\tResult(ERROR, map[string]interface{}{}, \"文件上传失败\", c)\n}\n\nfunc FailWithDetailed(data interface{}, message string, c *gin.Context) {\n\tResult(ERROR, data, message, c)\n}\n\nfunc FailByJsonError(c *gin.Context) {\n\tResult(ParamError, map[string]interface{}{}, \"JSON格式错误或数据类型错误\", c)\n}\n\nfunc FailByNoPermission(c *gin.Context) {\n\tResult(ERROR, map[string]interface{}{}, \"没有权限\", c)\n}\n\nfunc FailByFormError(c *gin.Context) {\n\tResult(ParamError, map[string]interface{}{}, \"form-data格式错误或数据类型错误\", c)\n}\n\nfunc FailByNoData(c *gin.Context) {\n\tResult(NoDataError, map[string]interface{}{}, \"数据不存在\", c)\n}\n\nfunc FailByIdValid(c *gin.Context) {\n\tResult(ERROR, map[string]interface{}{}, \"路径未找到\", c)\n}\n"
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Write([]byte(str))
	defer f.Close()
	fmt.Println(filePath, "response 完成")
}
