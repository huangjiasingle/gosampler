package main

import (
	"github.com/golang/glog"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

var (
	// accessKey = os.Getenv("QINIU_ACCESS_KEY")
	// secretKey = os.Getenv("QINIU_SECRET_KEY")
	// bucket    = os.Getenv("QINIU_TEST_BUCKET")
	// io        = os.Getenv("QINIU_IO")
	// up        = os.Getenv("QINIU_UP")
	// api       = os.Getenv("QINIU_API")
	mac       = qbox.NewMac("RlxNug-8lvTxdMrbK3lUzUDDzAWI0A6wX9MQ0Rw_", "14AGJ17e3X79_EV2QZ7_uZ0mhABhKYg--onyXhSA")
	putPolicy = storage.PutPolicy{Scope: "backup", Expires: 3600 * 24 * 365} //Expires 最大值为1年
	upToken   = putPolicy.UploadToken(mac)

	zone = &storage.Zone{
		SrcUpHosts: []string{
			"up.oss.test",
		},
		RsHost:    "rs.oss.test",
		ApiHost:   "apiserver.oss.test",
		IovipHost: "io.oss.test",
	}
)

// Upload upload file to qiuniu
func main() {
	glog.Info(upToken)
	cfg := storage.Config{UseHTTPS: false, UseCdnDomains: false}
	cfg.Zone = zone
	// formUploader := storage.NewFormUploader(&cfg)
	// glog.Info(formUploader.PutFile(context.Background(), &storage.PutRet{}, upToken, "main.go2", "main.go", nil))

	bucketManager := storage.NewBucketManager(mac, &cfg)
	glog.Info(bucketManager.Delete("backup", "demo.20180806103952.sql.gz"))

	// bucket := "backup"
	// keys := []string{
	// 	"main.go",
	// 	"main.go1",
	// 	"main.go2",
	// 	"prometheus-2.3.2.darwin-amd64.tar.gz",
	// 	"zrqtest",
	// }
	// deleteOps := make([]string, 0, len(keys))
	// for _, key := range keys {
	// 	deleteOps = append(deleteOps, storage.URIDelete(bucket, key))
	// }
	// _, err := bucketManager.Batch(deleteOps)
	// glog.Info(err)
}

func Download(privateAccessURL string) error {
	return nil
}
