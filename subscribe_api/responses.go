package subscribe_api

type ResponseCard struct {
	Number    string `json:"number"`
	Expire    string `json:"expire"`
	Token     string `json:"token"`
	Recurrent bool   `json:"recurrent"`
	Verify    bool   `json:"verify"`
}

type CardsCreateResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Card ResponseCard `json:"card"`
	} `json:"result"`
}
