package main

type downloadFromNetDisk struct {
	/*	1、登陆
		2、下载
	*/
	secret   DynamicSecret
	filePath string
}

func (dd *downloadFromNetDisk) DownlaodFile() {
	if err := dd.logincheck(); err != nil {
		// todo 重新登陆
	}
	dd.downloadFromAliYun(dd.filePath)
}

func (dd downloadFromNetDisk) logincheck() error {
	dd.checkSecret(dd.secret.GetSecret())
	return nil
}

func (dd *downloadFromNetDisk) downloadFromAliYun(file string) {
}

func (dd *downloadFromNetDisk) checkSecret(secret string) {
	//todo 调用验证接口验证密码是否有效
}

func main() {
	dd := &downloadFromNetDisk{
		secret:   &mobileTokenDynamic{mobileNumber: "123456"},
		filePath: "接口编程.ppt",
	}
	dd.DownlaodFile()
}

type DynamicSecret interface {
	GetSecret() string
}

type mobileTokenDynamic struct {
	mobileNumber string
}

func (d *mobileTokenDynamic) GetSecret() string {
	return "something"
}

// 通常开始的时候，第一个版本叫做 happy path
// 剩下的是通：它无法应对变更，简单的变更会带来痛苦的维护
