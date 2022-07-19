package request

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jtao539/autocode/template/common/response"
	"go.uber.org/zap"
)

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page         int `json:"page" form:"page"`           // 页码
	PageSize     int `json:"page_size" form:"page_size"` // 每页大小
	UserId       int `json:"user_id" form:"user_id"`
	Flag         int `json:"flag" form:"flag"`
	CreateUserId int `db:"create_user_id" json:"create_user_id"` // 创建人用户id（tbl_user_id）
}

// GetById Find by id structure
type GetById struct {
	ID int `json:"id" form:"id"` // 主键ID
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId string `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct{}

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
