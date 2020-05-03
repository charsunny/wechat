package marketcode

import (
	"github.com/charsunny/wechat/mp/core"
)

// 申请二维码接口
// @param count 申请码的数量 ≥10000，≤20000000，10000的整数倍
// @param isv_application_id 相同isv_application_id视为同一申请单
func ApplyCode(clt *core.Client, count int64, isv_application_id string) (application_id int64, err error) {
	const incompleteURL = "https://api.weixin.qq.com/intp/marketcode/applycode?access_token="

	req := map[string]interface{} {
		"code_count": count,
		"isv_application_id": isv_application_id,
	}
	var result struct{
		ApplicationId int64 `json:"application_id"`
		core.Error
	}
	if err = clt.PostJSON(incompleteURL, &req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	application_id = result.ApplicationId
	return
}


//  查询二维码申请单接口
// @param application_id 申请单号
// @param isv_application_id 相同isv_application_id视为同一申请单
func QueryApplyCode(clt *core.Client, application_id int64, isv_application_id string) (apply MarketCodeApply, err error) {
	const incompleteURL = "https://api.weixin.qq.com/intp/marketcode/applycodequery?access_token="

	req := map[string]interface{} {
		"application_id": application_id,
		"isv_application_id": isv_application_id,
	}
	var result struct{
		MarketCodeApply
		core.Error
	}
	if err = clt.PostJSON(incompleteURL, &req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	apply = result.MarketCodeApply
	return
}

// 下载二维码包接口
// @param code_start	开始位置	Uint64	Y	来自查询二维码申请接口
// @param code_end	结束位置	Uint64	Y	来自查询二维码申请接口
// @return buffer	文件buffer	String128	Y	需要先base64 decode，再做解密操作（解密参见3.1）
func DownloadApplyCode(clt *core.Client, code_start int64, code_end int64) (buffer string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/intp/marketcode/applycodedownload?access_token="

	req := map[string]interface{} {
		"code_start": code_start,
		"code_end": code_end,
	}
	var result struct{
		Buffer string `json:"buffer"`
		core.Error
	}
	if err = clt.PostJSON(incompleteURL, &req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	buffer = result.Buffer
	return
}

// 激活二维码接口
func ActiveCode(clt *core.Client, info *ActiveCodeInfo) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/intp/marketcode/codeactive?access_token="
	var result core.Error
	if err = clt.PostJSON(incompleteURL, info, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}

// 激活二维码接口
// 字段名	中文解释	类型	是否必填	备注
// application_id	申请单号	Uint64	N	无
// code_index	该码在批次中的偏移量	Uint64	N	传入application_id时必填
// code_url	28位普通码字符	String128	N	code与code_url二选一
// code	九位的字符串原始码	String16	N	code与code_url二选一
func QueryActiveCode(clt *core.Client, application_id , index int64,  url, code string ) (info ActiveCodeInfo, err error) {
	const incompleteURL = "https://api.weixin.qq.com/intp/marketcode/codeactive?access_token="

	req := map[string]interface{} {
		"application_id": application_id,
		"code_index": index,
		"code_url": url,
		"code": code,
	}

	var result struct{
		core.Error
		ActiveCodeInfo
	}
	if err = clt.PostJSON(incompleteURL, &req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	info = result.ActiveCodeInfo
	return
}

// code_ticket换code接口
// 字段名	中文解释	类型	是否必填	备注
// openid	用户的openid	String	Y	无
// code_ticket	跳转时带上的code_ticket参数	String256	Y	无
func TicketToCode(clt *core.Client, openid, code_ticket string ) (info ActiveCodeInfo, err error) {
	const incompleteURL = "https://api.weixin.qq.com/intp/marketcode/tickettocode?access_token="

	req := map[string]interface{} {
		"openid": openid,
		"code_ticket": code_ticket,
	}

	var result struct{
		core.Error
		ActiveCodeInfo
	}
	if err = clt.PostJSON(incompleteURL, &req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	info = result.ActiveCodeInfo
	return
}