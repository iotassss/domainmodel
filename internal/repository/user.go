package repository

import (
	"context"

	"github.com/iotassss/domainmodel/internal/domain"
	"github.com/iotassss/domainmodel/internal/repository/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, entityUser *domain.User) error {
	var existingUser model.User
	if err := r.db.WithContext(ctx).Where("uuid = ?", entityUser.UUID().Value()).First(&existingUser).Error; err == nil {
		return &domain.ConflictError{Msg: "uuid already exists"}
	}
	if err := r.db.WithContext(ctx).Where("email = ?", entityUser.Email().Address()).First(&existingUser).Error; err == nil {
		return &domain.ConflictError{Msg: "email already exists"}
	}

	tx := r.db.WithContext(ctx).Begin()
	if err := tx.Error; err != nil {
		return &domain.ServerError{Msg: "failed to start transaction"}
	}

	user := &model.User{
		UUID:      entityUser.UUID().Value(),
		FirstName: entityUser.FirstName(),
		LastName:  entityUser.LastName(),
		Email:     entityUser.Email().Address(),
	}
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return &domain.ServerError{Msg: "failed to create user"}
	}

	credential := &model.Credential{
		UserUUID:     entityUser.UUID().Value(),
		PasswordHash: entityUser.Credential().PasswordHash(),
	}
	if err := tx.Create(credential).Error; err != nil {
		tx.Rollback()
		return &domain.ServerError{Msg: "failed to create credential"}
	}

	if err := tx.Commit().Error; err != nil {
		return &domain.ServerError{Msg: "failed to commit transaction"}
	}

	return nil
}

func (r *UserRepository) FindByUUID(ctx context.Context, uuid domain.UUID) (*domain.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Where("uuid = ?", uuid.Value()).First(&user).Error; err != nil {
		return nil, &domain.NotFoundError{Msg: "user not found"}
	}

	uuidVO, err := domain.NewUUID(user.UUID)
	if err != nil {
		return nil, err
	}
	emailVO, err := domain.NewEmail(user.Email)
	if err != nil {
		return nil, err
	}
	entityUser, err := domain.NewUser(
		uuidVO,
		user.FirstName,
		user.LastName,
		emailVO,
	)
	if err != nil {
		return nil, err
	}

	return entityUser, nil
}
