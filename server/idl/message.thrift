namespace go message

include "base.thrift"

struct AddMessageRequest {
    1:base.LiveMessage message,
}

struct GetMessageRequest {
    1:i64 id,
}

struct GetMessageResponse {
    1:optional base.LiveMessage message,
}

struct GetHistoryMessagesRequest {
    1:i64 live_id,
    2:i64 start_time,
    3:i64 offset,
}

struct GetHistoryMessagesResponse {
    1:list<base.LiveMessage> messages,
}

service MessageService {
    void AddMessage(1:AddMessageRequest req),
    GetMessageResponse GetMessage(1:GetMessageRequest req),
    GetHistoryMessagesResponse GetHistoryMessages(1:GetHistoryMessagesRequest req),
}
