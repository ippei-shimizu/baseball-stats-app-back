package usecase

import (
	"baseball-stats-app-back/internal/entity"
	"baseball-stats-app-back/internal/repository"
	"context"
	"errors"
	"os"

	"google.golang.org/api/idtoken"
)

type UserUseCase interface {
	AuthenticateGoogleUser(idToken string) (*entity.User, error)
}

type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{userRepo: userRepo}
}

func (u *userUseCase) AuthenticateGoogleUser(idToken string) (*entity.User, error) {
	ctx := context.Background()
	clientID := os.Getenv("GOOGLE_CLIENT_ID")

	payload, err := idtoken.Validate(ctx, idToken, clientID)
	if err != nil {
		return nil, errors.New("GoogleIDToken validation failed")
	}

	googleID := payload.Claims["sub"].(string)
	email := payload.Claims["email"].(string)
	name := payload.Claims["name"].(string)

	user, err := u.userRepo.UpserUser(name, email, googleID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
