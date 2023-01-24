package init

import (
	"TinyTikTok/conf"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog/log"
)

var MinioClient *minio.Client
var Mctx context.Context

func Minio() {

	Mctx = context.Background()
	// Initialize minio client object.
	MinioClient, _ = minio.New(conf.MinioEndpoint, &minio.Options{
		Creds: credentials.NewStaticV4(conf.MinioAccessKeyID, conf.MinioSecretAccessKey, ""),
	})
	log.Info().Msgf("%#v\n", MinioClient)
}
