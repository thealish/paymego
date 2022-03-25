package subscribe_api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	create           = "receipts.create"
	pay              = "receipts.pay"
	createp2p        = "receipts.p2p"
	cardCheck        = "cards.check"
	cardCheckBalance = "cards.checkBalance"
	cardsCreate      = "cards.create"
	cardsRemove      = "cards.remove"
	receiptsCheck    = "receipts.check"
	receiptsCancel   = "receipts.cancel"
)

var (
	ErrEmptyOrInvalidPaycomID  = errors.New("invalid paycomID")
	ErrEmptyOrInvalidPaycomKey = errors.New("invalid paycomKey")
)

type SubscribeAPI struct {
	paycomID   string
	paycomKey  string
	headers    map[string]string
	baseURL    string
	httpClient http.Client
}

type SubsribeAPIOpts struct {
	PaycomID   string
	PaycomKey  string
	BaseURL    string
	httpClient http.Client
}

func (s SubscribeAPI) CardsCreate(
	ctx context.Context,
	requestID,
	cardNumber,
	cardExpiration string,
	save bool,
) (*CardsCreateResponse, error) {
	request, err := NewCardsCreateRequest(requestID, cardNumber, cardExpiration, save)
	if err != nil {
		return nil, err
	}
	res, err := do(ctx, s.httpClient, request, http.MethodPost, s.baseURL)
	response, ok := res.(CardsCreateResponse)
	if !ok {
		return nil, fmt.Errorf("unable to convert %s to CardsCreateResponse, err - %s", res, err)
	}
	return &response, err
}

// New returns new instance of SubscribeAPI
func New(args *SubsribeAPIOpts) (SubscribeAPI, error) {
	err := args.validate()
	if err != nil {
		return SubscribeAPI{}, err
	}
	subscribeAPI := SubscribeAPI{
		paycomID:   args.PaycomID,
		paycomKey:  args.PaycomKey,
		httpClient: args.httpClient,
		baseURL:    args.BaseURL,
	}
	subscribeAPI.setHeaders()
	return subscribeAPI, nil
}

func (s *SubscribeAPI) setHeaders() {
	s.headers = map[string]string{
		"X-Auth":       fmt.Sprintf("%s:%s", s.paycomID, s.paycomKey),
		"Content-Type": "application/json",
	}
}

func do(ctx context.Context, c http.Client, requestBody []byte, method, url string) (interface{}, error) {

	request, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	res, err := c.Do(request)
	if err != nil {
		return nil, err
	}
	var response interface{}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response, nil

}

func (s SubsribeAPIOpts) validate() error {
	if s.PaycomID == "" {
		return ErrEmptyOrInvalidPaycomID
	}
	if s.PaycomKey == "" {
		return ErrEmptyOrInvalidPaycomKey
	}
	return nil
}
