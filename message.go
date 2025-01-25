package lark_bot

type Message struct {
    MsgType string      `json:"msg_type"`
    Content interface{} `json:"content"`
}

type EventCardTemplate struct {
	Title string `json:"title"`
    ServiceName string `json:"service_name"`
	PodID string `json:"pod_id"`
	EventLevel string `json:"event_level"`
	EventMessage string `json:"event_message"`
	EventTime string `json:"event_time"`
}

// TextContent 文本消息内容
type TextContent struct {
    Text string `json:"text"`
}