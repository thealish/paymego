package subscribe

type ResponseCard struct {
	Number    string `json:"number"`
	Expire    string `json:"expire"`
	Token     string `json:"token"`
	Recurrent bool   `json:"recurrent"`
	Verify    bool   `json:"verify"`
}

type CardsCreateResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		Card ResponseCard `json:"card"`
	} `json:"result"`
}

type CardsGetVerifyCodeResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Sent  bool   `json:"sent"`
		Phone string `json:"phone"`
		Wait  int    `json:"wait"`
	} `json:"result"`
}
