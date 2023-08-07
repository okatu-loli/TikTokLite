// Code generated by goctl. DO NOT EDIT.
package types

type Message struct {
	Id         int64  `json:"id"`
	ToUserId   int64  `json:"to_user_id"`
	FromUserId int64  `json:"from_user_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

type DouyinMessageChatRequestt struct {
	Token      string `form:"token"`
	ToUserId   int64  `form:"to_user_id"`
	PreMsgTime int64  `form:"pre_msg_time"` // 上次最新消息的时间
}

type DouyinMessageChatResponse struct {
	StatusCode  int32     `json:"status_code"` // 0-成功, 其他-失败
	StatusMsg   string    `json:"status_msg"`
	MessageList []Message `json:"message_list"`
}

type DouyinMessageActionRequest struct {
	Token      string `form:"token"`
	ToUserId   int64  `form:"to_user_id"`
	ActionType int32  `form:"action_type"` // 1-发送消息
	Content    string `form:"content"`
}

type DouyinMessageActionResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}
