package quickpay

type Pong struct {
	Msg     string      `json:"msg"`
	Scope   string      `json:"scope"`
	Version string      `json:"version"`
	Params  interface{} `json:"params"`
}
