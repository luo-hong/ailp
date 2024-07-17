package spark

const (
	SparkProModel = "generalv3"
	SparkProUrl   = "wss://spark-api.xf-yun.com/v3.1/chat"
)

type SparkReq struct {
	Header    Header    `json:"header"`
	Parameter Parameter `json:"parameter"`
	Payload   Payload   `json:"payload"`
}

type Header struct {
	Appid string `json:"app_id"`
	Uid   string `json:"uid"`
}

type Parameter struct {
	Chat Chat `json:"chat"`
}

type Chat struct {
	Domain      string  `json:"domain"`
	Temperature float64 `json:"temperature"`
	TopK        int64   `json:"top_k"`
	MaxToken    int64   `json:"max_tokens"`
	Auditing    string  `json:"auditing"`
}

type Payload struct {
	Message Message `json:"message"`
}

type Message struct {
	Text []Text `json:"text"`
}

type Text struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
