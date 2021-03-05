package dao

import (
	"xq_music_go/conf/db"
	"xq_music_go/models/bean"
)

type UserDao interface {
	AddUser(user *bean.User) (int64, error)
	DeleteUserById(id int) error
	FindUser(*bean.User) (bool, error)
	FindUserByIds(ids []int) []*bean.User
	ExistUserByUserName(username string) (bool, error)
	UpdateUser(user *bean.User) (int64, error)
}

type MysqlUserDao struct {
}

func (*MysqlUserDao) AddUser(user *bean.User) (int64, error) {
	return db.Engine.InsertOne(user)
}
func (*MysqlUserDao) DeleteUserById(id int) error {
	return nil
}
func (*MysqlUserDao) FindUser(user *bean.User) (bool, error) {
	return db.Engine.Get(user)
}

func (*MysqlUserDao) FindUserByIds(ids []int) []*bean.User {
	return nil
}

func (*MysqlUserDao) ExistUserByUserName(username string) (bool, error) {
	return db.Engine.Exist(&bean.User{
		Username: username,
	})
}

func (*MysqlUserDao) UpdateUser(user *bean.User) (int64, error) {
	return db.Engine.ID(user.Id).Cols("password").Update(user)
}
