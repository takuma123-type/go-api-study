package router_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/takuma123-type/go-api-study/src/domain/userdm"
	"github.com/takuma123-type/go-api-study/src/interface/controller"
	"github.com/takuma123-type/go-api-study/src/interface/presenter"
	"github.com/takuma123-type/go-api-study/src/usecase/userusecase/userinput"
)

type MockUserRepository struct{}

func (m *MockUserRepository) FindAll(ctx context.Context) ([]*userdm.User, error) {
	return []*userdm.User{}, nil
}

func (m *MockUserRepository) FindByID(ctx context.Context, id userdm.UserID) (*userdm.User, error) {
	return &userdm.User{}, nil
}

func (m *MockUserRepository) Create(ctx context.Context, user *userdm.User) error {
	if user.FirstName == "" || user.LastName == "" {
		return fmt.Errorf("無効なユーザーデータ")
	}
	return nil
}

func (m *MockUserRepository) Update(ctx context.Context, user *userdm.User) error {
	if user.FirstName == "" || user.LastName == "" {
		return fmt.Errorf("無効なユーザーデータ")
	}
	return nil
}

func (m *MockUserRepository) Store(ctx context.Context, user *userdm.User) error {
	if user.FirstName == "" || user.LastName == "" {
		return fmt.Errorf("無効なユーザーデータ")
	}
	return nil
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	api := router.Group("/api")

	mockRepo := &MockUserRepository{}

	api.GET("/users", func(ctx *gin.Context) {
		controller := controller.NewUserController(presenter.NewUserPresenter(ctx), mockRepo)
		err := controller.GetUserList(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	api.GET("/users/:id", func(ctx *gin.Context) {
		type reqStruct struct {
			ID string `uri:"id"`
		}
		var req reqStruct
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		controller := controller.NewUserController(presenter.NewUserPresenter(ctx), mockRepo)
		in := userinput.GetUserByIDInput{ID: req.ID}
		err := controller.GetUserByID(ctx, &in)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	api.POST("/user", func(ctx *gin.Context) {
		var in userinput.CreateUserInput
		if err := ctx.ShouldBindJSON(&in); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		controller := controller.NewUserController(presenter.NewUserPresenter(ctx), mockRepo)
		err := controller.CreateUser(ctx, &in)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusCreated, gin.H{"status": "User created"})
		}
	})

	api.PUT("/user/:id", func(ctx *gin.Context) {
		var in userinput.UpdateUserInput
		if err := ctx.ShouldBindJSON(&in); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		in.ID = ctx.Param("id")
		controller := controller.NewUserController(presenter.NewUserPresenter(ctx), mockRepo)
		err := controller.UpdateUser(ctx, &in)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"status": "User updated"})
		}
	})

	return router
}

func TestUserRouter(t *testing.T) {
	router := setupRouter()

	t.Run("GET /api/users", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/api/users", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("GET /api/users/:id", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/api/users/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("POST /api/user", func(t *testing.T) {
		userData := map[string]string{
			"first_name": "Test",
			"last_name":  "User",
			"email":      "test@example.com",
		}
		jsonData, _ := json.Marshal(userData)
		req, _ := http.NewRequest(http.MethodPost, "/api/user", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusCreated {
			t.Logf("POST /api/user エラー内容: %s", w.Body.String())
		}

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("PUT /api/user/:id", func(t *testing.T) {
		userData := map[string]string{
			"first_name": "Updated",
			"last_name":  "User",
			"email":      "updated@example.com",
		}
		jsonData, _ := json.Marshal(userData)
		req, _ := http.NewRequest(http.MethodPut, "/api/user/1", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Logf("PUT /api/user/:id エラー内容: %s", w.Body.String())
		}

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
