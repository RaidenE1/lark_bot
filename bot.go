package lark_bot

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

// Bot 飞书机器人结构体
type Bot struct {
    webhook string
	serviceName string
	podID string
}


// NewBot 初始化机器人
func NewBot(webhook string, serviceName string, podID string) *Bot {
    return &Bot{
        webhook: webhook,
		serviceName: serviceName,
		podID: podID,
    }
}

// SendText 发送文本消息
func (b *Bot) SendText(text string) error {
    content := TextContent{
        Text: text,
    }
    
    msg := Message{
        MsgType: "text",
        Content: content,
    }
    
    jsonData, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal message failed: %v", err)
	}

    return b.send(jsonData)
}

func (b *Bot) SendCard(template EventCardTemplate) error {
	jsonData, err := json.Marshal(template)
	if err != nil {
		return fmt.Errorf("marshal message failed: %v", err)
	}

	return b.send(jsonData)
}

func (b *Bot) send(jsonData []byte) error {
    resp, err := http.Post(b.webhook, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return fmt.Errorf("send message failed: %v", err)
    }
    defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("send message failed with status code: %d", resp.StatusCode)
	}

	return nil
}
