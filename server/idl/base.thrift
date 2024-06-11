struct BaseResponse {
    1: i32 status_code,
    2: string status_msg,
}

struct LiveMessage {
     1:i8 type,
     2:i64 id,
     3:i64 user_id,
     4:i64 live_id,
     5:i64 live_time,
     6:i64 send_time,
     7:string content,
 }

struct Room{
    1:i64 live_id (go.tag = "gorm:\"type:bigint;primaryKey;autoIncrement:false\""),
    2:string room_name,
    3:string introduction,
    4:i64 owner,
    5:string cover,
}

struct Gift{
    1:i64 id (go.tag = "gorm:\"primaryKey;autoIncrement:false\"")
    2:required i64 live_id (go.tag = "gorm:\"index\"")
    3:required string gift
    4:required i32 count
    5:required i32 end_time
}
