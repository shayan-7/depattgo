package structural

import (
	"errors"
)

// Account sub-system
type account struct {
	name string
}

func NewAccount(name string) *account {
	return &account{name}
}

func (a *account) checkAccount(name string) error {
	if a.name != name {
		return errors.New("account name doesn't match")
	}
	return nil
}

// Security Code sub-system
type securityCode struct {
	code int
}

func NewSecurityCode(code int) *securityCode {
	return &securityCode{code}
}

func (sc *securityCode) checkCode(code int) error {
	if sc.code != code {
		return errors.New("security code doesn't match")
	}
	return nil
}

// Wallet sub-system
type wallet struct {
	balance int
}

func NewWallet() *wallet {
	return &wallet{0}
}

func (w *wallet) CreditBalance(amount int) {
	w.balance += amount
}

func (w *wallet) DebitBalance(amount int) error {
	if w.balance < amount {
		return errors.New("not enough balance")
	}
	w.balance -= amount
	return nil
}

// Transaction Registry sub-system
type txRegistry struct {
	isRegistered bool
}

func (tr *txRegistry) Add() {
	tr.isRegistered = !tr.isRegistered
}

// Notification sub-system
type notification struct {
	msg string
}

func (n *notification) sendWalletCreditNotification() {
	n.msg = "wallet credit is done"
}

func (n *notification) sendWalletDebitNotification() {
	n.msg = "wallet debit is done"
}

// Wallet Facade to handle complex sub-systems
type WalletFacade struct {
	account      *account
	wallet       *wallet
	securityCode *securityCode
	txRegistry   *txRegistry
	notification *notification
}

func NewWalletFacade(accountID string, code int) *WalletFacade {
	return &WalletFacade{
		account:      NewAccount(accountID),
		wallet:       NewWallet(),
		securityCode: NewSecurityCode(code),
		txRegistry:   &txRegistry{},
		notification: &notification{},
	}
}

func (wf *WalletFacade) AddMoney(
	accountId string, securityCode, amount int) error {

	err := wf.account.checkAccount(accountId)
	if err != nil {
		return err
	}

	err = wf.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}

	wf.wallet.CreditBalance(amount)
	wf.notification.sendWalletCreditNotification()
	wf.txRegistry.Add()
	return nil
}

func (wf *WalletFacade) DeductMoney(
	accountId string, securityCode, amount int) error {

	err := wf.account.checkAccount(accountId)
	if err != nil {
		return err
	}

	err = wf.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}

	err = wf.wallet.DebitBalance(amount)
	if err != nil {
		return err
	}

	wf.notification.sendWalletDebitNotification()
	wf.txRegistry.Add()
	return nil
}
