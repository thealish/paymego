package subscribe

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/textproto"
)

const (
	create             = "receipts.create"
	pay                = "receipts.pay"
	createp2p          = "receipts.p2p"
	cardCheck          = "cards.check"
	cardCheckBalance   = "cards.checkBalance"
	cardsCreate        = "cards.create"
	cardsGetVerifyCode = "cards.get_verify_code"
	cardsVerify        = "cards.verify"
	cardsRemove        = "cards.remove"
	receiptsCheck      = "receipts.check"
	receiptsCancel     = "receipts.cancel"
)

var (
	ErrEmptyOrInvalidPaycomID  = errors.New("invalid paycomID")
	ErrEmptyOrInvalidPaycomKey = errors.New("invalid paycomKey")
)

type SubscribeAPI struct {
	headers    xAuthHeaders
	baseURL    string
	Cards      ICardsRepository
	httpClient http.Client
	logger     *log.Logger
}

type SubsribeAPIOpts struct {
	PaycomID   string
	PaycomKey  string
	BaseURL    string
	Logger     *log.Logger
	httpClient http.Client
}

// New returns new instance of SubscribeAPI
func New(args *SubsribeAPIOpts) (SubscribeAPI, error) {
	err := args.validate()
	if err != nil {
		return SubscribeAPI{}, err
	}
	subscribeAPI := SubscribeAPI{
		httpClient: args.httpClient,
		baseURL:    args.BaseURL,
		logger:     args.Logger,
		headers:    getXAuthHeaders(args.PaycomID, args.PaycomKey),
	}

	subscribeAPI.Cards = newCardsRepository(subscribeAPI)

	return subscribeAPI, nil
}

func (s *SubsribeAPIOpts) getHeaders() map[string]string {
	return map[string]string{
		"X-Auth":       fmt.Sprintf("%s:%s", s.PaycomID, s.PaycomKey),
		"Content-Type": "application/json",
	}
}

func (s *SubscribeAPI) do(
	ctx context.Context,
	requestBody []byte,
	method string,
	response interface{},
	withPaycomKey bool,
) error {
	request, err := http.NewRequestWithContext(ctx, method, s.baseURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	if withPaycomKey {
		request.Header = s.headers.withPaycomKey()
	} else {
		request.Header = s.headers.withoutPaycomKey()
	}
	res, err := s.httpClient.Do(request)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	errResponse, err := hasError(body)

	if errResponse != nil || err != nil {
		return errResponse
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}
	return nil
}

func hasError(r []byte) (*ErrorResponse, error) {
	var errResponse *ErrorResponse
	err := json.Unmarshal(r, &errResponse)
	if err != nil {
		return nil, err
	}
	if errResponse.Err.Code != 0 {
		return errResponse, nil
	}
	return nil, nil
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

type xAuthHeaders struct {
	paycomID  string
	paycomKey string
}

func (x xAuthHeaders) withPaycomKey() http.Header {
	headers := http.Header{}
	textproto.MIMEHeader(headers).Add("X-Auth", fmt.Sprintf("%s:%s", x.paycomID, x.paycomKey))
	textproto.MIMEHeader(headers).Add("Content-Type", "application/json")
	return headers
}

func (x xAuthHeaders) withoutPaycomKey() http.Header {
	headers := http.Header{}
	textproto.MIMEHeader(headers).Add("X-Auth", x.paycomID)
	textproto.MIMEHeader(headers).Add("Content-Type", "application/json")
	return headers
}

func getXAuthHeaders(paycomID, paycomKey string) xAuthHeaders {
	return xAuthHeaders{paycomID: paycomID, paycomKey: paycomKey}
}
