package initialize

import (
	"github.com/jizizr/goligoli/server/service/record/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func InitMinio() *minio.Client {
	m := config.GlobalServerConfig.MinioInfo
	minioClient, err := minio.New(m.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(m.ID, m.Secret, ""),
		Secure: false,
	})
	if err != nil {
		panic(err)
	}
	return minioClient
}
