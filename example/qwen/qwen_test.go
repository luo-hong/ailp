package qwen_test

import (
	"ailp/internal/channel/qwen"
	"fmt"
	"strings"
	"testing"

	"github.com/alevinval/sse/pkg/base"
)

func TestRun(t *testing.T) {

	// 初始化QWEN聊天机器人客户端，使用您的API密钥
	apiKey := "your api key"
	qwenclient := qwen.NewWithDefaultChat(apiKey)

	//qwenclient.QWenModel = "new model"

	// 定义一条消息对话的历史记录
	messages := []qwen.Messages{
		{Role: qwen.ChatUser, Content: "你好"},
	}

	// 获取AI对消息的回复
	resp, err := qwenclient.GetAIReply(messages)
	if err != nil {
		fmt.Printf("获取AI回复失败：%v\n", err)
		return
	}

	// 打印收到的回复
	fmt.Printf("收到的回复：%v\n", resp.Output.Text)

}

func TestRunStream(t *testing.T) {
	apiKey := "your api key"
	qwenclient := qwen.NewWithDefaultChat(apiKey)

	//qwenclient.QWenModel = "new model"

	// 定义一条消息对话的历史记录
	messages := []qwen.Messages{
		{Role: qwen.ChatUser, Content: "你好"},
		{Role: qwen.ChatBot, Content: "你好！有什么我能为你做的吗？"},
		{Role: qwen.ChatUser, Content: "给我推荐一款便宜耐用的衬衫"},
	}

	// 获取AI对消息的回复
	err := qwenclient.GetAIReplyStream(messages, func(e *base.MessageEvent) error {
		if length := len(e.Data); length < 1 {
			return nil
		}

		// 解析并处理消息中的来源信息。
		switch strings.ToLower(e.Name) {
		case "result":
			fmt.Println("data:", e.Data)
		}
		return nil
	})
	if err != nil {
		fmt.Println("err", err)
	}

}
