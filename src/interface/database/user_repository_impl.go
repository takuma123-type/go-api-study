// src/interface/database/user_repository_impl.go
package database

import (
	"context"

	"github.com/takuma123-type/go-api-study/src/domain/userdm"
	"github.com/takuma123-type/go-api-study/src/support/smperr"
	"gorm.io/gorm"
)

// userRepositoryImpl は private 構造体として定義
type userRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepositoryImpl は public にして外部からインスタンスを生成できるようにする
func NewUserRepositoryImpl(db *gorm.DB) *userRepositoryImpl {
	return &userRepositoryImpl{
		db: db,
	}
}

func (repo *userRepositoryImpl) FindAll(ctx context.Context) ([]*userdm.User, error) {
	var users []*userdm.User
	if err := repo.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *userRepositoryImpl) FindByID(ctx context.Context, id userdm.UserID) (*userdm.User, error) {
	var user userdm.User
	if err := repo.db.WithContext(ctx).Where("id = ?", id.String()).First(&user).Error; err != nil {
		if smperr.IsRecordNotFound(err) {
			return nil, smperr.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (repo *userRepositoryImpl) Store(ctx context.Context, user *userdm.User) error {
	if err := repo.db.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}
