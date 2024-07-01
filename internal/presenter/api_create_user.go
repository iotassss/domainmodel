package presenter

import (
	"github.com/iotassss/domainmodel/internal/domain"
	"github.com/iotassss/domainmodel/internal/usecase"
)

type APICreateUserPresenter struct{}

func NewAPICreateUserPresenter() *APICreateUserPresenter {
	return &APICreateUserPresenter{}
}

func (p *APICreateUserPresenter) Present(user *domain.User) *usecase.CreateUserDTO {
	return &usecase.CreateUserDTO{
		UUID:      user.UUID().Value(),
		FirstName: user.FirstName(),
		LastName:  user.LastName(),
		Email:     user.Email().Address(),
	}
}
