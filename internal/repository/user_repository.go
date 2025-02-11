package repository

import (
	"baseball-stats-app-back/internal/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	UpserUser(name, email, googleID string) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) UpserUser(name, email, googleID string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).FirstOrCreate(&user, entity.User{Email: email, Name: name}).Error; err != nil {
		return nil, err
	}

	auth := entity.ExternalAuth{
		UserID:         user.ID,
		Provider:       "google",
		ProviderUserID: googleID,
	}
	if err := r.db.Where("provider_user_id = ?", googleID).FirstOrCreate(&auth, auth).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
