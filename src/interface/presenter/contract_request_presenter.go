package presenter

import (
	"net/http"

	"github.com/takuma123-type/go-api-study/src/usecase/contractrequestusecase/contractrequestoutput"
)

type contractRequestPresent struct {
	delivery Presenter
}

func NewContractRequestPresenter(p Presenter) ContractRequestPresenter {
	return &contractRequestPresent{
		delivery: p,
	}
}

type ContractRequestPresenter interface {
	CreateContractRequest(out *contractrequestoutput.CreateContractRequestOutput)
}

func (c *contractRequestPresent) CreateContractRequest(out *contractrequestoutput.CreateContractRequestOutput) {
	c.delivery.JSON(http.StatusCreated, out)
}
