package consts

type RespCode int32

const Success RespCode = 0

// service error
const (
	CodeNeedLogin RespCode = iota + 10000
	CodeBadRequest
	CodeTokenExpired
	CodeTokenInvalid
	CodeRegisterFailed
	CodeUserAlreadyExists
	CodeLoginFailed
	CodeUserNotFound
	CodeWrongPassword
	CodeSendBulletFailed
	CodeGetHistoryBulletsFailed
	CodeGetBulletByIDFailed
)

var codeMsgMap = map[RespCode]string{
	Success:                     "success",
	CodeNeedLogin:               "need login",
	CodeBadRequest:              "bad request",
	CodeTokenExpired:            "token expired",
	CodeTokenInvalid:            "token invalid",
	CodeRegisterFailed:          "register failed",
	CodeUserAlreadyExists:       "user already exists",
	CodeLoginFailed:             "login failed",
	CodeUserNotFound:            "user not found",
	CodeWrongPassword:           "wrong password",
	CodeSendBulletFailed:        "send bullet failed",
	CodeGetHistoryBulletsFailed: "get history bullets failed",
	CodeGetBulletByIDFailed:     "get bullet by id failed",
}

func (code RespCode) Msg() string {
	return codeMsgMap[code]
}
