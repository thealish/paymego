package main

import (
	"context"
	"fmt"
	"github.com/thealish/paymego"
)

func main() {
	ctx := context.Background()
	s, err := paymego.NewSubscribeAPI(paymego.SubsribeAPIOpts{
		PaycomID:  "5feb5dd783c40aed047fe655",
		PaycomKey: "rwAAUFwRSFI5&eYtuq5Q7jd7u@Y6kRcRw44g",
		BaseURL:   "https://checkout.paycom.uz/api/",
	})
	if err != nil {
		panic(err)
	}
	r, err := s.CreateReceipt(ctx, "2134567879654321", "blabl", "awdaw", 2500, paymego.Account{
		OrderID:  "2383",
		CardID:   "129812",
		ReasonID: "12",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("response %+v", r)

}
