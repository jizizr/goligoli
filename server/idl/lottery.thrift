namespace go lottery

include "base.thrift"

struct SetLotteryRequest {
    1:base.Gift gift
}

struct SetLotteryResponse {
    1:i64 id
}

struct GetLotteryRequest {
    1:i64 id
}

struct GetLotteryResponse {
    1:base.Gift gift
}

struct JoinLotteryRequest {
    1:i64 id
    2:i64 uid
}

struct JoinLotteryResponse {
    1:bool success
}

struct GetLiveRoomLotteryRequest {
    1:i64 live_id
}

struct GetLiveRoomLotteryResponse {
    2:list<base.Gift> gifts
}

struct DrawLotteryRequest {
    1:i64 id
}

struct DrawLotteryResponse {
    1:string msg
}

service LotteryService {
    SetLotteryResponse SetLottery(1:SetLotteryRequest req)
    GetLotteryResponse GetLottery(1:GetLotteryRequest req)
    JoinLotteryResponse JoinLottery(1:JoinLotteryRequest req)
    GetLiveRoomLotteryResponse GetLiveRoomLottery(1:GetLiveRoomLotteryRequest req)
    DrawLotteryResponse DrawLottery(1:DrawLotteryRequest req)
}
