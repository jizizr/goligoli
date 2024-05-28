struct BaseResponse {
    1: i32 status_code
    2: string status_msg
}

struct Bullet {
     1:i64 bullet_id,
     2:i64 user_id,
     3:i64 live_id,
     4:i64 live_time,
     5:i64 send_time,
     6:string content,
 }
