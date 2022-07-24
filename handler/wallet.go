package handler

import (
	"layered-architecture/model"
	"layered-architecture/service"
)

type WalletHandler interface {
	Create(model.Wallet) (model.Wallet, error)
	Update(model.Wallet) (model.Wallet, error)
}

type walletHandler struct {
	walletService service.WalletService
}

func (this *walletHandler) Create(wallet model.Wallet) (model.Wallet, error)

func (this *walletHandler) Update(wallet model.Wallet) (model.Wallet, error)
