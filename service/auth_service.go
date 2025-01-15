package service

import (
	"fmt"
	"strings"

	"github.com/Minyeng/go-auth-jwt/auth"
	"github.com/Minyeng/go-auth-jwt/entity"
	"github.com/Minyeng/go-auth-jwt/model"
	"github.com/Minyeng/go-auth-jwt/validator"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UserLogin(db *gorm.DB, req *model.LoginRequest) (string, int64, any, error) {
	user := &entity.User{}
	result := db.First(user, "email = ?", req.Email).Order("id DESC")
	if result.Error != nil {
		return "", 0, nil, result.Error
	}
	if result.RowsAffected == 0 {
		return "", 0, nil, fiber.NewError(fiber.StatusBadRequest, "invalid login credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", 0, nil, err
	}

	token, exp, err := auth.CreateJWTToken(*user)
	if err != nil {
		return "", 0, nil, err
	}

	return token, exp, user, nil
}

func UserSignup(db *gorm.DB, req *model.SignupRequest) (string, int64, any, error) {
	// Validation
	if errs := validator.MyValidator.Validate(req); len(errs) > 0 && errs[0].Error {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}

		return "", 0, nil, &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: strings.Join(errMsgs, " and "),
		}
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", 0, nil, err
	}

	user := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hash),
	}

	result := db.Create(user)

	if result.Error != nil {
		return "", 0, nil, result.Error
	}

	token, exp, err := auth.CreateJWTToken(*user)

	if err != nil {
		return "", 0, nil, err
	}

	return token, exp, user, nil
}
