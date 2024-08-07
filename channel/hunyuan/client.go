package hunyuan

import (
	"context"
	"errors"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"

)

const APIVersion = "2023-09-01"

type Client struct {
	common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
	cpf := profile.NewClientProfile()
	client = &Client{}
	client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
	return
}

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
	client = &Client{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

func NewChatCompletionsRequest() (request *ChatCompletionsRequest) {
	request = &ChatCompletionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("hunyuan", APIVersion, "ChatCompletions")

	return
}

func NewChatCompletionsResponse() (response *ChatCompletionsResponse) {
	response = &ChatCompletionsResponse{}
	return

}

// ChatCompletions
// 腾讯混元大模型是由腾讯研发的大语言模型，具备强大的中文创作能力，复杂语境下的逻辑推理能力，以及可靠的任务执行能力。本接口支持流式或非流式调用，当使用流式调用时为 SSE 协议。
//
//  1. 本接口暂不支持返回图片内容。
//
//  2. 默认每种模型单账号限制并发数为 5 路，如您有提高并发限制的需求请 [联系我们](https://cloud.tencent.com/act/event/Online_service) 。
//
//  3. 请使用 SDK 调用本接口，每种开发语言的 SDK Git 仓库 examples/hunyuan/v20230901/ 目录下有提供示例供参考。SDK 链接在文档下方 “**开发者资源 - SDK**” 部分提供。
//
//  4. 我们推荐您使用 API Explorer，方便快速地在线调试接口和下载各语言的示例代码，[点击打开](https://console.cloud.tencent.com/api/explorer?Product=hunyuan&Version=2023-09-01&Action=ChatCompletions)。
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_ENGINEREQUESTTIMEOUT = "FailedOperation.EngineRequestTimeout"
//	FAILEDOPERATION_ENGINESERVERERROR = "FailedOperation.EngineServerError"
//	FAILEDOPERATION_ENGINESERVERLIMITEXCEEDED = "FailedOperation.EngineServerLimitExceeded"
//	FAILEDOPERATION_FREERESOURCEPACKEXHAUSTED = "FailedOperation.FreeResourcePackExhausted"
//	FAILEDOPERATION_RESOURCEPACKEXHAUSTED = "FailedOperation.ResourcePackExhausted"
//	FAILEDOPERATION_SERVICENOTACTIVATED = "FailedOperation.ServiceNotActivated"
//	FAILEDOPERATION_SERVICESTOP = "FailedOperation.ServiceStop"
//	FAILEDOPERATION_SERVICESTOPARREARS = "FailedOperation.ServiceStopArrears"
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER = "InvalidParameter"
//	INVALIDPARAMETERVALUE = "InvalidParameterValue"
//	INVALIDPARAMETERVALUE_MODEL = "InvalidParameterValue.Model"
//	LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ChatCompletions(request *ChatCompletionsRequest) (response *ChatCompletionsResponse, err error) {
	return c.ChatCompletionsWithContext(context.Background(), request)
}

// ChatCompletions
// 腾讯混元大模型是由腾讯研发的大语言模型，具备强大的中文创作能力，复杂语境下的逻辑推理能力，以及可靠的任务执行能力。本接口支持流式或非流式调用，当使用流式调用时为 SSE 协议。
//
//  1. 本接口暂不支持返回图片内容。
//
//  2. 默认每种模型单账号限制并发数为 5 路，如您有提高并发限制的需求请 [联系我们](https://cloud.tencent.com/act/event/Online_service) 。
//
//  3. 请使用 SDK 调用本接口，每种开发语言的 SDK Git 仓库 examples/hunyuan/v20230901/ 目录下有提供示例供参考。SDK 链接在文档下方 “**开发者资源 - SDK**” 部分提供。
//
//  4. 我们推荐您使用 API Explorer，方便快速地在线调试接口和下载各语言的示例代码，[点击打开](https://console.cloud.tencent.com/api/explorer?Product=hunyuan&Version=2023-09-01&Action=ChatCompletions)。
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_ENGINEREQUESTTIMEOUT = "FailedOperation.EngineRequestTimeout"
//	FAILEDOPERATION_ENGINESERVERERROR = "FailedOperation.EngineServerError"
//	FAILEDOPERATION_ENGINESERVERLIMITEXCEEDED = "FailedOperation.EngineServerLimitExceeded"
//	FAILEDOPERATION_FREERESOURCEPACKEXHAUSTED = "FailedOperation.FreeResourcePackExhausted"
//	FAILEDOPERATION_RESOURCEPACKEXHAUSTED = "FailedOperation.ResourcePackExhausted"
//	FAILEDOPERATION_SERVICENOTACTIVATED = "FailedOperation.ServiceNotActivated"
//	FAILEDOPERATION_SERVICESTOP = "FailedOperation.ServiceStop"
//	FAILEDOPERATION_SERVICESTOPARREARS = "FailedOperation.ServiceStopArrears"
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER = "InvalidParameter"
//	INVALIDPARAMETERVALUE = "InvalidParameterValue"
//	INVALIDPARAMETERVALUE_MODEL = "InvalidParameterValue.Model"
//	LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ChatCompletionsWithContext(ctx context.Context, request *ChatCompletionsRequest) (response *ChatCompletionsResponse, err error) {
	if request == nil {
		request = NewChatCompletionsRequest()
	}

	if c.GetCredential() == nil {
		return nil, errors.New("ChatCompletions require credential")
	}

	request.SetContext(ctx)

	response = NewChatCompletionsResponse()
	err = c.Send(request, response)
	return
}

func NewGetEmbeddingRequest() (request *GetEmbeddingRequest) {
	request = &GetEmbeddingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("hunyuan", APIVersion, "GetEmbedding")

	return
}

func NewGetEmbeddingResponse() (response *GetEmbeddingResponse) {
	response = &GetEmbeddingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// GetEmbedding
// 腾讯混元 Embedding 接口，可以将文本转化为高质量的向量数据。
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_ENGINESERVERERROR = "FailedOperation.EngineServerError"
//	FAILEDOPERATION_ENGINESERVERLIMITEXCEEDED = "FailedOperation.EngineServerLimitExceeded"
//	INVALIDPARAMETER = "InvalidParameter"
//	INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetEmbedding(request *GetEmbeddingRequest) (response *GetEmbeddingResponse, err error) {
	return c.GetEmbeddingWithContext(context.Background(), request)
}

// GetEmbedding
// 腾讯混元 Embedding 接口，可以将文本转化为高质量的向量数据。
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_ENGINESERVERERROR = "FailedOperation.EngineServerError"
//	FAILEDOPERATION_ENGINESERVERLIMITEXCEEDED = "FailedOperation.EngineServerLimitExceeded"
//	INVALIDPARAMETER = "InvalidParameter"
//	INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetEmbeddingWithContext(ctx context.Context, request *GetEmbeddingRequest) (response *GetEmbeddingResponse, err error) {
	if request == nil {
		request = NewGetEmbeddingRequest()
	}

	if c.GetCredential() == nil {
		return nil, errors.New("GetEmbedding require credential")
	}

	request.SetContext(ctx)

	response = NewGetEmbeddingResponse()
	err = c.Send(request, response)
	return
}

func NewGetTokenCountRequest() (request *GetTokenCountRequest) {
	request = &GetTokenCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("hunyuan", APIVersion, "GetTokenCount")

	return
}

func NewGetTokenCountResponse() (response *GetTokenCountResponse) {
	response = &GetTokenCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// GetTokenCount
// 该接口用于计算文本对应Token数、字符数。
//
// 可能返回的错误码:
//
//	FAILEDOPERATION = "FailedOperation"
func (c *Client) GetTokenCount(request *GetTokenCountRequest) (response *GetTokenCountResponse, err error) {
	return c.GetTokenCountWithContext(context.Background(), request)
}

// GetTokenCount
// 该接口用于计算文本对应Token数、字符数。
//
// 可能返回的错误码:
//
//	FAILEDOPERATION = "FailedOperation"
func (c *Client) GetTokenCountWithContext(ctx context.Context, request *GetTokenCountRequest) (response *GetTokenCountResponse, err error) {
	if request == nil {
		request = NewGetTokenCountRequest()
	}

	if c.GetCredential() == nil {
		return nil, errors.New("GetTokenCount require credential")
	}

	request.SetContext(ctx)

	response = NewGetTokenCountResponse()
	err = c.Send(request, response)
	return
}

func NewQueryHunyuanImageJobRequest() (request *QueryHunyuanImageJobRequest) {
	request = &QueryHunyuanImageJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("hunyuan", APIVersion, "QueryHunyuanImageJob")

	return
}

func NewQueryHunyuanImageJobResponse() (response *QueryHunyuanImageJobResponse) {
	response = &QueryHunyuanImageJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// QueryHunyuanImageJob
// 混元生图接口基于混元大模型，将根据输入的文本描述，智能生成与之相关的结果图。分为提交任务和查询任务2个接口。
//
// 提交任务：输入文本等，提交一个混元生图异步任务，获得任务 ID。
//
// 查询任务：根据任务 ID 查询任务的处理状态、处理结果，任务处理完成后可获得生成图像结果。
//
// 并发任务数（并发）说明：并发任务数指能同时处理的任务数量。混元生图默认提供1个并发任务数，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//
//	FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryHunyuanImageJob(request *QueryHunyuanImageJobRequest) (response *QueryHunyuanImageJobResponse, err error) {
	return c.QueryHunyuanImageJobWithContext(context.Background(), request)
}

// QueryHunyuanImageJob
// 混元生图接口基于混元大模型，将根据输入的文本描述，智能生成与之相关的结果图。分为提交任务和查询任务2个接口。
//
// 提交任务：输入文本等，提交一个混元生图异步任务，获得任务 ID。
//
// 查询任务：根据任务 ID 查询任务的处理状态、处理结果，任务处理完成后可获得生成图像结果。
//
// 并发任务数（并发）说明：并发任务数指能同时处理的任务数量。混元生图默认提供1个并发任务数，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//
//	FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryHunyuanImageJobWithContext(ctx context.Context, request *QueryHunyuanImageJobRequest) (response *QueryHunyuanImageJobResponse, err error) {
	if request == nil {
		request = NewQueryHunyuanImageJobRequest()
	}

	if c.GetCredential() == nil {
		return nil, errors.New("QueryHunyuanImageJob require credential")
	}

	request.SetContext(ctx)

	response = NewQueryHunyuanImageJobResponse()
	err = c.Send(request, response)
	return
}

func NewSubmitHunyuanImageJobRequest() (request *SubmitHunyuanImageJobRequest) {
	request = &SubmitHunyuanImageJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("hunyuan", APIVersion, "SubmitHunyuanImageJob")

	return
}

func NewSubmitHunyuanImageJobResponse() (response *SubmitHunyuanImageJobResponse) {
	response = &SubmitHunyuanImageJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// SubmitHunyuanImageJob
// 混元生图接口基于混元大模型，将根据输入的文本描述，智能生成与之相关的结果图。分为提交任务和查询任务2个接口。
//
// 提交任务：输入文本等，提交一个混元生图异步任务，获得任务 ID。
//
// 查询任务：根据任务 ID 查询任务的处理状态、处理结果，任务处理完成后可获得生成图像结果。
//
// 并发任务数（并发）说明：并发任务数指能同时处理的任务数量。混元生图默认提供1个并发任务数，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//
//	FAILEDOPERATION = "FailedOperation"
func (c *Client) SubmitHunyuanImageJob(request *SubmitHunyuanImageJobRequest) (response *SubmitHunyuanImageJobResponse, err error) {
	return c.SubmitHunyuanImageJobWithContext(context.Background(), request)
}

// SubmitHunyuanImageJob
// 混元生图接口基于混元大模型，将根据输入的文本描述，智能生成与之相关的结果图。分为提交任务和查询任务2个接口。
//
// 提交任务：输入文本等，提交一个混元生图异步任务，获得任务 ID。
//
// 查询任务：根据任务 ID 查询任务的处理状态、处理结果，任务处理完成后可获得生成图像结果。
//
// 并发任务数（并发）说明：并发任务数指能同时处理的任务数量。混元生图默认提供1个并发任务数，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//
//	FAILEDOPERATION = "FailedOperation"
func (c *Client) SubmitHunyuanImageJobWithContext(ctx context.Context, request *SubmitHunyuanImageJobRequest) (response *SubmitHunyuanImageJobResponse, err error) {
	if request == nil {
		request = NewSubmitHunyuanImageJobRequest()
	}

	if c.GetCredential() == nil {
		return nil, errors.New("SubmitHunyuanImageJob require credential")
	}

	request.SetContext(ctx)

	response = NewSubmitHunyuanImageJobResponse()
	err = c.Send(request, response)
	return
}
