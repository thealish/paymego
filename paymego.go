package main

import (
	"context"
	"fmt"
	"github.com/thealish/paymego/subscribe"
	"log"
	"math/rand"
	"strconv"
)

var (
	prodURL = "https://checkout.paycom.uz/api"
	testURL = "https://checkout.test.paycom.uz/api"
)

func main() {
	s, err := subscribe.New(&subscribe.SubsribeAPIOpts{
		PaycomID:  "paycomID",
		PaycomKey: "paycomKey",
		BaseURL:   testURL,
	})
	if err != nil {
		panic(err)
	}
	r, err := s.Cards.Create(context.Background(), strconv.Itoa(rand.Int()), "8600495473316478", "03/99", false)
	if err != nil {
		log.Fatalf("err %s", err)
	}

	_, err = s.Cards.GetVerifyCode(context.Background(), strconv.Itoa(rand.Int()), r.Result.Card.Token)
	if err != nil {
		log.Fatalf("err %s", err)
	}

	verifyCode, err := s.Cards.Verify(context.Background(), strconv.Itoa(rand.Int()), "666666", r.Result.Card.Token)

	fmt.Println(verifyCode, err)

}
