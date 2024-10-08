package presenter

import (
	"net/http"

	"github.com/takuma123-type/go-api-study/src/usecase/userusecase/useroutput"
)

type userPresent struct {
	delivery Presenter
}

func NewUserPresenter(p Presenter) UserPresenter {
	return &userPresent{
		delivery: p,
	}
}

type UserPresenter interface {
	UserList(out *useroutput.UserList)
	UserByID(out *useroutput.UserByID)
	Create(out *useroutput.CreateUserOutput)
	Update(out *useroutput.UpdateUserOutput)
}

func (p *userPresent) UserList(out *useroutput.UserList) {
	p.delivery.JSON(http.StatusOK, out)
}

func (p *userPresent) UserByID(out *useroutput.UserByID) {
	p.delivery.JSON(http.StatusOK, out)
}
func (p *userPresent) Create(out *useroutput.CreateUserOutput) {
	p.delivery.JSON(http.StatusCreated, out)
}

func (p *userPresent) Update(out *useroutput.UpdateUserOutput) {
	p.delivery.JSON(http.StatusOK, out)
}
