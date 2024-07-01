package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/iotassss/domainmodel/internal/domain"
)

type CreateUserPresenter interface {
	Present(*domain.User) *CreateUserDTO
}

type CreateUserDTO struct {
	UUID      string `json:"uuid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type CreateUserUseCase interface {
	Execute(ctx context.Context, firstName, lastName, email, password string) (*CreateUserDTO, error)
}

type CreateUserInteractor struct {
	repo      domain.UserRepository
	presenter CreateUserPresenter
}

func NewCreateUserInteractor(repo domain.UserRepository, presenter CreateUserPresenter) *CreateUserInteractor {
	return &CreateUserInteractor{
		repo:      repo,
		presenter: presenter,
	}
}

func (cu *CreateUserInteractor) Execute(ctx context.Context, firstName, lastName, email, password string) (*CreateUserDTO, error) {
	uuid := uuid.New().String()
	uuidVO, err := domain.NewUUID(uuid)
	if err != nil {
		return nil, err
	}

	emailVO, err := domain.NewEmail(email)
	if err != nil {
		return nil, err
	}

	user, err := domain.NewUser(uuidVO, firstName, lastName, emailVO)
	if err != nil {
		return nil, err
	}

	credential := domain.NewCredential(uuidVO)
	if err := credential.SetPassword(password); err != nil {
		return nil, err
	}
	user.SetCredential(credential)

	if err := cu.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return cu.presenter.Present(user), nil

}
