package setup

import (
	"TinyTikTok/conf"
	"context"
	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog/log"
)

var MinioClient *minio.Client
var Mctx context.Context

func Minio() {
	Mctx = context.Background()
	// Initialize minio client object.
	MinioClient, _ = minio.New(conf.Conf.MinioEndpoint, &minio.Options{
		Creds: credentials.NewStaticV4(conf.Conf.MinioAccessKeyID, conf.Conf.MinioSecretAccessKey, ""),
	})
	log.Info().Msgf("%#v\n", MinioClient)
}
