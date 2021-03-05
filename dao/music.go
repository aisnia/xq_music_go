package dao

import (
	"xq_music_go/conf/db"
	"xq_music_go/models/bean"
)

type SongDao interface {
	AddSong(song *bean.Song) (int64, error)
	FindSongs(page int, size int) ([]bean.Song, error)
	SongCount() (int64, error)
	AddLove(love *bean.UserLove) (int64, error)
}
type SingerDao interface {
	AddSinger(singer *bean.Singer) (int64, error)
}

type SongMysqlDao struct {
}

func (*SongMysqlDao) AddSong(song *bean.Song) (int64, error) {
	return db.Engine.InsertOne(song)
}

func (s *SongMysqlDao) FindSongs(page int, size int) ([]bean.Song, error) {
	songs := make([]bean.Song, 0)
	err := db.Engine.Cols("id", "name", "singer_name", "album", "url", "upload_u_id", "duration").Limit(size, (page-1)*size).Find(&songs)
	return songs, err
}

func (s *SongMysqlDao) SongCount() (int64, error) {
	song := new(bean.Song)
	return db.Engine.Count(song)
}

func (s *SongMysqlDao) AddLove(love *bean.UserLove) (int64, error) {
	return db.Engine.InsertOne(love)
}

func (s *SongMysqlDao) FindMyLove(id int64) ([]bean.UserLove, error) {
	userloves := make([]bean.UserLove, 0)
	err := db.Engine.Where("u_id=?", id).Find(&userloves)
	return userloves, err
}

func (s *SongMysqlDao) DeleteMyLove(love *bean.UserLove) (int64, error) {
	return db.Engine.Where("u_id=?", love.UId).And("song_id=?", love.SongId).Unscoped().Delete(love)
}

func (s *SongMysqlDao) LoveCount() (int64, error) {
	userLove := new(bean.UserLove)
	return db.Engine.Count(userLove)
}

func (s *SongMysqlDao) FindLoveByUId(id int64, page int, size int) ([]bean.UserLove, error) {
	userLoves := make([]bean.UserLove, 0)
	err := db.Engine.Where("u_id=?", id).Limit(size, (page-1)*size).Find(&userLoves)
	return userLoves, err
}

func (s *SongMysqlDao) FindByIds(ids []int64) ([]bean.Song, error) {
	songs := make([]bean.Song, 0)
	err := db.Engine.In("id", ids).Find(&songs)
	return songs, err
}

func (s *SongMysqlDao) FindSongsLikeName(name string, page int, size int) ([]bean.Song, error) {
	songs := make([]bean.Song, 0)
	err := db.Engine.Cols("id", "name", "singer_name", "album", "url", "upload_u_id", "duration").Where("name like ?", "%"+name+"%").Or("singer_name like ?", "%"+name+"%").Limit(size, (page-1)*size).Find(&songs)
	return songs, err
}

type SingerMysqlDao struct {
}

func (*SingerMysqlDao) AddSinger(singer *bean.Singer) (int64, error) {
	return db.Engine.InsertOne(singer)
}
