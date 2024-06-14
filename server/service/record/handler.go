package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	record "github.com/jizizr/goligoli/server/kitex_gen/record"
	"github.com/jizizr/goligoli/server/service/record/config"
	"github.com/minio/minio-go/v7"
	"log"
	"net/http"
	"sync"
	"time"
)

// RecordServiceImpl implements the last service interface defined in the IDL.
type RecordServiceImpl struct {
	m *minio.Client
}

var liveMap sync.Map

// StartRecord implements the RecordServiceImpl interface.
func (s *RecordServiceImpl) StartRecord(ctx context.Context, req *record.StartRecordRecordRequest) (err error) {
	liveMap.Store(req.LiveId, nil)

	go func() {
		for ok := true; ok; _, ok = liveMap.Load(req.LiveId) {
			// 发起HTTP GET请求以读取FLV流
			resp, err := http.Get(fmt.Sprintf("%s/%d.flv", config.GlobalServerConfig.Stream.Address, req.LiveId))
			if err != nil {
				klog.Errorf("Failed to get FLV stream: %v", err)
				time.Sleep(1 * time.Second)
				continue
			}
			defer resp.Body.Close()

			// 检查HTTP响应状态码
			if resp.StatusCode != http.StatusOK {
				time.Sleep(1 * time.Second)
				continue
			}

			// 上传FLV流到MinIO
			uploadInfo, err := s.m.PutObject(ctx, config.GlobalServerConfig.MinioInfo.Bucket, fmt.Sprintf("%d-%d", req.LiveId, time.Now().Unix()), resp.Body, -1, minio.PutObjectOptions{ContentType: "video/x-flv"})
			if err != nil {
				log.Fatalf("Failed to upload stream: %v", err)
			}
			klog.Infof("Successfully uploaded stream: %v", uploadInfo)
		}
	}()
	return
}

// StopRecord implements the RecordServiceImpl interface.
func (s *RecordServiceImpl) StopRecord(ctx context.Context, req *record.StopRecordRecordRequest) (err error) {
	liveMap.Delete(req.LiveId)
	return
}
