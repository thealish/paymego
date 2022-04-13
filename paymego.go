package paymego

import (
	"errors"
)

var (
	prodURL        = "https://checkout.paycom.uz/api"
	testURL        = "https://checkout.test.paycom.uz/api"
	ErrInvalidMode = errors.New("invalid mode")
)

type PaymeGoMode int

const (
	Production PaymeGoMode = iota
	Test
)

type PaymeCardsForTest int

const (
	ExpiredTestCard PaymeCardsForTest = iota
	BlockedTestCard
	CardWithUnexpectedError
	CardWithTimeoutError
	CardWithSmsInformingDisabled
	CardWithNoError
)

func (p PaymeCardsForTest) String() string {
	switch p {
	case ExpiredTestCard:
		return "3333336415804657"
	case BlockedTestCard:
		return "4444445987459073"
	case CardWithTimeoutError:
		return "8600134301849596"
	case CardWithUnexpectedError:
		return "8600143417770323"
	case CardWithNoError:
		return "8600495473316478"
	case CardWithSmsInformingDisabled:
		return "8600060921090842"
	}
	return ""
}

func (p PaymeGoMode) string() (string, error) {
	switch p {
	case Production:
		return prodURL, nil
	case Test:
		return testURL, nil
	}
	return "", ErrInvalidMode
}
