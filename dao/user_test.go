package dao

import (
	"fmt"
	"testing"
	"xq_music_go/conf/db"
)

func init() {
	db.InitXorm()
}
func TestMysqlUserDao(t *testing.T) {
	//dao := &MysqlUserDao{}
	//user := &bean.User{Username: "xiaoqiang", Password: "123456"}
	//
	//fmt.Println(dao.AddUser(user))
	//fmt.Println(user)
	//dao := &SongMysqlDao{}
	//songs, _ := dao.FindSongs(1, 7)
	//fmt.Println(songs)

	//service := service.SongServiceImpl{}
	//songs, count, _ := service.GetHotMusic(1)
	songDao := &SongMysqlDao{}
	ids := []int64{1}
	songs,_ := songDao.FindByIds(ids)
	fmt.Println(songs)
}
