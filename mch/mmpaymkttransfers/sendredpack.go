package mmpaymkttransfers

import (
	"github.com/charsunny/wechat/mch/core"
)

// 红包发放.
//  NOTE: 请求需要双向证书
func SendRedPack(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/mmpaymkttransfers/sendredpack", req)
}

// 小程序红包发放.
//  NOTE: 请求需要双向证书
func SendMpRedPack(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/mmpaymkttransfers/sendminiprogramhb", req)
}
