package service

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"goWebDemo/utils"
	"goWebDemo/utils/errmsg"
	"mime/multipart"
)

var AccessKey = utils.AccessKey
var SecretKey = utils.SecretKey
var Bucket = utils.Bucket
var ImgUrl = utils.EndPoint

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
	url := ImgUrl + ret.Key
	return url, errmsg.Success
}
