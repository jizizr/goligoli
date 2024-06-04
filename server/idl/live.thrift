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

service LiveService {
    CreateLiveRoomResponse CreateLiveRoom(1:CreateLiveRoomRequest req);
    GetLiveRoomOwnerResponse GetLiveRoomOwner(1:GetLiveRoomOwnerRequest req);
}