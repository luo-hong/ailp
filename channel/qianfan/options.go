package qianfan

import "context"

type Option func(*Options)
type Options struct {
	Model                 *string
	Endpoint              *string
	LLMRetryCount         int
	LLMRetryTimeout       float32
	LLMRetryBackoffFactor float32
	Context               context.Context
}

// 用于模型类对象设置使用的模型
func WithModel(model string) Option {
	return func(options *Options) {
		options.Model = &model
	}
}

// 用于模型类对象设置使用的 endpoint
func WithEndpoint(endpoint string) Option {
	return func(options *Options) {
		options.Endpoint = &endpoint
	}
}

// 设置重试次数
func WithLLMRetryCount(count int) Option {
	return func(options *Options) {
		options.LLMRetryCount = count
	}
}

// 设置重试超时时间
func WithLLMRetryTimeout(timeout float32) Option {
	return func(options *Options) {
		options.LLMRetryTimeout = timeout
	}
}

// 设置重试退避因子
func WithLLMRetryBackoffFactor(factor float32) Option {
	return func(options *Options) {
		options.LLMRetryBackoffFactor = factor
	}
}

// 设置初始化时的 Context
func WithContext(ctx context.Context) Option {
	return func(options *Options) {
		options.Context = ctx
	}
}

// 将多个 Option 转换成最终的 Options 对象
func makeOptions(options ...Option) *Options {
	option := Options{
		LLMRetryCount:         GetConfig().LLMRetryCount,
		LLMRetryTimeout:       GetConfig().LLMRetryTimeout,
		LLMRetryBackoffFactor: GetConfig().LLMRetryBackoffFactor,
		Context:               context.TODO(),
	}
	for _, opt := range options {
		opt(&option)
	}
	return &option
}
