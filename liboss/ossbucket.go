package liboss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func GetBucket(bucketName string) (*oss.Bucket, error) {

	client, err := oss.New(EndpointPublic, AccessId, AccessKey)
	if err != nil {
		return nil, err
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}
