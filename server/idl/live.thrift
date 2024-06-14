namespace go live

include "base.thrift"

struct CreateLiveRoomRequest {
    1:base.Room room;
}

struct CreateLiveRoomResponse {
    1:i64 live_id;
    2:string key;
}

struct GetLiveRoomOwnerRequest {
    1:i64 live_id;
}

struct GetLiveRoomOwnerResponse {
    1:i64 owner;
}

struct GetLiveRoomStatusRequest {
    1:i64 live_id;
}

struct GetLiveRoomStatusResponse {
    1:bool is_live;
}

struct GetLiveRoomRequest {
    1:i64 live_id;
}

struct GetLiveRoomResponse {
    1:base.Room room;
}

struct GetLiveRoomKeyRequest {
    1:i64 live_id;
}

struct GetLiveRoomKeyResponse {
    1:string key;
}

struct StopLiveRoomRequest {
    1:i64 live_id;
}

struct GetAllOnlineLiveRoomResponse {
    1:list<i64> live_ids;
}

service LiveService {
    CreateLiveRoomResponse CreateLiveRoom(1:CreateLiveRoomRequest req);
    GetLiveRoomOwnerResponse GetLiveRoomOwner(1:GetLiveRoomOwnerRequest req);
    GetLiveRoomStatusResponse GetLiveRoomStatus(1:GetLiveRoomStatusRequest req);
    GetLiveRoomResponse GetLiveRoom(1:GetLiveRoomRequest req);
    GetLiveRoomKeyResponse GetLiveRoomKey(1:GetLiveRoomKeyRequest req);
    GetAllOnlineLiveRoomResponse GetAllOnlineLiveRoom();
    void StopLiveRoom(1:StopLiveRoomRequest req);
}
