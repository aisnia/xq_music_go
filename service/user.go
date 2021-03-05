package service

import (
	"fmt"
	"xq_music_go/commons"
	"xq_music_go/dao"
	"xq_music_go/models/bean"
	"xq_music_go/models/vo"
)

type UserService interface {
	Register(form *vo.RegisterForm) (int64, error)
	ExistUser(username string) bool
	Login(username, password string) (*bean.User, error)
	FindInfo(id int64) (*bean.UserInfo, error)
	UpdateUserInfo(info *bean.UserInfo) (int64, error)
	ConfirmPassword(u_id int64, password string) (bool, error)
	UpdateUser(user *bean.User) (int64, error)
}
type UserServiceImpl struct {
}

func (*UserServiceImpl) Register(form *vo.RegisterForm) (int64, error) {
	user := &bean.User{Username: form.Username, Password: commons.Md5(form.Username+commons.PROJECT_NAME+form.Password, commons.PWD_SECRET_KEY)}
	info := &bean.UserInfo{}
	commons.StructToStruct(form, info, "json")
	userDao := &dao.MysqlUserDao{}
	_, err := userDao.AddUser(user)
	if err != nil {
		return -1, fmt.Errorf("添加用户错误")
	}
	info.UId = user.Id
	infoDao := &dao.MysqlUserInfoDao{}
	_, err = infoDao.AddUserInfo(info)
	if err != nil {
		return -1, fmt.Errorf("用户信息添加错误")
	}
	// 权限管理
	userRoleDao := &dao.MysqlUserRoleDao{}
	userRoleDao.AddUserDefaultRole(user.Id)
	return user.Id, nil
}

func (*UserServiceImpl) ExistUser(username string) bool {
	userDao := dao.MysqlUserDao{}
	has, err := userDao.ExistUserByUserName(username)
	if err != nil || has {
		return true
	}
	return false
}

func (*UserServiceImpl) Login(username, password string) (*bean.User, error) {
	userDao := dao.MysqlUserDao{}
	user := &bean.User{Username: username}
	has, err := userDao.FindUser(user)
	password = commons.Md5(username+commons.PROJECT_NAME+password, commons.PWD_SECRET_KEY)
	if has && err == nil && user.Password == password {
		return user, nil
	}
	return nil, err

}

func (*UserServiceImpl) FindInfo(id int64) (*bean.UserInfo, error) {
	userInfoDao := &dao.MysqlUserInfoDao{}
	userInfo := &bean.UserInfo{UId: id}
	has, err := userInfoDao.FindUserInfo(userInfo)
	if has && err == nil {
		return userInfo, nil
	}
	return nil, err
}

func (*UserServiceImpl) UpdateUserInfo(info *bean.UserInfo) (int64, error) {
	userInfoDao := &dao.MysqlUserInfoDao{}
	return userInfoDao.UpdateUserInfo(info)
}

func (*UserServiceImpl) ConfirmPassword(u_id int64, username, password string) (bool, error) {
	userInfoDao := &dao.MysqlUserInfoDao{}
	password = commons.Md5(username+commons.PROJECT_NAME+password, commons.PWD_SECRET_KEY)
	return userInfoDao.FindUser(&bean.User{Id: u_id, Password: password})
}

func (*UserServiceImpl) UpdateUser(user *bean.User) (int64, error) {
	userDao := &dao.MysqlUserDao{}
	password := commons.Md5(user.Username+commons.PROJECT_NAME+user.Password, commons.PWD_SECRET_KEY)
	user.Password = password
	return userDao.UpdateUser(user)
}
