package paymego

import (
	"context"
	"net/http"
)

func (s SubscribeAPI) CreateReceipt(
	ctx context.Context,
	requestID, description, detail string,
	amount int,
	account Account,
) (CreateReceiptResponse, error) {
	request, err := newCreateReceiptRequest(requestID, description, detail, amount, account)
	if err != nil {
		return CreateReceiptResponse{}, err
	}
	var response CreateReceiptResponse
	if err = s.do(ctx, request, http.MethodPost, &response, true); err != nil {
		return CreateReceiptResponse{}, err
	}
	return response, nil
}

// PayReceipt method used to pay cheque(receipt).
// payer is optional argument it can be empty struct
func (s SubscribeAPI) PayReceipt(
	ctx context.Context, requestID, cardToken, receiptID string,
	payer Payer,
) (PayReceiptResponse, error) {
	request, err := newPayReceiptRequest(requestID, cardToken, receiptID, payer)
	if err != nil {
		return PayReceiptResponse{}, err
	}
	var response PayReceiptResponse
	err = s.do(ctx, request, http.MethodPost, &response, true)
	return response, err
}
