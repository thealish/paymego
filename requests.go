package paymego

import "encoding/json"

type baseRequest struct {
	RequestID string                 `json:"request_id"`
	Method    string                 `json:"method"`
	Params    map[string]interface{} `json:"params"`
}

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
	baseRequest
}

type cardsVerifyRequest struct {
	RequestID string            `json:"request_id"`
	Method    string            `json:"method"`
	Params    map[string]string `json:"params"`
}

type cardsCheckRequest struct {
	baseRequest
}

type cardsRemoveRequest struct {
	baseRequest
}

func newCardsRemoveRequest(requestID, token string) ([]byte, error) {
	return json.Marshal(cardsRemoveRequest{baseRequest{
		RequestID: requestID,
		Params:    map[string]interface{}{"token": token},
		Method:    cardsRemove,
	}})
}

func newCardsCheckRequest(requestID, token string) ([]byte, error) {
	return json.Marshal(cardsCheckRequest{baseRequest{
		RequestID: requestID,
		Params:    map[string]interface{}{"token": token},
		Method:    cardCheck,
	}})
}

func newCardsGetVerifyCodeRequest(requestID, token string) ([]byte, error) {
	return json.Marshal(cardsGetVerifyCodeRequest{
		baseRequest{
			RequestID: requestID,
			Method:    cardsGetVerifyCode,
			Params:    map[string]interface{}{"token": token},
		},
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

type Account struct {
	OrderID  string `json:"order_id"`
	CardID   string `json:"card_id"`
	ReasonID string `json:"reason_id"`
}

func newCreateReceiptRequest(requestID, desc, detail string, amount int, account Account) ([]byte, error) {
	baseRequest := baseRequest{
		RequestID: requestID,
		Method:    create,
		Params:    make(map[string]interface{}),
	}
	baseRequest.Params = map[string]interface{}{
		"account":     account,
		"amount":      amount,
		"description": desc,
		"detail":      detail,
	}
	return json.Marshal(baseRequest)
}

type Payer struct {
	ID    string `json:"id"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Name  string `json:"name"`
	IP    string `json:"ip"`
}

func newPayReceiptRequest(requestID, token, ID string, payer Payer) ([]byte, error) {
	baseRequest := baseRequest{
		RequestID: requestID,
		Method:    pay,
		Params:    make(map[string]interface{}),
	}
	baseRequest.Params = map[string]interface{}{
		"token": token,
		"id":    ID,
		"payer": payer,
	}
	return json.Marshal(baseRequest)
}

func newCheckReceiptRequest(requestID, receiptID string) ([]byte, error) {
	baseRequest := baseRequest{
		RequestID: requestID,
		Method:    cardCheck,
		Params:    make(map[string]interface{}),
	}
	baseRequest.Params = map[string]interface{}{
		"id": receiptID,
	}
	return json.Marshal(baseRequest)
}

func newCardsCheckBalanceRequest(requestID, token string, amount int) ([]byte, error) {
	baseRequest := baseRequest{
		RequestID: requestID,
		Method:    cardCheckBalance,
		Params:    make(map[string]interface{}),
	}
	baseRequest.Params = map[string]interface{}{
		"token":  token,
		"amount": amount,
	}
	return json.Marshal(baseRequest)
}
