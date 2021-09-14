package service

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/spf13/viper"
	"goWebDemo/utils/errmsg"
	"mime/multipart"
)

var AccessKey = viper.GetString("Qiniu.AccessKey")
var SecretKey = viper.GetString("Qiniu.SecretKey")
var Bucket = viper.GetString("Qiniu.Bucket")
var EndPoint = viper.GetString("Qiniu.EndPoint")

func UpLoadFile(file multipart.File, filSize int64) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)

	config := storage.Config{
		Zone:          &storage.Zone_as0,   // 东南亚
		UseHTTPS:      false,
		UseCdnDomains: false,
	}

	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&config)
	ret := storage.PutRet{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, filSize, &putExtra)

	if err != nil {
		return "", errmsg.Error
	}
	url := EndPoint + ret.Key
	return url, errmsg.Success
}
