package model

import (
	"errors"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type PixKey struct {
	Base      `valid:"required"`
	Kind      string `json:"kind" valid:"notnull"`
	Key       string `json:"key"`
	AccountID string `json:"account_id"`
	Account   *Account
	Status    string `json:"status"`
}

type iPixKeyRepository interface {
	Register(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccount(id string) (*Account, error)
}

func (pixKey *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(pixKey)

	if pixKey.Kind != "email" && pixKey.Kind != "cpf" {
		return errors.New("invalid type of key")
	}

	if pixKey.Status != "active" && pixKey.Status != "inactive" {
		return errors.New("invalid status")
	}

	if err != nil {
		return err
	}
	return nil
}

func NewPixKey(account *Account, kind string, key string) (*PixKey, error) {
	pixKey := PixKey{
		Kind:    kind,
		Key:     key,
		Account: account,
		Status:  "active",
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()

	err := pixKey.isValid()
	if err != nil {
		return nil, err
	}

	return &pixKey, nil
}