package conf

const (
	DSN                  string = "root:003127@tcp(127.0.0.1:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	SnowflakeId          int64  = 1
	BucketName           string = "videos"
	MinioEndpoint        string = "localhost:9000"
	MinioAccessKeyID     string = "minioadmin"
	MinioSecretAccessKey string = "minioadmin"
	RedisAddr            string = "localhost:49153"
	RedisPassword        string = "redispw"
	RedisDB              int    = 0
)
