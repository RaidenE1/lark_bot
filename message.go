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
	TemplateID string `json:"template_id"`
	TemplateVersion string `json:"template_version"`
	TemplateVariable interface{} `json:"template_variable"`
}

// TextContent 文本消息内容
type TextContent struct {
    Text string `json:"text"`
}

func newMessageCard(data eventCardTemplate) messageCard {
	card := card{
		Type: "template",
		Data: data,
	}
	msg := messageCard{
		MsgType: "interactive",
		Card: card,
	}
	return msg
}