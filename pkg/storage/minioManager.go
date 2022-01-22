package storage

import (
	"facegram_file_server/config"
	"github.com/minio/minio-go"
)

var minioConnection *minio.Client

func GetStorageConnection() (*minio.Client, error) {

	if minioConnection != nil {
		return minioConnection, nil
	}

	cfg := config.GetStorageServerConfig()

	minioClient, err := minio.New(cfg.Endpoint, cfg.AccessKey, cfg.SecretKey, cfg.UseSSL)

	if err != nil {
		return nil, err
	}

	minioConnection = minioClient
	return minioConnection, nil
}
