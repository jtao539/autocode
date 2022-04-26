package definiteError

import "errors"

var Array []error

var ProcessError = errors.New("不符合流程或用户无权限")

var ReqFieldError = errors.New("请求字段不符合逻辑")

var RepeatPhoneError = errors.New("手机号码已存在")

var CodeError = errors.New("唯一编码自动生成失败")

var InValidUpdateError = errors.New("非法的数据更新,更新缺少主键")

var InValidSwitchError = errors.New("非法的启用/禁用")

func init() {
	Array = []error{ProcessError, ReqFieldError, RepeatPhoneError, CodeError, InValidSwitchError, InValidUpdateError}
}

func Contain(err error) bool {
	for i := 0; i < len(Array); i++ {
		if Array[i] == err {
			return true
		}
	}
	return false
}
