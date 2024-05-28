namespace go bullet

include "base.thrift"

struct AddBulletRequest {
    1:base.BaseResponse base_resp,
    2:i64 user_id,
    3:i64 live_id,
    4:i64 live_time,
    5:i64 send_time,
    6:string content,
}

struct AddBulletResponse {
    1:base.BaseResponse base_resp,
    2:i64 bullet_id,
}

struct GetBulletRequest {
    1:base.BaseResponse base_resp,
    2:i64 live_id,
    3:i64 start_time,
    4:i64 offset,
}

struct GetBulletResponse {
    1:base.BaseResponse base_resp,
    2:list<base.Bullet> bullets,
}

service BulletService {
    AddBulletResponse AddBullet(1:AddBulletRequest req),
    GetBulletResponse GetBullet(1:GetBulletRequest req),
}
