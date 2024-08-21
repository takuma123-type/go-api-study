package mock_user

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/takuma123-type/go-api-study/src/domain/userdm"
)

func TestMockUserRepository_FindByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックのインスタンスを作成
	mockRepo := NewMockUserRepository(ctrl)
	expectedUserID := userdm.NewUserID()
	expectedUser := &userdm.User{}

	// モックの動作を定義
	mockRepo.EXPECT().FindByID(gomock.Any(), expectedUserID).Return(expectedUser, nil)

	// テストの実行
	user, err := mockRepo.FindByID(context.Background(), expectedUserID)

	// テストの検証
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestMockUserRepository_FindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックのインスタンスを作成
	mockRepo := NewMockUserRepository(ctrl)
	expectedUsers := []*userdm.User{
		&userdm.User{},
		&userdm.User{},
	}

	// モックの動作を定義
	mockRepo.EXPECT().FindAll(gomock.Any()).Return(expectedUsers, nil)

	// テストの実行
	users, err := mockRepo.FindAll(context.Background())

	// テストの検証
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, users)
}

func TestMockUserRepository_Store(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックのインスタンスを作成
	mockRepo := NewMockUserRepository(ctrl)
	userToStore := &userdm.User{}

	// モックの動作を定義
	mockRepo.EXPECT().Store(gomock.Any(), userToStore).Return(nil)

	// テストの実行
	err := mockRepo.Store(context.Background(), userToStore)

	// テストの検証
	assert.NoError(t, err)
}
