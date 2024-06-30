package qwen

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/alevinval/sse/pkg/base"
	"github.com/alevinval/sse/pkg/decoder"
)

type QWenChat struct {
	BaseUrl   string
	ApiKey    string
	QWenModel string
	Params    Parameters
}

func NewWithDefaultChat(apiKey string) *QWenChat {
	return &QWenChat{
		BaseUrl:   ModelBaseUrl,
		ApiKey:    apiKey,
		QWenModel: ModelQWenTurbo,
		Params:    Parameters{EnableSearch: true, ResponseFormat: "message"},
	}
}

// GetAIReply 获取聊天回复
func (c *QWenChat) GetAIReply(messages []Messages) (Response, error) {
	client := http.Client{}

	if !checkParams(c) {
		return Response{}, errors.New("invalid parameters")
	}
	// body
	body := QWenTurbo{
		Model:      c.QWenModel,
		Input:      Input{Messages: messages},
		Parameters: c.Params,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return Response{}, fmt.Errorf("json.Marshal failed,err:%v", err)
	}
	// 创建请求
	req, err := http.NewRequest("POST", c.BaseUrl, bytes.NewReader(jsonBody))
	if err != nil {
		return Response{}, fmt.Errorf("http.NewRequest failed,err:%v", err)
	}
	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.ApiKey)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return Response{}, fmt.Errorf("http.Client.Do failed,err:%v", err)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{}, fmt.Errorf("io.ReadAll failed,err:%v", err)
	}
	defer resp.Body.Close()
	// 读取响应
	var result Response

	if resp.StatusCode != http.StatusOK {
		var errResp ResponseError
		err = json.Unmarshal(b, &errResp)
		if err != nil {
			return Response{}, err
		}
		return Response{}, fmt.Errorf("failed,err:%v,code:%s,message:%s", err, errResp.Code, errResp.Message)
	}

	err = json.Unmarshal(b, &result)

	if err != nil {
		fmt.Printf("json.NewDecoder failed,err:%v\n", err)
	}

	return result, nil
}

// GetAIReplyStream 获取聊天回复
func (c *QWenChat) GetAIReplyStream(messages []Messages, fn func(e *base.MessageEvent) error) error {
	client := http.Client{}

	if !checkParams(c) {
		return fmt.Errorf("invalid parameters")
	}

	// Prepare request body
	body := QWenTurbo{
		Model:      c.QWenModel,
		Input:      Input{Messages: messages},
		Parameters: c.Params,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	// Create request
	req, err := http.NewRequest("POST", c.BaseUrl, bytes.NewReader(jsonBody))
	if err != nil {
		return fmt.Errorf("new request failed: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-DashScope-SSE", "enable")
	req.Header.Set("Authorization", "Bearer "+c.ApiKey)

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("code:%d", resp.StatusCode)
	}
	// 使用decoder解析服务端推送的SSE事件。
	code := decoder.New(resp.Body)
	for {
		// 解码SSE事件，如果解码失败则根据错误类型处理。
		event, err := code.Decode()
		if err != nil {
			if err == io.EOF {
				// 如果遇到EOF错误，表示连接断开，正常结束函数。
				return nil
			}
			// 其他错误直接返回。
			return err
		}
		// 调用回调函数处理解码成功的事件，如果处理失败则返回错误。
		if err := fn(event); err != nil {
			return err
		}
	}
}

// 效验参数
func checkParams(chat *QWenChat) bool {
	if chat.QWenModel == "" {
		log.Fatal("QWenModel is empty")
		return false
	}
	if chat.ApiKey == "" {
		log.Fatal("ApiKey is empty")
		return false
	}
	if chat.BaseUrl == "" {
		log.Fatal("BaseUrl is empty")
		return false
	}
	return true
}
