package paymego

import "encoding/json"

type baseRequest struct {
	RequestID string            `json:"request_id"`
	Method    string            `json:"method"`
	Params    map[string]string `json:"params"`
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
		Params:    map[string]string{"token": token},
		Method:    cardsRemove,
	}})
}

func newCardsCheckRequest(requestID, token string) ([]byte, error) {
	return json.Marshal(cardsCheckRequest{baseRequest{
		RequestID: requestID,
		Params:    map[string]string{"token": token},
		Method:    cardCheck,
	}})
}

func newCardsGetVerifyCodeRequest(requestID, token string) ([]byte, error) {
	return json.Marshal(cardsGetVerifyCodeRequest{
		baseRequest{
			RequestID: requestID,
			Method:    cardsGetVerifyCode,
			Params:    map[string]string{"token": token},
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