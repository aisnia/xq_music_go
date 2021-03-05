package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"xq_music_go/commons"
	"xq_music_go/models/bean"
	"xq_music_go/mymiddleware/mp3"
	"xq_music_go/mymiddleware/uuid"
	"xq_music_go/service"
)

func Upload(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claim := token.Claims.(*commons.JwtCustomClaims)
	id := claim.Id
	name := c.FormValue("music")
	singer := c.FormValue("singer")
	album := c.FormValue("album")
	file, err := c.FormFile("file")
	if name == "" || singer == "" || album == "" || err != nil {
		return c.JSON(http.StatusOK, commons.Resp{Code: -1, Message: "表单信息错误"})
	}
	//打开用户上传的文件
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusOK, commons.Resp{Code: -1, Message: "file is error"})
	}
	defer src.Close()
	fileUUID, err := uuid.New()
	if err != nil {
		return c.JSON(http.StatusOK, commons.Resp{Code: -1, Message: "UUID error"})
	}
	index := strings.LastIndexAny(file.Filename, ".")
	extend := file.Filename[index:]
	dir, _ := os.Getwd()
	//fmt.Println(dir)
	url := commons.MAPPING_Path + commons.MUSIC_PAYH + fileUUID + extend
	dst, err := os.Create(dir + url)
	if err != nil {
		return c.JSON(http.StatusOK, commons.Resp{Code: -1, Message: "file create error"})
	}
	if _, err := io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusOK, commons.Resp{Code: -1, Message: "file copy error"})
	}
	dst.Close()

	//写入到数据库
	//s := &bean.Singer{Name: singer}
	//singerService := &service.SingerServiceImpl{}
	//id, err := singerService.UploadSinger(s)
	//if err != nil {
	//	return c.JSON(http.StatusOK, commons.Resp{Code: -1, Message: "insert singer error"})
	//}

	//歌手id先不管了
	d, err := mp3.Calculate(dir + url)
	if err != nil {
		return c.JSON(http.StatusOK, commons.Resp{Code: -1, Message: "mp3 calculate error" + err.Error()})
	}
	time := fmt.Sprintf("%d:%d", int(d)/60, int(d)%60)
	song := &bean.Song{Name: name, Album: album, Url: url, SingerName: singer, UploadUId: id, Duration: time}
	songServie := &service.SongServiceImpl{}
	_, err = songServie.UploadSong(song)
	if err != nil {
		return c.JSON(http.StatusOK, commons.Resp{Code: -1, Message: "insert song error"})
	}
	return c.JSON(http.StatusOK, commons.Resp{Code: 0})
}

func GetHotMusic(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claim := token.Claims.(*commons.JwtCustomClaims)
	id := claim.Id

	pageStr := c.FormValue("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return c.JSON(http.StatusOK, commons.Resp{Code: -1, Message: "page is not num"})
	}
	songServie := &service.SongServiceImpl{}
	songs, total, err := songServie.GetHotMusic(page)
	if err != nil {
		return c.JSON(http.StatusOK, commons.Resp{Code: -1, Message: "Find Music error"})
	}
	userLoves, err := songServie.FindMyLove(id)
	if err != nil {
		return c.JSON(http.StatusOK, commons.Resp{Code: -1, Message: "My love find error"})
	}
	loves := make(map[int64]int64)
	for i, _ := range userLoves {
		if userLoves[i].UId == id {
			loves[userLoves[i].SongId] = userLoves[i].Id
		}
	}
	for i, _ := range songs {
		sId := songs[i].Id
		if v, ok := loves[sId]; ok {
			//在我喜欢里面则设置isMylove 为对应的id
			songs[i].IsLove = v
		}
	}
	data := commons.PageUtil(int(total), page, commons.PAGE_SIZE, songs)

	return c.JSON(http.StatusOK, commons.Resp{Code: 0, Data: data})
}

func SetMyLove(c echo.Context) error {
	songId := c.FormValue("mId")
	token := c.Get("user").(*jwt.Token)
	claim := token.Claims.(*commons.JwtCustomClaims)
	id := claim.Id

	songServie := &service.SongServiceImpl{}
	count, err := songServie.SetMyLove(songId, id)
	if count <= 0 || err != nil {
		return c.JSON(http.StatusOK, commons.Resp{Code: -1, Message: "Set Love error"})
	}
	return c.JSON(http.StatusOK, commons.Resp{Code: 0})
}
func RemoveMyLove(c echo.Context) error {
	songId := c.FormValue("mId")
	token := c.Get("user").(*jwt.Token)
	claim := token.Claims.(*commons.JwtCustomClaims)
	id := claim.Id

	songServie := &service.SongServiceImpl{}
	_, err := songServie.RemoveMyLove(songId, id)
	if err != nil {
		return c.JSON(http.StatusOK, commons.Resp{Code: 0, Message: "delete error"})
	}
	return c.JSON(http.StatusOK, commons.Resp{Code: 0})
}

func GetLoveMusic(c echo.Context) error {
	pageStr := c.FormValue("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return c.JSON(http.StatusOK, commons.Resp{Code: -1, Message: "page is not num"})
	}
	token := c.Get("user").(*jwt.Token)
	claim := token.Claims.(*commons.JwtCustomClaims)
	id := claim.Id
	songServie := &service.SongServiceImpl{}
	songs, total, err := songServie.GetLoveMusic(page, id)
	if err != nil {
		return c.JSON(http.StatusOK, commons.Resp{Code: -1, Message: "My love find error"})
	}
	data := commons.PageUtil(int(total), page, commons.PAGE_SIZE, songs)
	return c.JSON(http.StatusOK, commons.Resp{Code: 0, Data: data})
}
func FindMusicByName(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claim := token.Claims.(*commons.JwtCustomClaims)
	id := claim.Id

	musicName := c.FormValue("musicName")
	pageStr := c.FormValue("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return c.JSON(http.StatusOK, commons.Resp{Code: -1, Message: "page is not num"})
	}
	songServie := &service.SongServiceImpl{}
	songs, total, err := songServie.FindMusicByName(musicName,page)
	if err != nil {
		return c.JSON(http.StatusOK, commons.Resp{Code: -1, Message: "Find Music error"})
	}
	userLoves, err := songServie.FindMyLove(id)
	if err != nil {
		return c.JSON(http.StatusOK, commons.Resp{Code: -1, Message: "My love find error"})
	}
	loves := make(map[int64]int64)
	for i, _ := range userLoves {
		if userLoves[i].UId == id {
			loves[userLoves[i].SongId] = userLoves[i].Id
		}
	}
	for i, _ := range songs {
		sId := songs[i].Id
		if v, ok := loves[sId]; ok {
			//在我喜欢里面则设置isMylove 为对应的id
			songs[i].IsLove = v
		}
	}
	data := commons.PageUtil(int(total), page, commons.PAGE_SIZE, songs)

	return c.JSON(http.StatusOK, commons.Resp{Code: 0, Data: data})
}
