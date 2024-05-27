package usecase

import (
	"os"
	"time"
	"todo-rest-api-3/model"
	"todo-rest-api-3/repository"
	"todo-rest-api-3/validator"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// インターフェース
type IUserUsecase interface {
	SingUp(user model.User) (model.UserResponse, error)
	LogIn(user model.User) (string, error)
}

// インターフェースを実装するstruct
type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

// コンストラクタ
func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur: ur, uv: uv}
}

//処理部

// SignUp：ユーザー登録
func (uu *userUsecase) SingUp(user model.User) (model.UserResponse, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{
		Email:    user.Email,
		Password: string(hashPass),
	}
	if err := uu.uv.UserValidate(newUser); err != nil {
		return model.UserResponse{}, err
	}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	userRes := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return userRes, nil
}

func (uu *userUsecase) LogIn(user model.User) (string, error) {
	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
