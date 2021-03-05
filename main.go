package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"os"
	"time"
	"xq_music_go/commons"
	"xq_music_go/conf/db"
	"xq_music_go/controllers"
	"xq_music_go/mymiddleware"
	"xq_music_go/mymiddleware/myjwt"
)

type DefaultValidator struct {
	validator *validator.Validate
}

func (cv *DefaultValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
func main() {
	db.InitXorm()
	e := echo.New()
	//日志
	logFile, err := os.OpenFile("logs/music.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic("log file can't init")
	}
	//logFile := os.Stdout

	defer logFile.Close()
	e.Logger.SetOutput(logFile)
	e.Logger.SetHeader("${time_rfc3339} ${level}")
	e.Validator = &DefaultValidator{validator: validator.New()}
	//Server
	s := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}
	e.HideBanner = false
	e.Static("/static", "static")
	// Middleware
	//恢复
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.CORS())
	//日志
	e.Use(mymiddleware.LoggerWithConfig(mymiddleware.LoggerConfig{
		Format: "id=${id},status=${status},time=${time_rfc3339_nano},remote_ip=${remote_ip},host=${host},method=${method},uri=${uri},latency=${latency},query=${query},form=${form}\n",
		Output: logFile,
	}))

	// Routes
	userRoute := e.Group("/user")
	//登录后才能进行操作的
	userRoute.Use(myjwt.JWTWithConfig(myjwt.JWTConfig{
		SigningKey:  []byte(commons.SECRET),
		Claims:      &commons.JwtCustomClaims{},
		TokenLookup: "form:token",
	}))

	e.POST("/login", controllers.Login)
	e.POST("/register", controllers.Register)
	e.POST("/captcha", controllers.Captcha)
	e.POST("/verifyCode", controllers.VerifyCode)
	e.POST("/existUser", controllers.ExistUser)

	userRoute.POST("/findInfo", controllers.FindInfo)
	userRoute.POST("/updateInfo", controllers.UpdateInfo)
	userRoute.POST("/confirmPassword", controllers.ConfirmPassword)
	userRoute.POST("/updatePwd", controllers.UpdatePwd)
	userRoute.POST("/upload", controllers.Upload)
	userRoute.POST("/getHotMusic", controllers.GetHotMusic)
	userRoute.POST("/setMyLove", controllers.SetMyLove)
	userRoute.POST("/removeMyLove", controllers.RemoveMyLove)
	userRoute.POST("/getLoveMusic", controllers.GetLoveMusic)
	userRoute.POST("/findMusicByName", controllers.FindMusicByName)
	// Start server
	e.Logger.Fatal(e.StartServer(s))
}
