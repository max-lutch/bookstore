package creditcard

import "errors"

type card struct {
	Number string
}

func New(number string) (*card, error) {
	if number == "" {
		return nil, errors.New("invalid card number")
	}
	return &card{Number: number}, nil
}

func (c *card) SetCardNumber(number string) error {
	if c.Number == "" {
		return errors.New("Setting card number a zero")
	}
	c.Number = number
	return nil
}
