package testhelper

import "os"

func SetEnv() {
	os.Setenv("POSTGRES_USER", "testuser")
	os.Setenv("POSTGRES_PW", "testpassword")
	os.Setenv("POSTGRES_HOST", "localhost") //ローカルから実行するときは、127.0.0.1
	os.Setenv("POSTGRES_PORT", "15434")
	os.Setenv("POSTGRES_DB", "testmydb")
}

func ClearEnv() {
	os.Unsetenv("POSTGRES_USER")
	os.Unsetenv("POSTGRES_PW")
	os.Unsetenv("POSTGRES_HOST")
	os.Unsetenv("POSTGRES_PORT")
	os.Unsetenv("POSTGRES_DB")
}
