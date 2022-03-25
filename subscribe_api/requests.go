package subscribe_api

import "encoding/json"

type RequestParams struct {
	RequestCard RequestCard `json:"card"`
	Save        bool        `json:"save"`
}

type RequestCard struct {
	Number string `json:"number"`
	Expire string `json:"expire"`
}

type CardsCreateRequest struct {
	RequestID string        `json:"id"`
	Method    string        `json:"method"`
	Params    RequestParams `json:"params"`
}

func NewCardsCreateRequest(requestID, cardNumber, cardExpiration string, save bool) ([]byte, error) {
	return json.Marshal(CardsCreateRequest{
		RequestID: requestID,
		Method:    cardsCreate,
		Params: RequestParams{
			RequestCard: RequestCard{
				Number: cardNumber,
				Expire: cardExpiration,
			},
			Save: save,
		},
	})
}
