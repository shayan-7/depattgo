package behavioral

import (
	"errors"
)

type TransmitionRequest struct {
	Money     int
	ItemCount int
}

func NewTransmitionRequest(money, itemCount int) *TransmitionRequest {
	return &TransmitionRequest{money, itemCount}
}

type State interface {
	GetValue() string
	Transmit() (State, error)
}

type VendingMachine struct {
	Current State
	Request *TransmitionRequest

	ItemCount int
	ItemPrice int
}

func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		ItemCount: itemCount,
		ItemPrice: itemPrice,
	}
	switch itemCount {
	case 0:
		v.Current = NewNoItemState(v)
	default:
		v.Current = NewHasItemState(v)
	}
	return v
}

func (v *VendingMachine) AddItem(itemCount int) {
	v.ItemCount += itemCount
}

func (v *VendingMachine) MakeRequest(r *TransmitionRequest) error {
	if v.Request != nil {
		return errors.New("request already in process")
	}
	v.Request = r
	return nil
}

func (v *VendingMachine) RequestDone() {
	v.Request = nil
}

func (v *VendingMachine) Proceed() error {
	current, err := v.Current.Transmit()
	if err != nil {
		v.RequestDone()
	}
	v.Current = current
	return err
}

type HasItemState struct {
	v string
	*VendingMachine
}

func NewHasItemState(v *VendingMachine) *HasItemState {
	return &HasItemState{"has item", v}
}

func (s *HasItemState) GetValue() string {
	return s.v
}

func (s *HasItemState) Transmit() (State, error) {
	if s.Request == nil {
		return s, errors.New("no request to proceed")
	}
	if s.ItemCount < s.Request.ItemCount {
		return s, errors.New("not enough items")
	}
	return NewItemRequestedState(s.VendingMachine), nil
}

type NoItemState struct {
	v string
	*VendingMachine
}

func NewNoItemState(v *VendingMachine) *NoItemState {
	return &NoItemState{"no item", v}
}

func (s *NoItemState) GetValue() string {
	return s.v
}

func (s *NoItemState) Transmit() (State, error) {
	if s.Request == nil {
		return s, errors.New("no request to proceed")
	}
	s.ItemCount += s.Request.ItemCount
	return NewHasItemState(s.VendingMachine), nil
}

type ItemRequestedState struct {
	v string
	*VendingMachine
}

func NewItemRequestedState(v *VendingMachine) *ItemRequestedState {
	return &ItemRequestedState{"item requested", v}
}

func (s *ItemRequestedState) GetValue() string {
	return s.v
}

func (s *ItemRequestedState) Transmit() (State, error) {
	r := s.Request
	if r == nil {
		return s, errors.New("no request to proceed")
	}
	if r.ItemCount*s.ItemPrice > r.Money {
		return NewHasItemState(s.VendingMachine), errors.New("not enough money")
	}
	return NewHasMoneyState(s.VendingMachine), nil
}

type HasMoneyState struct {
	v string
	*VendingMachine
}

func NewHasMoneyState(v *VendingMachine) *HasMoneyState {
	return &HasMoneyState{"has money", v}
}

func (s *HasMoneyState) GetValue() string {
	return s.v
}

func (s *HasMoneyState) Transmit() (State, error) {
	r := s.Request
	if s.Request == nil {
		return s, errors.New("no request to proceed")
	}
	r.Money -= r.ItemCount * s.ItemPrice
	s.ItemCount -= r.ItemCount

	s.RequestDone()
	if s.ItemCount == 0 {
		return NewNoItemState(s.VendingMachine), nil
	}
	return NewHasItemState(s.VendingMachine), nil
}
