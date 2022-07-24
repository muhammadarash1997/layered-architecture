package service

import (
	"layered-architecture/model"
	"layered-architecture/repository"
)

type WalletService interface {
	Create(model.Wallet) (model.Wallet, error)
	Update(model.Wallet) (model.Wallet, error)
}

type walletService struct {
	walletRepository repository.WalletRepository
}

func (this *walletService) Create(wallet model.Wallet) (model.Wallet, error)

func (this *walletService) Update(wallet model.Wallet) (model.Wallet, error)