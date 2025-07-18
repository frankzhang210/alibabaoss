package liboss

import (
	"encoding/json"
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

var (
	// kmsID      = os.Getenv("OSS_TEST_KMS_ID")

	EndpointPublic, _  = beego.AppConfig.String("OSSEndPoint::Public")
	EndpointClassic, _ = beego.AppConfig.String("OSSEndPoint::ClassicFromECS")
	EndpointVPC, _     = beego.AppConfig.String("OSSEndPoint::VPCFromECS")
	AccessId, _        = beego.AppConfig.String("OSSEndPoint::AccessKeyId")
	AccessKey, _       = beego.AppConfig.String("OSSEndPoint::AccessKeySecret")
	BucketName, _      = beego.AppConfig.String("OSSEndPoint::Bucket")
)

type OSSConfig struct {
	Public          string
	ClassicFromECS  string
	VPCFromECS      string
	AccessKeyId     string
	AccessKeySecret string
	Bucket          string
}

/*	If EndpointPublic is empty, in the caller main.go, add below two lines of code into var section and make sure such json file exists
 *	//go:embed conf/oss.config.json
 *	ossConfigJson string
 *
 *	Then call this method and pass in ossConfigJson string
 */
func OverrideOSS(configJson string) {
	var oss OSSConfig
	if err3 := json.Unmarshal([]byte(configJson), &oss); err3 != nil {
		fmt.Println(err3)
	}

	AccessId = oss.AccessKeyId
	AccessKey = oss.AccessKeySecret
	BucketName = oss.Bucket
	EndpointClassic = oss.ClassicFromECS
	EndpointPublic = oss.Public
	EndpointVPC = oss.VPCFromECS
}
