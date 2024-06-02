namespace go push

include "base.thrift"

struct PushBulletRequest {
    1:base.Bullet bullet,
}

service PushService {
    void PushBullet(1:PushBulletRequest req),
}



