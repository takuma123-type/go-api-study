package database

import (
	"context"

	"github.com/takuma123-type/go-api-study/src/domain/userdm"
	"golang.org/x/xerrors"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (repo *UserRepositoryImpl) FindAll(ctx context.Context) ([]*userdm.User, error) {
	var users []*userdm.User
	if err := repo.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepositoryImpl) FindByID(ctx context.Context, id userdm.UserID) (*userdm.User, error) {
	var user userdm.User
	if err := repo.db.WithContext(ctx).Where("id = ?", id.String()).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, xerrors.Errorf("user not found: %s", id.String())
		}
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepositoryImpl) FindByName(ctx context.Context, firstName, lastName string) (*userdm.User, error) {
	var user userdm.User
	if err := repo.db.WithContext(ctx).
		Where("first_name = ? AND last_name = ?", firstName, lastName).
		First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepositoryImpl) Store(ctx context.Context, user *userdm.User) error {
	if err := repo.db.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}
