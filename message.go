package lark_bot

type message struct {
    MsgType string      `json:"msg_type"`
    Content interface{} `json:"content"`
}

type messageCard struct {
	MsgType string      `json:"msg_type"`
    Card card `json:"card"`
}

type card struct {
	Type string `json:"type"`
	Data eventCardTemplate `json:"data"`
}

type eventCardTemplate struct {
	Title string `json:"title"`
	TemplateID string `json:"template_id"`
	TemplateVersion string `json:"template_version"`
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

func newMessageCard(eventCardTemplate eventCardTemplate) messageCard {
	card := card{
		Type: "template",
		Data: eventCardTemplate,
	}
	msg := messageCard{
		MsgType: "message_card",
		Card: card,
	}
	return msg
}