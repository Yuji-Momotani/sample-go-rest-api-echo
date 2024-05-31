package db

import (
	"database/sql"
	testhelper "sample-go-rest-api-echo/testHelper"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestNewDB(t *testing.T) {
	tests := []struct {
		name        string
		envSetup    func()
		expectError bool
	}{
		{
			name: "Success",
			envSetup: func() {
				testhelper.SetEnv()
			},
			expectError: false,
		},
		{
			name: "missing enviroment variables",
			envSetup: func() {
				testhelper.ClearEnv()
			},
			expectError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.envSetup()
			defer testhelper.ClearEnv()
			db, err := NewDB()
			if tt.expectError {
				// 失敗するテストケース
				assert.Error(t, err)
			} else {
				// 成功するテストケース
				assert.NotNil(t, db)
				sqlDB, err := db.DB()
				assert.NoError(t, err)
				assert.NoError(t, sqlDB.Ping())
			}
		})
	}
}

type MockDB struct {
	sqlDB *sql.DB
	err   error
}

func (mdb *MockDB) DB() (*sql.DB, error) {
	return mdb.sqlDB, mdb.err
}

func TestCloseDB(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name        string
		args        args
		expectError bool
	}{
		{
			name: "Success",
			args: args{
				db: func() *gorm.DB {
					testhelper.SetEnv()
					defer testhelper.ClearEnv()
					db, _ := NewDB()
					return db
				}(),
			},
			expectError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectError {
				// エラーのパターンはとりあえずスキップ
				t.Skip()
			} else {
				err := CloseDB(tt.args.db)
				assert.NoError(t, err)
				sqlDB, _ := tt.args.db.DB()
				assert.Error(t, sqlDB.Ping())
			}
		})
	}
}
