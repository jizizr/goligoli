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

// Bullet

struct AddBulletRequest {
    1:required i64 live_id,
    2:required i64 live_time,
    3:required string content,
}

struct AddBulletResponse {
    1:base.BaseResponse base_resp,
    2:i64 bullet_id,
}

struct GetBulletByIDRequest {
    1:required i64 bullet_id,
}

struct GetBulletByIDResponse {
    1:base.BaseResponse base_resp,
    2:base.Bullet bullet,
}

struct GetHistoryBulletsRequest {
    1:required i64 live_id,
    2:required i64 start_time,
    3:required i64 offset,
}

struct GetHistoryBulletsResponse {
    1:base.BaseResponse base_resp,
    2:list<base.Bullet> bullets,
}

service ApiService {
    RegisterResponse Register(1:RegisterRequest req)(api.post="/register"),
    LoginResponse Login(1:LoginRequest req)(api.post="/login"),
    AddBulletResponse SendBullet(1:AddBulletRequest req)(api.post="/bullet/live"),
    GetBulletByIDResponse GetBulletByID(1:GetBulletByIDRequest req)(api.get="/bullet/history/single/"),
    GetHistoryBulletsResponse GetHistoryBullets(1:GetHistoryBulletsRequest req)(api.get="/bullet/history/multi/"),
}
