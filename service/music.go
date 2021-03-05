package service

import (
	"fmt"
	"strconv"
	"xq_music_go/commons"
	"xq_music_go/dao"
	"xq_music_go/models/bean"
	"xq_music_go/models/dto"
)

type SongService interface {
	UploadSong(song *bean.Song) (int64, error)
	GetHotMusic(page int) ([]dto.Song, int64, error)
	SetMyLove(songId string, uId int64) (int64, error)
	FindMyLove(id int64) ([]bean.UserLove, error)
	RemoveMyLove(songId string, uId int64) (int64, error)
	GetLoveMusic(page int, id int64) ([]dto.Song, int64, error)
}
type SingerService interface {
	UploadSinger(singer *bean.Singer) (int64, error)
}
type SongServiceImpl struct {
}

func (s SongServiceImpl) UploadSong(song *bean.Song) (int64, error) {
	songDao := &dao.SongMysqlDao{}
	return songDao.AddSong(song)
}

func (s SongServiceImpl) GetHotMusic(page int) ([]dto.Song, int64, error) {
	songDao := &dao.SongMysqlDao{}
	count, err := songDao.SongCount()
	if err != nil {
		return nil, 0, err
	}
	songs, err := songDao.FindSongs(page, commons.PAGE_SIZE)
	if err != nil {
		return nil, 0, err
	}
	data := make([]dto.Song, 0)
	for i, _ := range songs {
		d := dto.Song{}
		_ = commons.StructToStruct(&songs[i], &d, "json")
		data = append(data, d)
	}
	return data, count, nil
}

func (s SongServiceImpl) SetMyLove(songId string, uId int64) (int64, error) {
	id, err := strconv.Atoi(songId)
	if err != nil {
		return -1, err
	}
	love := &bean.UserLove{
		UId:    uId,
		SongId: int64(id),
	}
	songDao := &dao.SongMysqlDao{}
	return songDao.AddLove(love)
}

func (s SongServiceImpl) FindMyLove(id int64) ([]bean.UserLove, error) {
	songDao := &dao.SongMysqlDao{}
	return songDao.FindMyLove(id)
}

func (s SongServiceImpl) RemoveMyLove(songId string, uId int64) (int64, error) {
	songDao := &dao.SongMysqlDao{}
	id, err := strconv.Atoi(songId)
	if err != nil {
		return -1, err
	}
	love := &bean.UserLove{
		UId:    uId,
		SongId: int64(id),
	}
	return songDao.DeleteMyLove(love)
}

func (s SongServiceImpl) GetLoveMusic(page int, id int64) ([]dto.Song, int64, error) {
	songDao := &dao.SongMysqlDao{}
	count, err := songDao.LoveCount()
	if err != nil {
		return nil, 0, fmt.Errorf("count error")
	}
	userLoves, err := songDao.FindLoveByUId(id, page, commons.PAGE_SIZE)

	ids := make([]int64, 0)
	for i, _ := range userLoves {
		ids = append(ids, userLoves[i].SongId)
	}
	songs, err := songDao.FindByIds(ids)
	if err != nil {
		return nil, 0, fmt.Errorf("find music error")
	}
	data := make([]dto.Song, 0)
	for i := range songs {
		s := dto.Song{IsLove: songs[i].Id}
		_ = commons.StructToStruct(&songs[i], &s, "json")
		data = append(data, s)
	}
	return data,count,nil
}

func (s SongServiceImpl) FindMusicByName(name string, page int)  ([]dto.Song, int64, error) {
	songDao := &dao.SongMysqlDao{}
	count, err := songDao.SongCount()
	if err != nil {
		return nil, 0, err
	}
	songs, err := songDao.FindSongsLikeName(name,page, commons.PAGE_SIZE)
	if err != nil {
		return nil, 0, err
	}
	data := make([]dto.Song, 0)
	for i, _ := range songs {
		d := dto.Song{}
		_ = commons.StructToStruct(&songs[i], &d, "json")
		data = append(data, d)
	}
	return data, count, nil
}

type SingerServiceImpl struct {
}

func (s SingerServiceImpl) UploadSinger(singer *bean.Singer) (int64, error) {
	singerDao := &dao.SingerMysqlDao{}
	return singerDao.AddSinger(singer)
}
