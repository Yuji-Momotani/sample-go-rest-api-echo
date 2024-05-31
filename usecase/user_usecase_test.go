package usecase

import (
	"errors"
	"os"
	"reflect"
	"sample-go-rest-api-echo/model"
	"sample-go-rest-api-echo/repository"
	"sample-go-rest-api-echo/validator"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type MockIUserRepository struct {
	mock.Mock
}

func (m *MockIUserRepository) CreateUser(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockIUserRepository) GetUserByEmail(user *model.User, email string) error {
	args := m.Called(user, email)
	return args.Error(0)
}

type MockIUserValidator struct {
	mock.Mock
}

func (m *MockIUserValidator) UserValidate(user model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func TestNewUserUsecase(t *testing.T) {
	type args struct {
		ur repository.IUserRepository
		uv validator.IUserValidator
	}
	tests := []struct {
		name string
		args args
		want IUserUsecase
	}{
		{
			name: "Create new user usecase",
			args: args{
				ur: &MockIUserRepository{},
				uv: &MockIUserValidator{},
			},
			want: &userUsecase{
				ur: &MockIUserRepository{},
				uv: &MockIUserValidator{},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserUsecase(tt.args.ur, tt.args.uv); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingUp(t *testing.T) {
	type fields struct {
		ur repository.IUserRepository
		uv validator.IUserValidator
	}
	type args struct {
		user model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.UserResponse
		wantErr bool
	}{
		{
			name: "Successful sign up",
			fields: fields{
				ur: func() repository.IUserRepository {
					mockUR := &MockIUserRepository{}
					mockUR.On("CreateUser", mock.AnythingOfType("*model.User")).Run(func(args mock.Arguments) {
						arg := args.Get(0).(*model.User)
						arg.ID = 1
					}).Return(nil)
					return mockUR
				}(),
				uv: func() validator.IUserValidator {
					mockUV := &MockIUserValidator{}
					mockUV.On("UserValidate", mock.AnythingOfType("model.User")).Return(nil)
					return mockUV
				}(),
			},
			args: args{
				user: model.User{
					Email:    "test@example.com",
					Password: "password",
				},
			},
			want: model.UserResponse{
				ID:    1,
				Email: "test@example.com",
			},
			wantErr: false,
		},
		{
			name: "Validation error",
			fields: fields{
				ur: &MockIUserRepository{},
				uv: func() validator.IUserValidator {
					mockUV := &MockIUserValidator{}
					mockUV.On("UserValidate", mock.AnythingOfType("model.User")).Return(errors.New("validation error"))
					return mockUV
				}(),
			},
			args: args{
				user: model.User{
					Email:    "test@example.com",
					Password: "password",
				},
			},
			want:    model.UserResponse{},
			wantErr: true,
		},
		{
			name: "Create user error",
			fields: fields{
				ur: func() repository.IUserRepository {
					mockUR := &MockIUserRepository{}
					mockUR.On("CreateUser", mock.AnythingOfType("*model.User")).Return(errors.New("create user error"))
					return mockUR
				}(),
				uv: func() validator.IUserValidator {
					mockUV := &MockIUserValidator{}
					mockUV.On("UserValidate", mock.AnythingOfType("model.User")).Return(nil)
					return mockUV
				}(),
			},
			args: args{
				user: model.User{
					Email:    "test@example.com",
					Password: "password",
				},
			},
			want:    model.UserResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			uu := &userUsecase{
				ur: tt.fields.ur,
				uv: tt.fields.uv,
			}
			got, err := uu.SingUp(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.SingUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.SingUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogIn(t *testing.T) {
	type fields struct {
		ur repository.IUserRepository
		uv validator.IUserValidator
	}
	type args struct {
		user model.User
	}
	stringPassword := "password"
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(stringPassword), 10)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Successful login",
			fields: fields{
				ur: func() repository.IUserRepository {
					mockUR := &MockIUserRepository{}
					mockUR.On("GetUserByEmail", mock.AnythingOfType("*model.User"), "test@example.com").Run(func(args mock.Arguments) {
						arg := args.Get(0).(*model.User)
						arg.ID = 1
						arg.Email = "test@example.com"
						arg.Password = string(hashPassword)
					}).Return(nil)
					return mockUR
				}(),
				uv: &MockIUserValidator{},
			},
			args: args{
				user: model.User{
					Email:    "test@example.com",
					Password: stringPassword,
				},
			},
			want:    "someTokenString", // Sucessの具体的なJWTの値チェックはしない
			wantErr: false,
		},
		{
			name: "Invalid password",
			fields: fields{
				ur: func() repository.IUserRepository {
					mockUR := &MockIUserRepository{}
					mockUR.On("GetUserByEmail", mock.AnythingOfType("*model.User"), "test@example.com").Run(func(args mock.Arguments) {
						arg := args.Get(0).(*model.User)
						arg.ID = 1
						arg.Email = "test@example.com"
						arg.Password = string(hashPassword)
					}).Return(nil)
					return mockUR
				}(),
				uv: &MockIUserValidator{},
			},
			args: args{
				user: model.User{
					Email:    "test@example.com",
					Password: "wrongpassword",
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "User not found",
			fields: fields{
				ur: func() repository.IUserRepository {
					mockUR := &MockIUserRepository{}
					mockUR.On("GetUserByEmail", mock.AnythingOfType("*model.User"), "test@example.com").Return(errors.New("record not found"))
					return mockUR
				}(),
				uv: &MockIUserValidator{},
			},
			args: args{
				user: model.User{
					Email:    "test@example.com",
					Password: stringPassword,
				},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			// Set the SECRET environment variable for testing
			os.Setenv("SECRET", "mysecret")
			defer os.Unsetenv("SECRET")

			uu := &userUsecase{
				ur: tt.fields.ur,
				uv: tt.fields.uv,
			}
			got, err := uu.LogIn(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.LogIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if got != tt.want {
					t.Errorf("userUsecase.LogIn() = %v, want %v", got, tt.want)
				}
			} else {
				// 成功時はJWTの具体的な値まではチェックしない
				if strings.TrimSpace(got) == "" {
					t.Errorf("userUsecase.LogIn() = not get JWT token : %v", got)
				}
			}
		})
	}
}
