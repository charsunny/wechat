package oauth2

import (
	"github.com/charsunny/wechat/open/core"
	"fmt"
)

// 获取预授权码 配合 授权链接获取的authcode 一起服用， 换取token
func GetPreAuthCode(clt *core.Client) (code string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode?component_access_token="

	var result struct {
		core.Error
		PreAuthCode string `json:"pre_auth_code"`
		ExpiresIn int `json:"expires_in"`
	}
	if err = clt.PostJSON(incompleteURL, map[string]interface{} {"component_appid": clt.AuthServer.AppId()}, &result); err != nil {
		return
	}
	fmt.Printf("%v, %v", result.ErrCode, result.ErrMsg)
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
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
	req := map[string]interface{} {
		"component_appid": clt.AuthServer.AppId(),
		"authorization_code": auth_code,
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	authInfo = result.authorizationInfo
	return
}

// 刷新第三方帐号的授权信息
// 注意： 此处appid和refresh token 均为第三方帐号的appid和refreshtoken， 不是开放平台的appid和token
func RefreshAuthInfo(clt *core.Client, appId string, refreshToken string) (authInfo AuthorizationInfo, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/component/api_query_auth?component_access_token="

	var result struct{
		core.Error
		authorizationInfo AuthorizationInfo `json:"authorization_info"`
	}
	req := map[string]interface{} {
		"component_appid": clt.AuthServer.AppId(),
		"authorizer_appid": appId,
		"authorizer_refresh_token": refreshToken,
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	authInfo = result.authorizationInfo
	return
}

// 获取授权方的授权信息
// 获取授权app的具体信息
func GetAuthAppInfo(clt *core.Client, appId string) (authInfo map[string]interface{}, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_info?component_access_token="

	var result struct{
		core.Error
		authorizationInfo AuthorizationInfo `json:"authorization_info"`
	}
	req := map[string]interface{} {
		"component_appid": clt.AuthServer.AppId(),
		"authorizer_appid": appId,
	}
	authInfo = make(map[string]interface{})
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}