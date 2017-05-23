package models

type MessageParameterData struct {
	Type    int    `json:"type"`    // 数据类型
	Content string `json:"content"` // 数据内容
}

type MessageParameter struct {
	Action int                     `json:"action"` // 增删改查
	Target int                     `json:"target"` // 处理目标对象
	Data   []*MessageParameterData `json:"data"`   // 请求数据
}

type MessageRole struct {
	Type int   `json:"type"` // 类型
	Id   int64 `json:"id"`   // 唯一ID
}

type MessageRequest struct {
	Token     string            `json:"token"`     // message有效性判断
	SessionId string            `json:"sessionId"` // 会话ID
	From      *MessageRole      `json:"from"`      // 请求发起者
	To        *MessageRole      `json:"to"`        // 请求接受者
	Parameter *MessageParameter `json:"parameter"` // 参数
}

type MessageResponse struct {
	Token      string `json:"token"`      // message有效性判断
	SessionId  string `json:"sessionId"`  // 会话ID
	ResultCode int    `json:"resultCode"` // 成功失败标识号
	Reason     string `json:"reason"`     // 成功失败原因
	Target     int    `json:"target"`     // 返回数据类型
	Result     string `json:"result"`     // 返回数据内容
}
