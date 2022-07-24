package repository

import (
	"layered-architecture/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type WalletRepository interface {
	Create(model.Wallet) (model.Wallet, error)
	Update(model.Wallet) (model.Wallet, error)
}

type walletRepository struct {
	db *mongo.Database
}

func (this *walletRepository) Create(wallet model.Wallet) (model.Wallet, error)

func (this *walletRepository) Update(wallet model.Wallet) (model.Wallet, error)