namespace go push

include "base.thrift"

struct PushMessageRequest {
    1:base.LiveMessage message,
}

struct ReceiveMessageRequest {
    1:i64 user_id,
    2:i64 live_id,
}

struct ReceiveMessageResponse {
    1:base.LiveMessage message,
}

struct StopMessageRequest {
    1:i64 live_id,
}

struct InitLiveRoomReciverRequest {
    1:i64 live_id,
}

service PushService {
    void PushMessage(1:PushMessageRequest req),
    void StopMessage(1:StopMessageRequest req),
    void InitLiveRoomReciver(1:InitLiveRoomReciverRequest req),
    ReceiveMessageResponse ReceiveMessage(1:ReceiveMessageRequest req) (streaming.mode="bidirectional"),
}



