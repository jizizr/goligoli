namespace go api

include "base.thrift"

// User
struct RegisterRequest {
    1:string username,
    2:string password,
}

struct RegisterResponse {
    1:base.BaseResponse base_resp,
    2:string token,
}

struct LoginRequest {
    1:string username,
    2:string password,
}

struct LoginResponse {
    1:base.BaseResponse base_resp,
    2:string token,
}

// Bullet

struct AddBulletRequest {
    1:string token,
    2:i64 live_id,
    3:i64 live_time,
    4:i64 send_time,
    5:string content,
}

struct AddBulletResponse {
    1:base.BaseResponse base_resp,
    2:i64 bullet_id,
}

struct GetHistoryBulletsRequest {
    1:i64 live_id,
    2:i64 bullet_id,
}

struct GetBulletResponse {
    1:base.BaseResponse base_resp,
    2:list<base.Bullet> bullets,
}

service ApiService {
    RegisterResponse Register(1:RegisterRequest req)(api.post="/register"),
    LoginResponse Login(1:LoginRequest req)(api.post="/login"),
    AddBulletResponse SendBullet(1:AddBulletRequest req)(api.post="/bullet/live"),
    GetBulletResponse GetHistoryBullets(1:GetHistoryBulletsRequest req)(api.get="/bullet/history"),
}
