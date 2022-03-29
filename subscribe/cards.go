package subscribe

import (
	"context"
	"net/http"
)

type ICardsRepository interface {
	GetVerifyCode(
		ctx context.Context,
		requestID,
		token string,
	) (*CardsGetVerifyCodeResponse, error)
	Create(
		ctx context.Context,
		requestID,
		cardNumber,
		cardExpiration string,
		save bool,
	) (*CardsCreateResponse, error)
	Verify(
		ctx context.Context,
		requestID,
		code,
		token string,
	) (*CardsVerifyResponse, error)
}

type cardsRepository struct {
	subscribe SubscribeAPI
}

func newCardsRepository(subscribe SubscribeAPI) cardsRepository {
	return cardsRepository{subscribe: subscribe}
}

func (c cardsRepository) GetVerifyCode(
	ctx context.Context,
	requestID,
	token string,
) (*CardsGetVerifyCodeResponse, error) {
	request, err := newCardsGetVerifyCodeRequest(requestID, token)
	if err != nil {
		return nil, err
	}
	response := CardsGetVerifyCodeResponse{}
	err = c.subscribe.do(ctx, request, http.MethodPost, &response, false)
	return &response, nil

}

func (c cardsRepository) Create(
	ctx context.Context,
	requestID,
	cardNumber,
	cardExpiration string,
	save bool,
) (*CardsCreateResponse, error) {
	request, err := newCardsCreateRequest(requestID, cardNumber, cardExpiration, save)
	if err != nil {
		return nil, err
	}
	response := CardsCreateResponse{}
	err = c.subscribe.do(ctx, request, http.MethodPost, &response, false)
	if err != nil {
		return nil, err
	}

	return &response, err
}

func (c cardsRepository) Verify(
	ctx context.Context,
	requestID,
	code,
	token string,
) (*CardsVerifyResponse, error) {
	request, err := newCardsVerifyRequest(requestID, token, code)
	if err != nil {
		return nil, err
	}
	response := CardsVerifyResponse{}
	err = c.subscribe.do(ctx, request, http.MethodPost, &response, false)
	if err != nil {
		return nil, err
	}

	return &response, err
}
