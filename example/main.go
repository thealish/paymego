package main

import (
	"context"
	"fmt"
	"github.com/thealish/paymego"
	"math/rand"
	"strconv"
)

func main() {
	ctx := context.Background()
	s, err := paymego.NewSubscribeAPI(&paymego.SubsribeAPIOpts{
		PaycomID:  "-",
		PaycomKey: "-",
		Mode:      paymego.Test,
	})
	if err != nil {
		panic(err)
	}
	c, err := s.CardsCreate(ctx, strconv.Itoa(rand.Int()), false, paymego.CardWithNoError.String(), "03/99")
	if err != nil {
		panic(err)
	}
	x, err := s.CardsGetVerifyCode(ctx, strconv.Itoa(rand.Int()), c.Result.Card.Token)
	fmt.Printf("VERIFY CODE --------------%+v-------------\n\n", x)
	if err != nil {
		panic(err)
	}
	v, err := s.CardsVerify(ctx, strconv.Itoa(rand.Int()), "666666", c.Result.Card.Token)
	fmt.Printf("CardsVerify --------------%+v-------------\n\n", v)
	if err != nil {
		panic(err)
	}
	z, err := s.CardsCheck(ctx, strconv.Itoa(rand.Int()), c.Result.Card.Token)
	fmt.Printf("CardsCheck --------------%+v-------------\n\n", z)
	if err != nil {
		panic(err)
	}
	r, err := s.CardsRemove(ctx, strconv.Itoa(rand.Int()), c.Result.Card.Token)
	fmt.Printf("CardsRemove --------------%+v-------------\n\n", r)
	if err != nil {
		panic(err)
	}
}
