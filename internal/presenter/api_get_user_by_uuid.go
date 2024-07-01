package presenter

import (
	"github.com/iotassss/domainmodel/internal/domain"
	"github.com/iotassss/domainmodel/internal/usecase"
)

type APIGetUserByUUIDPresenter struct{}

func NewAPIGetUserByUUIDPresenter() *APIGetUserByUUIDPresenter {
	return &APIGetUserByUUIDPresenter{}
}

func (p *APIGetUserByUUIDPresenter) Present(user *domain.User) *usecase.GetUserByUUIDDTO {
	return &usecase.GetUserByUUIDDTO{
		UUID:      user.UUID().Value(),
		FirstName: user.FirstName(),
		LastName:  user.LastName(),
		Email:     user.Email().Address(),
	}
}
