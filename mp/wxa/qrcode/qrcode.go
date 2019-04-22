package qrcode

import (
	"github.com/charsunny/wechat/mp/core"
)

// 接口生成的小程序码，永久有效，有数量限制
// path	string		是	扫码进入的小程序页面路径，
// width	number	430	否	二维码的宽度，单位 px。最小 280px，最大 1280px
func CreateQRCode(clt *core.Client, path string, width int) (data []byte, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token="

	var request = struct {
		Path string `json:"path"`	//路径
		Width int `json:"width"`	//宽度
	}{
		Path:path,
		Width:width,
	}
	var result struct {
		core.Error
	}
	if data, err = clt.PostJsonData(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}

type LineColor struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

// 接口生成的小程序码，永久有效，有数量限制
// path	string		是	扫码进入的小程序页面路径，必须是已经发布的小程序存在的页面（否则报错），例如 pages/index/index, 根路径前不要填加 /,不能携带参数（参数请放在scene字段里），如果不填写这个字段，默认跳主页面
// width	number	430	否	二维码的宽度，单位 px。最小 280px，最大 1280px
// auto_color	boolean	false	否	自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调，默认 false
// line_color	Object	{"r":0,"g":0,"b":0}	否	auto_color 为 false 时生效，
// is_hyaline	boolean	false	否	是否需要透明底色，为 true 时，生成透明底色的小程序
func CreateWxaCode(clt *core.Client, path string, width int, lineColor *LineColor, auto_color, is_hyaline bool) (data []byte, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/getwxacode?access_token="

	var request = struct {
		Path string `json:"path"`	//路径
		Width int `json:"width"`	//宽度
		AutoColor bool `json:"auto_color"`
		LineColor *LineColor `json:"line_color,omitempty"`
		IsHyaline bool `json:"is_hyaline"`
	}{
		Path:path,
		Width:width,
		AutoColor:auto_color,
		LineColor:lineColor,
		IsHyaline:is_hyaline,
	}
	var result struct {
		core.Error
	}
	if data, err = clt.PostJsonData(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}

// 接口生成的小程序码，永久有效，有数量限制
// path	string		是	扫码进入的小程序页面路径，必须是已经发布的小程序存在的页面（否则报错），例如 pages/index/index, 根路径前不要填加 /,不能携带参数（参数请放在scene字段里），如果不填写这个字段，默认跳主页面
// width	number	430	否	二维码的宽度，单位 px。最小 280px，最大 1280px
// auto_color	boolean	false	否	自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调，默认 false
// line_color	Object	{"r":0,"g":0,"b":0}	否	auto_color 为 false 时生效，
// is_hyaline	boolean	false	否	是否需要透明底色，为 true 时，生成透明底色的小程序
func CreateWxaCodeUnlimited(clt *core.Client, scene string, path string, width int, lineColor *LineColor, auto_color, is_hyaline bool) (data []byte, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token="

	var request = struct {
		Path string `json:"path"`	//路径
		Width int `json:"width"`	//宽度
		Scene string `json:"scene"`
		AutoColor bool `json:"auto_color"`
		LineColor *LineColor `json:"line_color,omitempty"`
		IsHyaline bool `json:"is_hyaline"`
	}{
		Path:path,
		Width:width,
		AutoColor:auto_color,
		LineColor:lineColor,
		Scene: scene,
		IsHyaline:is_hyaline,
	}
	var result struct {
		core.Error
	}
	if data, err = clt.PostJsonData(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}