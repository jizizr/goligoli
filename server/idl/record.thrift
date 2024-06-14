struct StartRecordRecordRequest {
    1:i64 liveId,
}

struct StopRecordRecordRequest {
    1:i64 liveId,
}

struct GetRecordListRequest {
    1:i64 liveId,
}

service RecordService {
    void StartRecord(1:StartRecordRecordRequest req),
    void StopRecord(1:StopRecordRecordRequest req),
}
