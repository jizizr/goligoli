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

service PushService {
    void PushMessage(1:PushMessageRequest req),
    ReceiveMessageResponse ReceiveMessage(1:ReceiveMessageRequest req) (streaming.mode="bidirectional"),
}



