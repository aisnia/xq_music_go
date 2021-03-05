package dao

import (
	"xq_music_go/conf/db"
	"xq_music_go/models/bean"
)

type UserInfoDao interface {
	AddUserInfo(UserInfo *bean.UserInfo) (int64, error)
	DeleteUserInfoById(id int) error
	FindUserInfo(info *bean.UserInfo) (bool, error)
	FindUserInfoByIds(ids []int) []*bean.UserInfo
	UpdateUserInfo(info *bean.UserInfo) (int64, error)
	FindUser(user *bean.User) (bool, error)
}

type MysqlUserInfoDao struct {
}

func (*MysqlUserInfoDao) AddUserInfo(userInfo *bean.UserInfo) (int64, error) {
	return db.Engine.InsertOne(userInfo)
}
func (*MysqlUserInfoDao) DeleteUserInfoById(id int) error {
	return nil
}
func (*MysqlUserInfoDao) FindUserInfo(info *bean.UserInfo) (bool, error) {
	return db.Engine.Get(info)
}

func (*MysqlUserInfoDao) FindUserInfoByIds(ids []int) []*bean.UserInfo {
	return nil
}

func (*MysqlUserInfoDao) UpdateUserInfo(info *bean.UserInfo) (int64, error) {
	return db.Engine.ID(info.Id).Cols("name", "phone", "email", "sex").Update(info)
}

func (*MysqlUserInfoDao) FindUser(user *bean.User) (bool, error) {
	return db.Engine.Get(user)
}
