package subscribe

import "fmt"

// ErrorResponse type represents error that Payme returns
type ErrorResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Err     struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
		Data    string `json:"data"`
	} `json:"error"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("request finished with error, message - %s code %d", e.Err.Message, e.Err.Code)
}
