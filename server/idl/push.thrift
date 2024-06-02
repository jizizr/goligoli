namespace go push

include "base.thrift"

struct PushBulletRequest {
    1:base.Bullet bullet,
}

struct ReceiveBulletRequest {
    1:i64 user_id,
    2:i64 live_id,
}

struct ReceiveBulletResponse {
    1:base.Bullet bullet,
}

service PushService {
    void PushBullet(1:PushBulletRequest req),
    ReceiveBulletResponse ReceiveBullet(1:ReceiveBulletRequest req) (streaming.mode="bidirectional"),
}



