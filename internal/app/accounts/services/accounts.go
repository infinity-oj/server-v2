package services

import (
	"errors"
	"github.com/infinity-oj/server-v2/internal/app/accounts/repositories"
	"github.com/infinity-oj/server-v2/internal/pkg/crypto"
	"github.com/infinity-oj/server-v2/internal/pkg/models"
	"github.com/infinity-oj/server-v2/internal/pkg/utils/random"
	"go.uber.org/zap"
)

// You should never change it, otherwise all old credentials will turn invalid
const specialKey = "imf1nlTy0j"

type Service interface {
	GetAccount(name string) (account *models.Account, err error)
	UpdateAccount(account *models.Account, nickname, email, gender, locale string) (*models.Account, error)
	CreateAccount(username, password, email string) (account *models.Account, err error)

	VerifyCredential(username, password string) (isValid bool, err error)
}

type DefaultService struct {
	logger     *zap.Logger
	Repository repositories.Repository
}

func (s *DefaultService) GetAccount(name string) (account *models.Account, err error) {
	account, err = s.Repository.QueryAccount(name)
	return
}
func (s *DefaultService) UpdateAccount(account *models.Account, nickname, email, gender, locale string) (*models.Account, error) {
	s.logger.Debug("update account", zap.String("name", account.Name))
	account.Nickname = nickname
	account.Email = email
	account.Gender = gender
	account.Locale = locale
	err := s.Repository.UpdateAccount(account)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (s *DefaultService) CreateAccount(username, password, email string) (account *models.Account, err error) {
	s.logger.Debug("create account", zap.String("username", username))
	{
		account, err := s.GetAccount(username)
		if err != nil {
			return nil, err
		}
		if account != nil {
			return nil, errors.New("username exists")
		}
	}

	salt := random.RandStringRunes(64)
	hash := crypto.Sha256(salt + password + specialKey)
	if account, err = s.Repository.CreateAccount(username, hash, salt, email); err != nil {
		return nil, err
	}
	return
}

func (s *DefaultService) VerifyCredential(username, password string) (isValid bool, err error) {
	s.logger.Debug("verify credential", zap.String("username", username))
	u := new(models.Credential)
	if u, err = s.Repository.QueryCredential(username); err != nil {
		s.logger.Error("verify credential error", zap.Error(err))
		return false, err
	}
	hash := crypto.Sha256(u.Salt + password + specialKey)

	return hash == u.Hash, nil
}

func New(logger *zap.Logger, Repository repositories.Repository) Service {
	return &DefaultService{
		logger:     logger.With(zap.String("type", "Account Repository")),
		Repository: Repository,
	}
}