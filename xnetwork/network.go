package xnetwork

import (
	"utils/xos"
	"github.com/imroc/req"
)

/**
下载文件到本地, 按网络路径规则自动创建本地目录
*/
func DownloadFileByUrl(url string, rootDir string) {
	uri := ParsePath(url)

	// 创建本地目录
	localFullPath := rootDir + uri
	DownloadFileByLocalPath(url, localFullPath)
}

/**
下载文件到本地
*/
func DownloadFileByLocalPath(url string, localPath string) error {
	// 创建本地目录
	err := xos.CreateFileDir(localPath)
	if err != nil {
		panic(err)
	}

	resp, err := req.Get(url)
	if err != nil {
		return err
	}

	err = resp.ToFile(localPath)
	if err != nil {
		panic(err)
	}

}
