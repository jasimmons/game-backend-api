package storer

import (
	"context"

	"github.com/jasimmons/game-backend-api/pkg/models"
)

type Storer interface {
	Ping(context.Context) error

	CreateAccount(context.Context, CreateAccountOpts) (*models.Account, error)
	GetAccount(context.Context, GetAccountOpts) (*models.Account, error)
	ListAccounts(context.Context, ListAccountOpts) ([]models.Account, error)
	UpdateAccount(context.Context, UpdateAccountOpts) (*models.Account, error)
	DeleteAccount(context.Context, DeleteAccountOpts) error
}

type CreateAccountOpts struct{}
type GetAccountOpts struct{}
type ListAccountOpts struct{}
type UpdateAccountOpts struct{}
type DeleteAccountOpts struct{}
