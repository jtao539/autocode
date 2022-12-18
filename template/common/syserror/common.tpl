package syserror

import (
	"errors"
	"fmt"
)

const Sign = "SysERR"

var Array []error

var ProcessError = errors.New("不符合流程或用户无权限")

var ReqFieldError = errors.New("请求字段不符合逻辑")

var RepeatPhoneError = errors.New("手机号码已存在")

var EmptyPhoneError = errors.New("手机号码不能为空")

var CodeError = errors.New("唯一编码自动生成失败")

var InValidUpdateError = errors.New("非法的数据更新,更新缺少主键")

var InValidSwitchError = errors.New("非法的状态变更")

var GetRoleError = errors.New("获取角色失败")

func init() {
	Array = []error{
		ProcessError,
		ReqFieldError,
		RepeatPhoneError,
		EmptyPhoneError,
		CodeError,
		InValidUpdateError,
		InValidSwitchError,
		GetRoleError,
	}
}

func Contain(err *error) bool {
	e := *err
	for i := 0; i < len(Array); i++ {
		if Array[i] == e {
			return true
		}
	}
	l := len(e.Error()) - len(Sign)
	if l > len(Sign) && e.Error()[l:] == Sign {
		a := errors.New(e.Error()[:l])
		*err = a
		return true
	}
	return false
}

func WriteError(msg string) error {
	e := errors.New(fmt.Sprintf("%s%s", msg, Sign))
	return e
}
