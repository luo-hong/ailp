package qwen

const (
	ChatUser   = "user"
	ChatSystem = "system"
	ChatBot    = "assistant"
)

const (
	ModelQWenTurbo          = "qwen-turbo"
	ModelQWenPlus           = "qwen-plus"
	ModelQWenMax            = "qwen-max"
	ModelQWenMax0428        = "qwen-max-0428"
	ModelQWenMax0403        = "qwen-max-0403"
	ModelQWenMax0107        = "qwen-max-0107"
	ModelQWenMaxLongContext = "qwen-max-longcontext"
	ModelBaseUrl            = "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation"
)

type Input struct {
	Messages []Messages `json:"messages"`
}

type QWenTurbo struct {
	Model      string     `json:"model"`
	Input      Input      `json:"input"`
	Parameters Parameters `json:"parameters"`
}

type Parameters struct {
	EnableSearch      bool   `json:"enable_search"`
	IncrementalOutput bool   `json:"incremental_output"`
	ResponseFormat    string `json:"response_format"`
}

type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Result

type Output struct {
	Text         string `json:"text"`
	FinishReason string `json:"finish_reason"`
}

type Usage struct {
	OutputTokens int `json:"output_tokens"`
	InputTokens  int `json:"input_tokens"`
}

type Response struct {
	Output    Output `json:"output"`
	Usage     Usage  `json:"usage"`
	RequestID string `json:"request_id"`
}

type ResponseError struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	RequestId string `json:"request_id"`
}
