package spark

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type SparkChat struct {
	HostUrl    string
	Appid      string
	ApiKey     string
	ApiSecret  string
	SparkModel string
}

func NewWithDefaultSparkChat(hosturl, appid, apiKey, apiSecret, sparkmodel string) *SparkChat {
	return &SparkChat{
		HostUrl:    hosturl,
		Appid:      appid,
		ApiKey:     apiKey,
		ApiSecret:  apiSecret,
		SparkModel: sparkmodel,
	}
}

// 生成参数
func (s *SparkChat) SparkRequest(uid, question string) SparkReq {
	messages := []Text{
		{Role: "user", Content: question},
	}

	data := SparkReq{
		Header: Header{
			Appid: s.Appid,
			Uid:   uid,
		},
		Parameter: Parameter{
			Chat: Chat{
				Domain:      s.SparkModel,
				Temperature: 0.8,
				TopK:        6,
				MaxToken:    2048,
				Auditing:    "default",
			},
		},
		Payload: Payload{
			Message: Message{
				Text: messages,
			},
		},
	}

	return data
}

// 创建鉴权url  apikey 即 hmac username
func (s *SparkChat) AssembleAuthUrl1() string {
	ul, err := url.Parse(s.HostUrl)
	if err != nil {
		fmt.Println(err)
	}
	//签名时间
	date := time.Now().UTC().Format(time.RFC1123)
	//date = "Tue, 28 May 2019 09:10:42 MST"
	//参与签名的字段 host ,date, request-line
	signString := []string{"host: " + ul.Host, "date: " + date, "GET " + ul.Path + " HTTP/1.1"}
	//拼接签名字符串
	sgin := strings.Join(signString, "\n")
	// fmt.Println(sgin)
	//签名结果
	sha := HmacWithShaTobase64("hmac-sha256", sgin, s.ApiSecret)
	// fmt.Println(sha)
	//构建请求参数 此时不需要urlencoding
	authUrl := fmt.Sprintf("hmac username=\"%s\", algorithm=\"%s\", headers=\"%s\", signature=\"%s\"", s.ApiKey,
		"hmac-sha256", "host date request-line", sha)
	//将请求参数使用base64编码
	authorization := base64.StdEncoding.EncodeToString([]byte(authUrl))

	v := url.Values{}
	v.Add("host", ul.Host)
	v.Add("date", date)
	v.Add("authorization", authorization)
	//将编码后的字符串url encode后添加到url后面
	callurl := s.HostUrl + "?" + v.Encode()
	return callurl
}

func HmacWithShaTobase64(algorithm, data, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	encodeData := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(encodeData)
}

func ReadResp(resp *http.Response) string {
	if resp == nil {
		return ""
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("code=%d,body=%s", resp.StatusCode, string(b))
}
