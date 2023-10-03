package storer

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jasimmons/game-backend-api/pkg/models"
)

type mysqlStorer struct {
	db *sql.DB
}

var _ Storer = (*mysqlStorer)(nil)

func New(db *sql.DB) (Storer, error) {
	s := &mysqlStorer{db: db}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Ping(ctx); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *mysqlStorer) Ping(ctx context.Context) error {
	var one int
	row := s.db.QueryRowContext(ctx, "SELECT 1")
	return row.Scan(&one)
}

func (s *mysqlStorer) CreateAccount(ctx context.Context, opts CreateAccountOpts) (*models.Account, error) {
	return nil, errors.New("not implemented")
}

func (s *mysqlStorer) GetAccount(ctx context.Context, opts GetAccountOpts) (*models.Account, error) {
	return nil, errors.New("not implemented")
}

func (s *mysqlStorer) ListAccounts(ctx context.Context, opts ListAccountOpts) ([]models.Account, error) {
	return nil, errors.New("not implemented")
}

func (s *mysqlStorer) UpdateAccount(ctx context.Context, opts UpdateAccountOpts) (*models.Account, error) {
	return nil, errors.New("not implemented")
}

func (s *mysqlStorer) DeleteAccount(ctx context.Context, opts DeleteAccountOpts) error {
	return errors.New("not implemented")
}
