struct DelayTaskRequest {
    1:i64 id,
    2:i64 end_time,
}

service DelayTaskService {
    void delayTask(1:DelayTaskRequest req),
}
