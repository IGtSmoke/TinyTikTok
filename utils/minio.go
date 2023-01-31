package utils

import (
	"TinyTikTok/conf/setup"
	"github.com/minio/minio-go/v7"
	"github.com/rs/zerolog/log"
	"io"
	"net/url"
	"time"
)

// UploadFile 上传文件（提供reader）至 minio
func UploadFile(bucketName string, objectName string, reader io.Reader, objectsize int64) error {
	n, err := setup.MinioClient.PutObject(setup.Mctx, bucketName, objectName, reader, objectsize, minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	if err != nil {
		log.Error().Msgf("upload %s of size %d failed, %s", bucketName, objectsize, err)
		return err
	}
	log.Info().Msgf("upload %s of bytes %d successfully", objectName, n.Size)
	return nil
}

// GetFileUrl 从 minio 获取文件Url
func GetFileUrl(bucketName string, fileName string, expires time.Duration) (*url.URL, error) {
	reqParams := make(url.Values)
	if expires <= 0 {
		expires = time.Second * 60 * 60 * 24
	}
	preSignedUrl, err := setup.MinioClient.PresignedGetObject(setup.Mctx, bucketName, fileName, expires, reqParams)
	if err != nil {
		log.Error().Msgf("get url of file %s from bucket %s failed, %s", fileName, bucketName, err)
		return nil, err
	}
	return preSignedUrl, nil
}
