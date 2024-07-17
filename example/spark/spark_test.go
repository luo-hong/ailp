package spark_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/luo-hong/ailp/internal/channel/spark"
)

func TestRun(t *testing.T) {
	sparkclient := spark.NewWithDefaultSparkChat(
		spark.SparkProUrl,
		"78e11dbe",
		"3df4d485895f06c224da9ed6a2686251",
		"MDg5NGFiMGYyYWYxNDc1ZjA4ZjJmZWEz",
		spark.SparkProModel,
	)
	d := websocket.Dialer{
		HandshakeTimeout: 5 * time.Second,
	}
	//握手并建立websocket 连接
	conn, resp, err := d.Dial(sparkclient.AssembleAuthUrl1(), nil)
	if err != nil {
		panic(spark.ReadResp(resp) + err.Error())
	} else if resp.StatusCode != 101 {
		panic(spark.ReadResp(resp) + err.Error())
	}

	go func() {
		data := sparkclient.SparkRequest("12345", "你是谁，可以干什么？")
		conn.WriteJSON(data)
	}()

	var answer = ""
	//获取返回的数据
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read message error:", err)
			break
		}

		var data map[string]interface{}
		err1 := json.Unmarshal(msg, &data)
		if err1 != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}
		fmt.Println(string(msg))
		//解析数据
		payload := data["payload"].(map[string]interface{})
		choices := payload["choices"].(map[string]interface{})
		header := data["header"].(map[string]interface{})
		code := header["code"].(float64)

		if code != 0 {
			fmt.Println(data["payload"])
			return
		}
		status := choices["status"].(float64)
		fmt.Println(status)
		text := choices["text"].([]interface{})
		content := text[0].(map[string]interface{})["content"].(string)
		if status != 2 {
			answer += content
		} else {
			fmt.Println("收到最终结果")
			answer += content
			usage := payload["usage"].(map[string]interface{})
			temp := usage["text"].(map[string]interface{})
			totalTokens := temp["total_tokens"].(float64)
			fmt.Println("total_tokens:", totalTokens)
			conn.Close()
			break
		}

	}
	//输出返回结果
	fmt.Println(answer)

	time.Sleep(1 * time.Second)
}
