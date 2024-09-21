package database

import (
	"context"
	"log"

	"github.com/takuma123-type/go-api-study/src/domain/userdm"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

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
		return nil, err
	}
	return &user, nil
}

func (repo *userRepositoryImpl) Store(ctx context.Context, user *userdm.User) error {
	log.Printf("Storing user: %+v", user)

	createdAt := user.CreatedAt.Value()

	if err := repo.db.WithContext(ctx).Exec(`
		INSERT INTO users (id, first_name, last_name, created_at) 
		VALUES (?, ?, ?, ?)`,
		user.ID.String(), user.FirstName, user.LastName, createdAt).Error; err != nil {
		log.Printf("Failed to store user: %v", err)
		return err
	}
	return nil
}

func (repo *userRepositoryImpl) Update(ctx context.Context, user *userdm.User) error {
	log.Printf("Updating user: %+v", user)

	if err := repo.db.WithContext(ctx).Model(&user).Updates(map[string]interface{}{
		"first_name": user.FirstName,
		"last_name":  user.LastName,
	}).Error; err != nil {
		log.Printf("Failed to update user: %v", err)
		return err
	}
	return nil
}
