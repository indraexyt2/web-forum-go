package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/indraexyt2/web-forum-go/internal/model/memberships"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(ctx context.Context, req memberships.SignUpRequest) (string, string, error) {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, req.Username, 0)
	if err != nil {
		return "", "", err
	}

	if user != nil {
		return "", "", errors.New("username or email already exist")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}

	now := time.Now()
	model := memberships.UserModel{
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(pass),
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: req.Email,
		UpdatedBy: req.Email,
	}

	err = s.membershipRepo.CreateUser(ctx, model)
	if err != nil {
		return "", "", err
	}

	return model.Username, model.Email, nil
}
