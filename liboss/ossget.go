package liboss

import "io/ioutil"

/*	Get the string text from the blob file
*	key is the source path in blob
 */
func GetString(key string) (string, error) {

	bucket, err := GetBucket(BucketName)
	if err != nil {
		return "", err
	}

	body, err := bucket.GetObject(key)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(body)
	body.Close()
	if err != nil {
		return "", err
	}

	return string(data), nil
}

/*	Detect if the blob file exists
*	key is the source path in blob
 */
func IsObjectExisted(key string) bool {

	bucket, err := GetBucket(BucketName)
	if err != nil {
		return false
	}

	if existed, err := bucket.IsObjectExist(key); err != nil {
		return false
	} else {
		return existed
	}
}
