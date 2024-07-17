package hunyuan_test

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin/internal/channel/hunyuan"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
)

func TestRun(t *testing.T) {
	credential := common.NewCredential("TENCENTCLOUD_SECRET_ID", "TENCENTCLOUD_SECRET_KEY")

	cpf := profile.NewClientProfile()

	client, _ := hunyuan.NewClient(credential, regions.Guangzhou, cpf)

	request := hunyuan.NewChatCompletionsRequest()

	msg := &hunyuan.Message{
		Role:    common.StringPtr("user"),
		Content: common.StringPtr("你好, 可以讲个笑话吗"),
	}
	request.Messages = append(request.Messages, msg)

	// hunyuan 同时支持 stream 和非 stream 的情况
	request.Stream = common.BoolPtr(true)
	model := "hunyuan-pro"
	request.Model = &model

	response, err := client.ChatCompletions(request)

	// 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}

	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
	}

	if *request.Stream {
		// stream 示例
		for event := range response.Events {
			fmt.Println(string(event.Data))
		}
	} else {
		// 非 stream 示例
		// 通过 Stream=false 参数来指定非 stream 协议, 一次性拿到结果
		fmt.Println(response.ToJsonString())
	}
}
