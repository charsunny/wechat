package wxa

import (
	"github.com/charsunny/wechat/open/core"
)

type WxaCodeDraftInfo struct {
	DraftId string `json:"draft_id"`	// 展示的公众号appid
	UserDesc string `json:"user_desc"`	// 展示的公众号nickname
	UserVersion string `json:"user_version"`	// 展示的公众号头像
	CreateTime int `json:"create_time"`	// 是否可以设置 1 可以，0，不可以
}

type WxaCodeTemplateInfo struct {
	TemplateId string `json:"template_id"`	// 展示的公众号appid
	UserDesc string `json:"user_desc"`	// 展示的公众号nickname
	UserVersion string `json:"user_version"`	// 展示的公众号头像
	CreateTime int `json:"create_time"`	// 是否可以设置 1 可以，0，不可以
}

// 绑定微信用户为小程序体验者
func GetCodeTemplateDraftList(clt *core.Client) (list []*WxaCodeDraftInfo, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/gettemplatedraftlist?access_token="

	var result struct {
		core.Error
		TemplateList []*WxaCodeDraftInfo `json:"template_list"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.TemplateList
	return
}

// 获取代码模版库中的所有小程序代码模版
func GetCodeTemplateList(clt *core.Client) (list []*WxaCodeTemplateInfo, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/gettemplatelist?access_token="

	var result struct {
		core.Error
		TemplateList []*WxaCodeTemplateInfo `json:"template_list"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.TemplateList
	return
}


// 将草稿箱的草稿选为小程序代码模版
func AddToTemplate(clt *core.Client, draft_id int) (err error)  {
	const incompleteURL = "https://api.weixin.qq.com/wxa/addtotemplate?access_token="

	var request = struct {
		DraftId int `json:"draft_id"`	// 反馈内容，不超过200字, 只有上个版本被驳回，才能使用“feedback_info”、“feedback_stuff”这两个字段，否则忽略处理
	}{
		DraftId:draft_id,
	}
	var result struct {
		core.Error
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}

// 将草稿箱的草稿选为小程序代码模版
func DeleteTemplate(clt *core.Client, template_id int) (err error)  {
	const incompleteURL = "https://api.weixin.qq.com/wxa/deletetemplate?access_token="

	var request = struct {
		TemplateId int `json:"template_id"`	// 反馈内容，不超过200字, 只有上个版本被驳回，才能使用“feedback_info”、“feedback_stuff”这两个字段，否则忽略处理
	}{
		TemplateId:template_id,
	}
	var result struct {
		core.Error
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}
