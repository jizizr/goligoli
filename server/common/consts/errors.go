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
	CodeSendMessageFailed
	CodeGetHistoryMessagesFailed
	CodeGetMessageByIDFailed
	CodeCreateLiveRoomFailed
	CodeLiveRoomAlreadyExists
	CodeInternalError
	CodeNoLiveRoomRight
	CodeLotteryNotFound
	CodeLiveRoomNotLive
)

var codeMsgMap = map[RespCode]string{
	Success:                      "success",
	CodeNeedLogin:                "need login",
	CodeBadRequest:               "bad request",
	CodeTokenExpired:             "token expired",
	CodeTokenInvalid:             "token invalid",
	CodeRegisterFailed:           "register failed",
	CodeUserAlreadyExists:        "user already exists",
	CodeLoginFailed:              "login failed",
	CodeUserNotFound:             "user not found",
	CodeWrongPassword:            "wrong password",
	CodeSendMessageFailed:        "send message failed",
	CodeGetHistoryMessagesFailed: "get history messages failed",
	CodeGetMessageByIDFailed:     "get message by id failed",
	CodeCreateLiveRoomFailed:     "create live room failed",
	CodeLiveRoomAlreadyExists:    "live room already exists",
	CodeInternalError:            "internal error",
	CodeNoLiveRoomRight:          "no live room right",
	CodeLotteryNotFound:          "lottery not found",
	CodeLiveRoomNotLive:          "live room not live",
}

func (code RespCode) Msg() string {
	return codeMsgMap[code]
}
