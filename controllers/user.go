package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"math/rand"
	"net/http"
	"strings"
	"xq_music_go/commons"
	"xq_music_go/models/bean"
	"xq_music_go/models/dto"
	"xq_music_go/models/vo"
	"xq_music_go/mymiddleware/uuid"
	"xq_music_go/service"
)

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	fmt.Println(username, password)
	userService := &service.UserServiceImpl{}
	user, err := userService.Login(username, password)
	if user == nil || err != nil {
		return c.JSON(http.StatusOK, commons.Resp{Code: -1, Message: "账号或者密码错误"})
	}
	// Create token
	t, err := commons.GetToken(user.Id, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, commons.Resp{Code: -1, Message: "加密错误"})
	}
	return c.JSON(http.StatusOK, commons.Resp{Code: 0, Data: map[string]string{"token": t}})
}
func ExistUser(c echo.Context) error {
	username := c.FormValue("username")
	if username == "" {
		return c.JSON(http.StatusOK, commons.Resp{Code: -1, Message: "username must be"})
	}
	userService := &service.UserServiceImpl{}
	has := userService.ExistUser(username)
	data := map[string]interface{}{
		"": has,
	}
	return c.JSON(http.StatusOK, commons.Resp{Code: 0, Data: data})
}
func Register(c echo.Context) error {
	param := &vo.RegisterForm{}
	if err := c.Bind(param); err != nil {
		return c.JSON(http.StatusOK, commons.Resp{Code: -1, Message: "服务器错误"})
	}
	if err := c.Validate(param); err != nil {
		return c.JSON(http.StatusOK, commons.Resp{Code: -1, Message: "参数错误"})
	}
	//fmt.Println(param)
	if !commons.VerifyCode(param.CId, param.Code) {
		return c.JSON(http.StatusOK, commons.Resp{Code: -1, Message: "验证码错误"})
	}

	userService := &service.UserServiceImpl{}
	id, err := userService.Register(param)
	if err != nil {
		return c.JSON(http.StatusOK, commons.Resp{Code: -1, Message: "服务器错误", Data: err})
	}
	data := map[string]interface{}{
		"id": id,
	}
	return c.JSON(http.StatusOK, commons.Resp{Code: 0, Data: data})
}

//base64Captcha create http handler

func Captcha(c echo.Context) error {
	//parse request parameters
	color := &color.RGBA{
		uint8(255 * rand.Int()),
		uint8(255 * rand.Int()),
		uint8(255 * rand.Int()),
		uint8(255 * rand.Int()),
	}
	fonts := []string{
		"wqy-microhei.ttc", "ApothecaryFont.ttf",
	}

	//create base64 encoding captcha
	captcha := base64Captcha.NewCaptcha(base64Captcha.NewDriverString(80, 240, 15, 100, 4, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", color, fonts), commons.Store)
	id, b64s, err := captcha.Generate()
	if err != nil {
		return err
	}
	println(commons.Store.Get(id, false))
	data := &dto.CaptchaResp{CaptchaId: id, Base64: b64s}
	return c.JSON(http.StatusOK, commons.Resp{Code: 0, Data: data})
}

func VerifyCode(c echo.Context) error {
	param := new(vo.CaptchaBody)
	if err := c.Bind(param); err != nil {
		return c.JSON(http.StatusBadRequest, commons.Resp{Code: -1, Message: "服务器错误"})
	}
	fmt.Println(param)
	if err := c.Validate(param); err != nil {
		return c.JSON(http.StatusBadRequest, commons.Resp{Code: -1, Message: "参数错误"})
	}
	flag := strings.EqualFold(commons.Store.Get(param.Id, false), param.VerifyValue)
	if flag {
		return c.JSON(http.StatusOK, commons.Resp{Code: 0, Data: flag})

	}
	return c.JSON(http.StatusOK, commons.Resp{Code: -1, Message: "验证码错误"})
}

func FindInfo(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	//fmt.Println(token)

	claim := token.Claims.(*commons.JwtCustomClaims)
	userService := &service.UserServiceImpl{}
	userInfo, err := userService.FindInfo(claim.Id)
	if err != nil {
		return c.JSON(http.StatusOK, commons.Resp{Code: 0, Message: "查找用户信息失败"})
	}
	return c.JSON(http.StatusOK, commons.Resp{Code: 0, Data: userInfo})
}

func UpdateInfo(c echo.Context) error {
	info := &bean.UserInfo{}
	if err := c.Bind(info); err != nil {
		return err
	}
	//fmt.Println(info)
	userService := &service.UserServiceImpl{}
	_, err := userService.UpdateUserInfo(info)
	if err != nil {
		return c.JSON(http.StatusOK, commons.Resp{Code: 0, Message: "查找用户信息失败"})
	}
	return c.JSON(http.StatusOK, commons.Resp{Code: 0})
}

func ConfirmPassword(c echo.Context) error {
	pwd := c.FormValue("password")
	userService := &service.UserServiceImpl{}
	token := c.Get("user").(*jwt.Token)
	claim := token.Claims.(*commons.JwtCustomClaims)
	user, err := userService.ConfirmPassword(claim.Id, claim.UserName, pwd)
	if !user || err != nil {
		return c.JSON(http.StatusOK, commons.Resp{Code: -1, Message: "密码错误"})
	}
	id, err := uuid.New()
	if err != nil {
		return c.JSON(http.StatusOK, commons.Resp{Code: -1, Message: "UUID生成错误"})
	}
	commons.Store.Set(id, fmt.Sprintf("%v", claim.Id))
	data := map[string]string{
		"uuid": id,
	}
	return c.JSON(http.StatusOK, commons.Resp{Code: 0, Data: data})
}

func UpdatePwd(c echo.Context) error {
	pwd := c.FormValue("password")
	uuid := c.FormValue("uuid")
	token := c.Get("user").(*jwt.Token)
	claim := token.Claims.(*commons.JwtCustomClaims)
	uId := claim.Id
	if commons.Store.Get(uuid, true) != fmt.Sprintf("%v", uId) {
		return c.JSON(http.StatusOK, commons.Resp{Code: -1, Message: "秘钥错误"})
	}
	userService := &service.UserServiceImpl{}
	user := &bean.User{
		Id:       uId,
		Password: pwd,
		Username: claim.UserName,
	}
	count, err := userService.UpdateUser(user)
	if count <= 0 && err != nil {
		return c.JSON(http.StatusOK, commons.Resp{Code: -1, Message: "没有该用户"})
	}
	return c.JSON(http.StatusOK, commons.Resp{Code: 0})
}

