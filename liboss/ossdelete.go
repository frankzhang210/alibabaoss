package liboss

import (
	"errors"
	"strings"
)

func Delete(key string) error {

	key = strings.TrimSpace(key)
	if key == "" {
		return errors.New("Blob file name cannot be empty")
	}

	bucket, err := GetBucket(BucketName)
	if err != nil {
		return err
	}

	err = bucket.DeleteObject(key)
	if err != nil {
		return err
	}

	return nil
}
