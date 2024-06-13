namespace go live

include "base.thrift"

struct CreateLiveRoomRequest {
    1:base.Room room;
}

struct CreateLiveRoomResponse {
    1:i64 live_id;
}

struct GetLiveRoomOwnerRequest {
    1:i64 live_id;
}

struct GetLiveRoomOwnerResponse {
    1:i64 owner;
}

struct GetLiveRoomRequest {
    1:i64 live_id;
}

struct GetLiveRoomResponse {
    1:base.Room room;
}

struct StopLiveRoomRequest {
    1:i64 live_id;
}

service LiveService {
    CreateLiveRoomResponse CreateLiveRoom(1:CreateLiveRoomRequest req);
    GetLiveRoomOwnerResponse GetLiveRoomOwner(1:GetLiveRoomOwnerRequest req);
    GetLiveRoomResponse GetLiveRoom(1:GetLiveRoomRequest req);
    void StopLiveRoom(1:StopLiveRoomRequest req);
}
