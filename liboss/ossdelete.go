package liboss

import (
	"errors"
	"log"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
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

func DeleteByPrefix(key string) error {

	key = strings.TrimSpace(key)
	if key == "" {
		return errors.New("Prefix cannot be empty")
	}

	bucket, err := GetBucket(BucketName)
	if err != nil {
		return err
	}

	lor, err := bucket.ListObjects(oss.Prefix(key))
	if err != nil {
		return err
	}
	// fmt.Println("my objects prefix :", getObjectsFromResponse(lor))

	objects := make([]string, len(lor.Objects))
	for i, obj := range lor.Objects {
		log.Println("key with same prefix ", obj.Key, i)
		objects[i] = obj.Key
	}

	delRes, err := bucket.DeleteObjects(objects, oss.DeleteObjectsQuiet(false))
	if err != nil {
		return err
	}
	log.Println("response: ", delRes)

	return nil
}
