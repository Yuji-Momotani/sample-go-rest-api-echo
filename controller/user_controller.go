package controller

import (
	"net/http"
	"os"
	"time"
	"todo-rest-api-3/model"
	"todo-rest-api-3/usecase"

	"github.com/labstack/echo/v4"
)

// インターフェース
type IUserController interface {
	SingUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
}

// インターフェースを実装するstruct
type userController struct {
	uu usecase.IUserUsecase
}

// コンストラクタ
func NewUserUsecase(uu usecase.IUserUsecase) IUserController {
	return &userController{uu: uu}
}

// 実装部
const notFound string = "record not found"

func (uc *userController) SingUp(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := uc.uu.SingUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, userRes)
}
func (uc *userController) LogIn(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tokenString, err := uc.uu.LogIn(user)
	if err != nil {
		if err.Error() == notFound {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(time.Hour * 12)
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}
func (uc *userController) LogOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Path = "/"
	cookie.Expires = time.Now()
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uc *userController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}
