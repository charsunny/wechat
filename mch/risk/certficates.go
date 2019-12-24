package risk

import (
	"github.com/charsunny/wechat/mch/core"
)

// 查询代金券信息.
func GetCertficates(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	if req["sign_type"] == "" {
		req["sign_type"] = core.SignType_HMAC_SHA256
	}
	return clt.PostXML(core.APIBaseURL()+"/risk/getcertficates", req)
}
