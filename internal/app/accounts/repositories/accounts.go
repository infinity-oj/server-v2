package repositories

import (
	"github.com/infinity-oj/server-v2/internal/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Repository interface {
	CreateAccount(name, hash, salt, email string) (u *models.Account, err error)
	QueryAccount(name string) (p *models.Account, err error)
	UpdateAccount(p *models.Account) (err error)

	UpdateCredential(u *models.Credential) (err error)
	QueryCredential(username string) (u *models.Credential, err error)
}

type DefaultRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

// CreateAccount
func (s *DefaultRepository) CreateAccount(username, hash, salt, email string) (account *models.Account, err error) {
	s.logger.Debug("create account",
		zap.String("username", username),
		zap.String("email", email),
	)

	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	account = &models.Account{
		Name:     username,
		Nickname: username,
		Email:    email,
	}
	if err = tx.Create(account).Error; err != nil {
		tx.Rollback()
		s.logger.Error("create account", zap.String("username", username), zap.Error(err))
		return nil, errors.Wrapf(err, " create account %s ", username)
	}

	credential := &models.Credential{
		Username: username,
		Hash:     hash,
		Salt:     salt,
	}
	if err = tx.Create(credential).Error; err != nil {
		tx.Rollback()
		s.logger.Error("create account", zap.String("username", username), zap.Error(err))
		return nil, errors.Wrapf(err, " create user with username: %s", username)
	}

	return account, tx.Commit().Error
}

func (s *DefaultRepository) UpdateAccount(p *models.Account) (err error) {
	// TODO: find a better way...
	err = s.db.Save(&p).Error
	return
}

func (s *DefaultRepository) UpdateCredential(u *models.Credential) (err error) {
	// TODO: find a better way...
	err = s.db.Save(&u).Error
	return
}

func (s *DefaultRepository) QueryCredential(username string) (credential *models.Credential, err error) {
	credential = &models.Credential{}
	if err = s.db.Where(&models.Credential{Username: username}).First(credential).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		} else {
			s.logger.Error("Query credential failed", zap.String("username", username), zap.Error(err))
		}
		return nil, err
	}
	return
}

func (s *DefaultRepository) QueryAccount(name string) (account *models.Account, err error) {
	account = &models.Account{}
	if err = s.db.Where(&models.Account{Name: name}).First(account).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		} else {
			s.logger.Error("query account failed", zap.String("name", name), zap.Error(err))
		}
		return nil, err
	}
	return
}

func New(logger *zap.Logger, db *gorm.DB) Repository {
	return &DefaultRepository{
		logger: logger.With(zap.String("type", "Account Repository")),
		db:     db,
	}
}