package respcode

const (
	// 全局响应码
	Success = 0
	Error   = -1

	// 用户操作相关响应码
	USER_NOT_FIND        = 1000 // 没有该用户
	USER_LOCKED          = 1001 // 用户已被锁定状态
	USER_HAVED           = 1002 // 用户已存在
	USER_AUTH_FAIL       = 1003 // 用户认证失败
	USER_CREATE_FINISHED = 1004 // 用户创建完成
	USER_LOCK_FINISHED   = 1005 // 用户锁定完成
	USER_FIND_FINISHED   = 1006 // 用户查询完成
)

var errMsg = map[int]string{
	USER_NOT_FIND:        "没有找到该用户",
	USER_LOCKED:          "用户已被锁定状态",
	USER_HAVED:           "用户已存在",
	USER_AUTH_FAIL:       "用户认证失败",
	USER_CREATE_FINISHED: "用户创建完成",
	USER_LOCK_FINISHED:   "用户锁定完成",
	USER_FIND_FINISHED:   "用户查询完成",
}

func GetRespMsg(code int) string {
	msg, ok := errMsg[code]
	if ok {
		return msg
	}
	return "未知的错误码"
}
