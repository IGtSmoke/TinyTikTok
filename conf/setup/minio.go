package setup

import (
	"TinyTikTok/conf"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog/log"
)

// MinioClient Export global Minio client
var MinioClient *minio.Client

// Mctx Export global minio client context
var Mctx context.Context

// Minio Initialize Minio client
func Minio() {
	Mctx = context.Background()
	// Initialize minio client object.
	MinioClient, _ = minio.New(conf.Conf.MinioEndpoint, &minio.Options{
		Creds: credentials.NewStaticV4(conf.Conf.MinioAccessKeyID, conf.Conf.MinioSecretAccessKey, ""),
	})
	log.Info().Msgf("%#v\n", MinioClient)
}
