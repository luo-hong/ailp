package qianfan

// 模型请求的前缀
const (
	modelAPIPrefix = "/rpc/2.0/ai_custom/v1/wenxinworkshop"
	authAPIPrefix  = "/oauth/2.0/token"
)

// API URL
const (
	serviceListURL = "/wenxinworkshop/service/list"
)

// 默认使用的模型
const (
	DefaultChatCompletionModel = "ERNIE-Lite-8K"
	DefaultCompletionModel     = "ERNIE-Lite-8K"
	DefaultEmbeddingModel      = "Embedding-V1"
	DefaultText2ImageModel     = "Stable-Diffusion-XL"
)

// API 错误码
const (
	NoErrorErrCode                    = 0
	UnknownErrorErrCode               = 1
	ServiceUnavailableErrCode         = 2
	UnsupportedMethodErrCode          = 3
	RequestLimitReachedErrCode        = 4
	NoPermissionToAccessDataErrCode   = 6
	GetServiceTokenFailedErrCode      = 13
	AppNotExistErrCode                = 15
	DailyLimitReachedErrCode          = 17
	QPSLimitReachedErrCode            = 18
	TotalRequestLimitReachedErrCode   = 19
	InvalidRequestErrCode             = 100
	APITokenInvalidErrCode            = 110
	APITokenExpiredErrCode            = 111
	InternalErrorErrCode              = 336000
	InvalidArgumentErrCode            = 336001
	InvalidJSONErrCode                = 336002
	InvalidParamErrCode               = 336003
	PermissionErrorErrCode            = 336004
	APINameNotExistErrCode            = 336005
	ServerHighLoadErrCode             = 336100
	InvalidHTTPMethodErrCode          = 336101
	InvalidArgumentSystemErrCode      = 336104
	InvalidArgumentUserSettingErrCode = 336105

	ConsoleInternalErrorErrCode = 500000
)

const Version = "v0.0.9"
const versionIndicator = "qianfan_go_sdk_" + Version
