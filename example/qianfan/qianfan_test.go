package qianfan_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/luo-hong/ailp/internal/channel/qianfan"
)

func TestRun(t *testing.T) {
	// 使用前请先设置 AccessKey 和 SecretKey，通过环境变量设置可省略如下两行
	qianfan.GetConfig().AccessKey = "your_access_key"
	qianfan.GetConfig().SecretKey = "your_secret_key"

	// 多轮对话，调用默认模型
	chat := qianfan.NewChatCompletion(
		qianfan.WithModel("ERNIE-Speed-128K"),
	)
	resp, _ := chat.Do(
		context.TODO(),
		&qianfan.ChatCompletionRequest{
			Messages: []qianfan.ChatCompletionMessage{
				qianfan.ChatCompletionUserMessage("你好！"),
				qianfan.ChatCompletionAssistantMessage("你好！有什么我可以帮助你的吗？"),
				qianfan.ChatCompletionUserMessage("我在成都，周末可以去哪里玩？"),
			},
		},
	)
	fmt.Println(resp.Result)

}

func TestRunByStream(t *testing.T) {
	// 使用前请先设置 AccessKey 和 SecretKey，通过环境变量设置可省略如下两行
	qianfan.GetConfig().AccessKey = "your_access_key"
	qianfan.GetConfig().SecretKey = "your_secret_key"

	// 多轮对话，调用默认模型
	chat := qianfan.NewChatCompletion(
		qianfan.WithModel("ERNIE-Speed-128K"),
	)
	resp, _ := chat.Stream( // Stream 启用流式返回，参数与 Do 相同
		context.TODO(),
		&qianfan.ChatCompletionRequest{
			Messages: []qianfan.ChatCompletionMessage{
				qianfan.ChatCompletionUserMessage("你好,在成都周末推荐去那玩？"),
			},
		},
	)
	for {
		r, err := resp.Recv()
		if err != nil {
			panic(err)
		}
		if resp.IsEnd { // 判断是否结束
			break
		}
		fmt.Println(r.Result)
	}

}
