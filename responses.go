package paymego

type ResponseCard struct {
	Number    string `json:"number"`
	Expire    string `json:"expire"`
	Token     string `json:"token"`
	Recurrent bool   `json:"recurrent"`
	Verify    bool   `json:"verify"`
	CardHash  string `json:"number_hash"`
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

type CardsVerifyResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Card ResponseCard `json:"card"`
	} `json:"result"`
}

type CardsCheckResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Card ResponseCard `json:"card"`
	} `json:"result"`
}

type CardsRemoveResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Success bool `json:"success"`
	} `json:"result"`
}

type CreateReceiptResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  struct {
		Receipt Receipt `json:"receipt"`
	} `json:"result"`
}

type Merchant struct {
	Id           string `json:"_id"`
	Name         string `json:"name"`
	Organization string `json:"organization"`
	Address      string `json:"address"`
	Epos         struct {
		MerchantId string `json:"merchantId"`
		TerminalId string `json:"terminalId"`
	} `json:"epos"`
	Date  int64       `json:"date"`
	Logo  interface{} `json:"logo"`
	Type  string      `json:"type"`
	Terms interface{} `json:"terms"`
}

type ResponseAccount struct {
	Name  string `json:"name"`
	Title string `json:"title"`
	Value string `json:"value"`
}

type Receipt struct {
	Id          string            `json:"_id"`
	CreateTime  int64             `json:"create_time"`
	PayTime     int               `json:"pay_time"`
	CancelTime  int               `json:"cancel_time"`
	State       int               `json:"state"`
	Type        int               `json:"type"`
	External    bool              `json:"external"`
	Operation   int               `json:"operation"`
	Category    interface{}       `json:"category"`
	Error       interface{}       `json:"error"`
	Description string            `json:"description"`
	Detail      interface{}       `json:"detail"`
	Amount      int               `json:"amount"`
	Commission  int               `json:"commission"`
	Account     []ResponseAccount `json:"account"`
	Card        ResponseCard      `json:"card"`
	Merchant    Merchant          `json:"merchant"`
	Meta        interface{}       `json:"meta"`
}

type PayReceiptResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	Id      interface{} `json:"id"`
	Result  struct {
		Receipt Receipt `json:"receipt"`
	} `json:"result"`
}
type CheckReceiptResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  struct {
		State int `json:"state"`
	} `json:"result"`
}

type CardsCheckBalanceResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  struct {
		Success bool `json:"success"`
	} `json:"result"`
}

type CreateReceiptP2PResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  struct {
		Receipt struct {
			Id         string `json:"_id"`
			CreateTime int64  `json:"create_time"`
			PayTime    int    `json:"pay_time"`
			CancelTime int    `json:"cancel_time"`
			State      int    `json:"state"`
			Type       int    `json:"type"`
			External   bool   `json:"external"`
			Operation  int    `json:"operation"`
			Category   struct {
				Id        string `json:"_id"`
				Title     string `json:"title"`
				Color     string `json:"color"`
				Sort      int    `json:"sort"`
				Operation int    `json:"operation"`
				Indoor    bool   `json:"indoor"`
			} `json:"category"`
			Error       interface{} `json:"error"`
			Description string      `json:"description"`
			Detail      interface{} `json:"detail"`
			Amount      int         `json:"amount"`
			Currency    int         `json:"currency"`
			Commission  int         `json:"commission"`
			Account     []struct {
				Name  string `json:"name"`
				Title string `json:"title"`
				Value string `json:"value"`
			} `json:"account"`
			Card     interface{} `json:"card"`
			Merchant struct {
				Id           string `json:"_id"`
				Name         string `json:"name"`
				Organization string `json:"organization"`
				Address      string `json:"address"`
				Epos         struct {
					MerchantId string `json:"merchantId"`
					TerminalId string `json:"terminalId"`
				} `json:"epos"`
				Logo  string `json:"logo"`
				Type  string `json:"type"`
				Terms struct {
					Ru string `json:"ru"`
					Uz string `json:"uz"`
				} `json:"terms"`
			} `json:"merchant"`
			Meta struct {
				Source string `json:"source"`
				Owner  string `json:"owner"`
			} `json:"meta"`
			ProcessingId interface{} `json:"processing_id"`
		} `json:"receipt"`
	} `json:"result"`
}
