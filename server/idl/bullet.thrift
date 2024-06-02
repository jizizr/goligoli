namespace go bullet

include "base.thrift"

struct AddBulletRequest {
    1:base.Bullet bullet,
}

struct GetBulletRequest {
    1:i64 bullet_id,
}

struct GetBulletResponse {
    1:optional base.Bullet bullet,
}

struct GetHistoryBulletsRequest {
    1:i64 live_id,
    2:i64 start_time,
    3:i64 offset,
}

struct GetHistoryBulletsResponse {
    1:list<base.Bullet> bullets,
}

service BulletService {
    void AddBullet(1:AddBulletRequest req),
    GetBulletResponse GetBullet(1:GetBulletRequest req),
    GetHistoryBulletsResponse GetHistoryBullets(1:GetHistoryBulletsRequest req),
}
