package autoreply

import "github.com/charsunny/wechat/mp/core"

const (
	MsgTypeText  core.MsgType = "text"  // 文本消息
	MsgTypeImage core.MsgType = "image" // 图片消息
	MsgTypeVoice core.MsgType = "voice" // 语音消息
	MsgTypeVideo core.MsgType = "video" // 视频消息
	MsgTypeNews  core.MsgType = "news"  // 图文消息
)

type AutoReply struct {
	IsAddFriendReplyOpen   int `json:"is_add_friend_reply_open,omitempty"`
	IsAutoreplyOpen        int `json:"is_autoreply_open,omitempty"`
	AddFriendAutoreplyInfo struct {
		Type    string `json:"type,omitempty"`
		Content string `json:"content,omitempty"`
	} `json:"add_friend_autoreply_info"`
	MessageDefaultAutoreplyInfo struct {
		Type    string `json:"type,omitempty"`
		Content string `json:"content,omitempty"`
	} `json:"message_default_autoreply_info"`
	KeywordAutoreplyInfo struct {
		List []KeywordAutoreplyInfo `json:"list,omitempty"`
	} `json:"keyword_autoreply_info"`
}

type KeywordAutoreplyInfo struct {
	RuleName        string            `json:"rule_name,omitempty"`
	CreateTime      int64             `json:"create_time,omitempty"`
	ReplyMode       string            `json:"reply_mode,omitempty"`
	KeywordListInfo []KeywordListInfo `json:"keyword_list_info,omitempty"`
	ReplyListInfo   []ReplyListInfo   `json:"reply_list_info,omitempty"`
}

type KeywordListInfo struct {
	Type      string `json:"type,omitempty"`
	MatchMode string `json:"match_mode,omitempty"`
	Content   string `json:"content,omitempty"`
}

type ReplyListInfo struct {
	Type     string `json:"type,omitempty"`
	Content  string `json:"content,omitempty"`
	NewsInfo struct {
		List []NewsInfo `json:"list,omitempty"`
	} `json:"news_info"`
}

type NewsInfo struct {
	Title      string `json:"title,omitempty"`
	Author     string `json:"author,omitempty"`
	Digest     string `json:"digest,omitempty"`
	ShowCover  int    `json:"show_cover,omitempty"`
	CoverUrl   string `json:"cover_url,omitempty"`
	ContentUrl string `json:"content_url,omitempty"`
	SourceUrl  string `json:"source_url,omitempty"`
}
