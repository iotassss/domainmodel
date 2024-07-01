package usecase

import (
	"context"

	"github.com/iotassss/domainmodel/internal/domain"
)

type GetUserByUUIDPresenter interface {
	Present(*domain.User) *GetUserByUUIDDTO
}

type GetUserByUUIDDTO struct {
	UUID      string `json:"uuid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type GetUserByUUIDUseCase interface {
	Execute(ctx context.Context, uuid string) (*GetUserByUUIDDTO, error)
}

type GetUserByUUIDInteractor struct {
	repo      domain.UserRepository
	presenter GetUserByUUIDPresenter
}

func NewGetUserByUUIDInteractor(repo domain.UserRepository, presenter GetUserByUUIDPresenter) *GetUserByUUIDInteractor {
	return &GetUserByUUIDInteractor{
		repo:      repo,
		presenter: presenter,
	}
}

func (cu *GetUserByUUIDInteractor) Execute(ctx context.Context, uuid string) (*GetUserByUUIDDTO, error) {
	uuidVO, err := domain.NewUUID(uuid)
	if err != nil {
		return nil, err
	}

	user, err := cu.repo.FindByUUID(ctx, uuidVO)
	if err != nil {
		return nil, err
	}

	dto := &GetUserByUUIDDTO{
		UUID:      user.UUID().Value(),
		FirstName: user.FirstName(),
		LastName:  user.LastName(),
		Email:     user.Email().Address(),
	}

	return dto, nil
}
