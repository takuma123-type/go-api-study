package presenter

import (
	"net/http"

	"github.com/takuma123-type/go-api-study/src/usecase/contractusecase/contractoutput"
)

type contractPresent struct {
	delivery Presenter
}

func NewContractPresenter(p Presenter) ContractPresenter {
	return &contractPresent{
		delivery: p,
	}
}

type ContractPresenter interface {
	CreateContract(out *contractoutput.CreateContractOutput)
}

func (c *contractPresent) CreateContract(out *contractoutput.CreateContractOutput) {
	c.delivery.JSON(http.StatusCreated, out)
}
