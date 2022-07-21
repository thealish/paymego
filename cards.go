package paymego

import (
	"context"
	"net/http"
)

func (s SubscribeAPI) CardsGetVerifyCode(
	ctx context.Context,
	requestID,
	token string,
) (*CardsGetVerifyCodeResponse, error) {
	request, err := newCardsGetVerifyCodeRequest(requestID, token)
	if err != nil {
		return nil, err
	}
	response := CardsGetVerifyCodeResponse{}
	err = s.do(ctx, request, http.MethodPost, &response, false)
	return &response, nil
}

func (s SubscribeAPI) CardsCreate(
	ctx context.Context,
	requestID string,
	save bool,
	cardNumber,
	cardExpiration string,
) (*CardsCreateResponse, error) {
	request, err := newCardsCreateRequest(requestID, cardNumber, cardExpiration, save)
	if err != nil {
		return nil, err
	}
	response := CardsCreateResponse{}
	err = s.do(ctx, request, http.MethodPost, &response, false)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (s SubscribeAPI) CardsVerify(
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
	err = s.do(ctx, request, http.MethodPost, &response, false)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (s SubscribeAPI) CardsCheck(
	ctx context.Context,
	requestID,
	token string,
) (*CardsCheckResponse, error) {
	request, err := newCardsCheckRequest(requestID, token)
	if err != nil {
		return nil, err
	}
	response := CardsCheckResponse{}
	err = s.do(ctx, request, http.MethodPost, &response, true)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (s SubscribeAPI) CardsRemove(
	ctx context.Context,
	requestID,
	token string,
) (*CardsRemoveResponse, error) {
	request, err := newCardsRemoveRequest(requestID, token)
	if err != nil {
		return nil, err
	}
	response := CardsRemoveResponse{}
	err = s.do(ctx, request, http.MethodPost, &response, true)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (s SubscribeAPI) CardsCheckBalance(
	ctx context.Context,
	requestID, token string,
	amount int,
) (CardsCheckBalanceResponse, error) {
	request, err := newCardsCheckBalanceRequest(requestID, token, amount)
	if err != nil {
		return CardsCheckBalanceResponse{}, err
	}
	var response CardsCheckBalanceResponse
	err = s.do(ctx, request, http.MethodPost, &response, true)
	if err != nil {
		return CardsCheckBalanceResponse{}, err
	}
	return response, err
}
