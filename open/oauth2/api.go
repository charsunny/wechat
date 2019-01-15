package oauth2

import (
	"github.com/charsunny/wechat/open/core"
)

// 获取预授权码 配合 授权链接获取的authcode 一起服用， 换取token
func GetPreAuthCode(clt *core.Client) (code string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode?component_access_token="

	var result struct{
		core.Error
		PreAuthCode string `json:"pre_auth_code"`
		ExpiresIn int `json:"expires_in"`
	}
	if err = clt.PostJSON(incompleteURL, map[string]interface{} {"component_appid": clt.AuthServer.AppId()}, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	code = result.PreAuthCode
	return
}

// 获取授权信息  授权链接获取的authcode 一起服用， 换取token
func GetAuthInfo(clt *core.Client, auth_code string) (authInfo AuthorizationInfo, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/component/api_query_auth?component_access_token="

	var result struct{
		core.Error
		authorizationInfo AuthorizationInfo `json:"authorization_info"`
	}
	if err = clt.PostJSON(incompleteURL, map[string]interface{} {"component_appid": clt.AuthServer.AppId(), "authorization_code": auth_code}, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	authInfo = result.authorizationInfo
	return
}