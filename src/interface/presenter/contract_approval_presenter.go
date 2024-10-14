package presenter

import (
	"net/http"

	"github.com/takuma123-type/go-api-study/src/usecase/contractapprovalusecase/contractapprovaloutput"
)

type contractApprovalPresent struct {
	delivery Presenter
}

func NewContractApprovalPresenter(p Presenter) ContractApprovalPresenter {
	return &contractApprovalPresent{
		delivery: p,
	}
}

type ContractApprovalPresenter interface {
	CreateContractApproval(out *contractapprovaloutput.CreateContractApprovalOutput)
}

func (c *contractApprovalPresent) CreateContractApproval(out *contractapprovaloutput.CreateContractApprovalOutput) {
	c.delivery.JSON(http.StatusCreated, out)
}
