namespace go api

include "base.thrift"

// User
struct RegisterRequest {
    1:required string username (api.vd="len($)>3 && len($)<32"),
    2:required string password,
}

struct RegisterResponse {
    1:base.BaseResponse base_resp,
    2:string token,
}

struct LoginRequest {
    1:required string username (api.vd="len($)>3 && len($)<32"),
    2:required string password,
}

struct LoginResponse {
    1:base.BaseResponse base_resp,
    2:string token,
}

// Message

struct AddMessageRequest {
    1:required i64 live_id,
    2:required i8  type,
    3:required string content,
}

struct AddMessageResponse {
    1:base.BaseResponse base_resp,
    2:i64 id,
}

struct GetMessageByIDRequest {
    1:required i64 id,
}

struct GetMessageByIDResponse {
    1:base.BaseResponse base_resp,
    2:base.LiveMessage message,
}

struct GetHistoryMessagesRequest {
    1:required i64 live_id,
    2:required i64 start_time,
    3:required i64 offset,
}

struct GetHistoryMessagesResponse {
    1:base.BaseResponse base_resp,
    2:list<base.LiveMessage> messages,
}

struct GetMessageRTRequest {
    1:i64 live_id,
}

struct GetMessageRTResponse {
    1:base.BaseResponse base_resp,
}

//struct BroadcastMessageRequest {
//    1:i64 live_id
//    2:base.LiveMessage message
//}
//
//struct BroadcastMessageResponse{
//    1:base.BaseResponse base_resp
//}

// Live
struct CreateLiveRequest {
    1:required string room_name,
    2:required string description,
}

struct CreateLiveResponse {
    1:base.BaseResponse base_resp,
    2:i64 live_id,
}

struct DeleteLiveRequest {
    1:required i64 live_id,
}

struct DeleteLiveResponse {
    1:base.BaseResponse base_resp,
}

// Lottery
struct PublishLotteryRequest {
    1:required base.Gift gift,
}

struct PublishLotteryResponse {
    1:base.BaseResponse base_resp,
    2:i64 id,
}

struct GetLotteryByIDRequest {
    1:required i64 id,
}

struct GetLotteryByIDResponse {
    1:base.BaseResponse base_resp,
    2:base.Gift gift,
}

struct GetLiveRoomLotteryRequest {
    1:required i64 live_id,
}

struct GetLiveRoomLotteryResponse {
    1:base.BaseResponse base_resp,
    2:list<base.Gift> gifts,
}

struct JoinLotteryRequest {
    1:required i64 lottery_id,
}

struct JoinLotteryResponse {
    1:base.BaseResponse base_resp,
}

service ApiService {
    RegisterResponse Register(1:RegisterRequest req)(api.post="/register"),
    LoginResponse Login(1:LoginRequest req)(api.post="/login"),
    AddMessageResponse SendMessage(1:AddMessageRequest req)(api.post="/message/live"),
    GetMessageByIDResponse GetMessageByID(1:GetMessageByIDRequest req)(api.get="/message/history/single"),
    GetHistoryMessagesResponse GetHistoryMessages(1:GetHistoryMessagesRequest req)(api.get="/message/history/multi"),
    GetMessageRTResponse GetMessageRT(1:GetMessageRTRequest req) (api.get="/message/live"),
    CreateLiveResponse CreateLive(1:CreateLiveRequest req)(api.post="/room/live"),
    DeleteLiveResponse DeleteLive(1:DeleteLiveRequest req)(api.delete="/room/live"),
    PublishLotteryResponse PublishLottery(1:PublishLotteryRequest req)(api.post="/lottery"),
    GetLotteryByIDResponse GetLotteryByID(1:GetLotteryByIDRequest req)(api.get="/lottery/single"),
    GetLiveRoomLotteryResponse GetLiveRoomLottery(1:GetLiveRoomLotteryRequest req)(api.get="/lottery/multi"),
    JoinLotteryResponse JoinLottery(1:JoinLotteryRequest req)(api.post="/lottery/entry"),
}
