package liboss

import (
	"fmt"
	"strings"

	"github.com/beego/beego/v2/core/logs"
)

func PutString(key string, val string) error {

	logs.Info(fmt.Sprintf("Hello: %s/%s/%s/%s/%s, %d", EndpointPublic, AccessId, AccessKey, BucketName, key, len(val)))

	bucket, err := GetBucket(BucketName)
	if err != nil {
		return err
	}

	err = bucket.PutObject(key, strings.NewReader(val))
	if err != nil {
		return err
	}

	return nil
}
