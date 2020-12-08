package wxa

import (
	"github.com/charsunny/wechat/mp/core"
)

type TemplateCategory struct {
	Id   string `json:"id"`   // 作为模板时返回
	Name string `json:"name"` // 作为模板时和作为添加到模板库后都返回
}

type TemplateKey struct {
	Id         string `json:"kid"`         // 作为模板时返回
	Name       string `json:"name"`        // 作为模板时和作为添加到模板库后都返回
	TemplateId string `json:"template_id"` // 作为添加到模板库后返回
	Rule       string `json:"rule"`        // 作为添加到模板库后返回
	Example    string `json:"example"`     // 作为添加到模板库后返回
}

type TemplateTitle struct {
	Id         string `json:"tid"`        // 作为模板时返回
	Title      string `json:"title"`      // 作为模板时和作为添加到模板库后都返回
	Type       int    `json:"type"`       // 作为添加到模板库后返回
	CategoryId int    `json:"categoryId"` // 作为添加到模板库后返回
}

type Template struct {
	Id      string `json:"priTmplId"` // 作为模板时返回
	Title   string `json:"title"`     // 作为模板时和作为添加到模板库后都返回
	Content string `json:"content"`   // 作为添加到模板库后返回
	Type    int    `json:"type"`      // 作为添加到模板库后返回
	Example string `json:"example"`   // 作为添加到模板库后返回
}

type WxaTemplateMsg struct {
	TemplateId      string                 `json:"template_id"`      // 模板id
	Page            string                 `json:"page"`             //跳转小程序页面
	FormId          string                 `json:"form_id"`          // formid
	Data            map[string]interface{} `json:"data"`             // 模板数据
	EmphasisKeyword string                 `json:"emphasis_keyword"` //加重数据
}

type MpTemplateMsg struct {
	TemplateId  string                 `json:"template_id"` // 模板id
	AppId       string                 `json:"appid"`       //跳转小程序页面
	Url         string                 `json:"url"`         // formid
	Data        map[string]interface{} `json:"data"`        // 模板数据
	MiniProgram struct {               //跳转小程序页面
		AppId    string `json:"appid"`
		PagePath string `json:"pagepath"`
	} `json:"miniprogram"`
}

// 组合模板并添加至帐号下的个人模板库
// id	string		是	模板标题id，可通过接口获取，也可登录小程序后台查看获取
// keyword_id_list	[]int 是  开发者自行组合好的模板关键词列表，关键词顺序可以自由搭配（例如[3,5,4]或[4,5,3]），最多支持10个关键词组合
func AddTemplate(clt *core.Client, id string, keyIds []int, desc string) (template_id string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxaapi/newtmpl/addtemplate?access_token="

	var request = struct {
		Id            string `json:"tid"`     //路径
		KeywordIdList []int  `json:"kidList"` //宽度
		SceneDesc     string `json:"sceneDesc"`
	}{
		Id:            id,
		KeywordIdList: keyIds,
		SceneDesc:     desc,
	}
	var result struct {
		core.Error
		TemplateId string `json:"priTmplId"`
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	template_id = result.TemplateId
	return
}

// 组合模板并添加至帐号下的个人模板库
// template_id	string		要删除的模板id
func DeleteTemplate(clt *core.Client, template_id string) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxaapi/newtmpl/deltemplate?access_token="

	var request = struct {
		TemplateId string `json:"priTmplId"` //路径
	}{
		TemplateId: template_id,
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

// 获取模板的keyword列表
// id	string		是	模板标题id，可通过接口获取，也可登录小程序后台查看获取
func GetTemplateCategory(clt *core.Client) (list []*TemplateCategory, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxaapi/newtmpl/getcategory?access_token="

	var request = struct {
	}{}
	var result struct {
		core.Error
		Data []*TemplateCategory `json:"data"`
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.Data
	return
}

// 获取模板的keyword列表
// id	string		是	模板标题id，可通过接口获取，也可登录小程序后台查看获取
func GetPubTemplateKeywords(clt *core.Client, id string) (count int64, templates []*TemplateKey, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatekeywords?access_token="

	var request = struct {
		Id string `json:"tid"` //路径
	}{
		Id: id,
	}
	var result struct {
		core.Error
		Count int64          `json:"count"`
		Data  []*TemplateKey `json:"data"`
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	count = result.Count
	templates = result.Data
	return
}

// 获取模板的keyword列表
// id	string		是	模板标题id，可通过接口获取，也可登录小程序后台查看获取
func GetPubTemplateTitles(clt *core.Client, ids string, page, count int) (total_count int64, list []*TemplateTitle, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatetitles?access_token="

	var request = struct {
		Ids    string `json:"ids"`   //类目 id，多个用逗号隔开
		Offset int    `json:"start"` //路径
		Count  int    `json:"limit"` //路径
	}{
		Ids:    ids,
		Offset: page,
		Count:  count,
	}
	var result struct {
		core.Error
		Count int64            `json:"count"`
		Data  []*TemplateTitle `json:"data"`
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	total_count = result.Count
	list = result.Data
	return
}

// 获取小程序模板库列表
// count: 用于分页，表示拉取count条记录。最大为 20。最后一页的list长度可能小于请求的count。
func GetTemplateList(clt *core.Client, page, count int) (list []*Template, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxaapi/newtmpl/gettemplate?access_token="

	var request = struct {
	}{}
	var result struct {
		core.Error
		Data []*Template `json:"data"`
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.Data
	return
}

// 发送模板消息
// data: { "k1":{"value":""}}
// miniprogram_state: developer为开发版；trial为体验版；formal为正式版；默认为正式版
func SendMessage(clt *core.Client, touser, template_id, page, miniprogram_state string, data map[string]interface{}) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token="

	var request = struct {
		ToUser           string                 `json:"touser"`            // 用户openid
		TemplateId       string                 `json:"template_id"`       // 模板id
		Page             string                 `json:"page"`              //跳转小程序页面
		MiniprogramState string                 `json:"miniprogram_state"` // formid
		Data             map[string]interface{} `json:"data"`              // 模板数据
	}{
		ToUser:           touser,
		TemplateId:       template_id,
		Page:             page,
		MiniprogramState: miniprogram_state,
		Data:             data,
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
