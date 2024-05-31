package repository

import (
	"sample-go-rest-api-echo/db"
	"sample-go-rest-api-echo/model"
	testhelper "sample-go-rest-api-echo/testHelper"
	"testing"

	"gorm.io/gorm"
)

func cleanupDB(db *gorm.DB) {
	db.Exec("DELETE FROM users")
}

func TestNewUserRepository(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want IUserRepository
	}{
		{
			name: "Create New User Repository",
			args: args{
				db: func() *gorm.DB {
					testhelper.SetEnv()
					defer testhelper.ClearEnv()
					db, _ := db.NewDB()
					return db
				}(),
			},
			want: &userRepository{
				db: func() *gorm.DB {
					testhelper.SetEnv()
					defer testhelper.ClearEnv()
					db, _ := db.NewDB()
					return db
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUserRepository(tt.args.db)
			// Check if the type is correct
			if _, ok := got.(*userRepository); !ok {
				t.Errorf("NewUserRepository() = %T, want *userRepository", got)
			}
			// Check if the db field is set
			gotRepo := got.(*userRepository)
			if gotRepo.db != tt.args.db {
				t.Errorf("NewUserRepository().db = %v, want %v", gotRepo.db, tt.args.db)
			}
		})
	}
}

func TestGetUserByEmail(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		user  *model.User
		email string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		initfunc func(dbConn *gorm.DB, t *testing.T)
		wantErr  bool
	}{
		{
			name: "User found",
			fields: fields{
				db: func() *gorm.DB {
					testhelper.SetEnv()
					defer testhelper.ClearEnv()
					dbConn, _ := db.NewDB()
					dbConn.AutoMigrate(&model.User{})
					return dbConn
				}(),
			},
			initfunc: func(dbConn *gorm.DB, t *testing.T) {
				err := dbConn.Create(&model.User{Email: "test@example.com", Password: "password"}).Error
				if err != nil {
					t.Fatal("create error")
				}
			},
			args: args{
				user:  &model.User{},
				email: "test@example.com",
			},
			wantErr: false,
		},
		{
			name: "User not found",
			fields: fields{
				db: func() *gorm.DB {
					testhelper.SetEnv()
					defer testhelper.ClearEnv()
					dbConn, err := db.NewDB()
					if err != nil {
						t.Fatalf("Error Connection DB: %v", err)
					}
					dbConn.AutoMigrate(&model.User{})
					return dbConn
				}(),
			},
			initfunc: nil,
			args: args{
				user:  &model.User{},
				email: "nonexistent@example.com",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ur := &userRepository{
				db: tt.fields.db,
			}
			cleanupDB(ur.db)
			if tt.initfunc != nil {
				tt.initfunc(ur.db, t)
			}
			if err := ur.GetUserByEmail(tt.args.user, tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("userRepository.GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		user *model.User
	}
	tests := []struct {
		name     string
		fields   fields
		initfunc func(dbConn *gorm.DB, t *testing.T)
		args     args
		wantErr  bool
	}{
		{
			name: "Successful user creation",
			fields: fields{
				db: func() *gorm.DB {
					testhelper.SetEnv()
					defer testhelper.ClearEnv()
					dbConn, _ := db.NewDB()
					dbConn.AutoMigrate(&model.User{})
					return dbConn
				}(),
			},
			initfunc: nil,
			args: args{
				user: &model.User{
					Email:    "test@example.com",
					Password: "password",
				},
			},
			wantErr: false,
		},
		{
			name: "Duplicate email error",
			fields: fields{
				db: func() *gorm.DB {
					testhelper.SetEnv()
					defer testhelper.ClearEnv()
					dbConn, _ := db.NewDB()
					dbConn.AutoMigrate(&model.User{})
					return dbConn
				}(),
			},
			initfunc: func(dbConn *gorm.DB, t *testing.T) {
				err := dbConn.Create(&model.User{Email: "test@example.com", Password: "password"}).Error
				if err != nil {
					t.Fatal("initfunc err")
				}
			},
			args: args{
				user: &model.User{
					Email:    "test@example.com",
					Password: "password",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			ur := &userRepository{
				db: tt.fields.db,
			}
			cleanupDB(ur.db)
			if tt.initfunc != nil {
				tt.initfunc(ur.db, t)
			}
			if err := ur.CreateUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("userRepository.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
