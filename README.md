# AILP
使用golang集成国内开源ai的平台 Go SDK

目前已有以下平台：
- [阿里通义千问](https://help.aliyun.com/zh/dashscope/developer-reference/quick-start)
- [百度千帆](https://cloud.baidu.com/doc/WENXINWORKSHOP/s/xlmokikxe)
- [腾讯混元](https://cloud.tencent.com/document/product/1729/105701)
- [讯飞星火](https://www.xfyun.cn/doc/spark/Web.html)

各平台的授权认证都已处理，直接调用即可
注意除讯飞星火外的流式接口是websocket，其余的均是sse

## 如何使用

首先可以通过如下命令安装 SDK：

```
go get github.com/luo-hong/ailp
```

之后就可以在代码中通过如下方式引入 SDK：

```
import (
	"github.com/luo-hong/ailp"
)
```

> 我们提供了一些 [示例](./example)，可以帮助快速了解 SDK 的使用方法并完成常见功能。

### 通义千问对话

```go
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
```

### 通义千问流式

```go
    // 初始化QWEN聊天机器人客户端，使用您的API密钥
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
```
