package merchant

import (
	"github.com/charsunny/wechat/mp/core"
)

// 绑定微信用户为小程序体验者
func GetDistrict(clt *core.Client) (result map[string]interface{}, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/get_district?access_token="
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	return
}
