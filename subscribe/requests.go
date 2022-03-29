package subscribe

import "encoding/json"

type requestParams struct {
	RequestCard requestCard `json:"card"`
	Save        bool        `json:"save"`
}

type requestCard struct {
	Number string `json:"number"`
	Expire string `json:"expire"`
}

type cardsCreateRequest struct {
	RequestID string        `json:"id"`
	Method    string        `json:"method"`
	Params    requestParams `json:"params"`
}

type cardsGetVerifyCodeRequest struct {
	RequestID string            `json:"request_id"`
	Method    string            `json:"method"`
	Params    map[string]string `json:"params"`
}

type cardsVerifyRequest struct {
	RequestID string            `json:"request_id"`
	Method    string            `json:"method"`
	Params    map[string]string `json:"params"`
}

type CardsVerifyResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Card ResponseCard `json:"card"`
	} `json:"result"`
}

func newCardsGetVerifyCodeRequest(requestID, token string) ([]byte, error) {
	return json.Marshal(cardsGetVerifyCodeRequest{
		RequestID: requestID,
		Method:    cardsGetVerifyCode,
		Params:    map[string]string{"token": token},
	})
}

func newCardsCreateRequest(requestID, cardNumber, cardExpiration string, save bool) ([]byte, error) {
	return json.Marshal(cardsCreateRequest{
		RequestID: requestID,
		Method:    cardsCreate,
		Params: requestParams{
			RequestCard: requestCard{
				Number: cardNumber,
				Expire: cardExpiration,
			},
			Save: save,
		},
	})
}

func newCardsVerifyRequest(requestID, token, code string) ([]byte, error) {
	return json.Marshal(cardsVerifyRequest{
		RequestID: requestID,
		Method:    cardsVerify,
		Params: map[string]string{
			"token": token,
			"code":  code,
		},
	})
}
