package dao

import (
	"xq_music_go/conf/db"
	"xq_music_go/models/bean"
)

type UserRoleDao interface {
	AddUserRole(UserRole *bean.UserRole) (int64, error)
	DeleteUserRoleById(id int) error
	FindUserRoleById(id int) *bean.UserRole
	FindUserRoleByName(name string) *bean.UserRole
	FindUserRoleByIds(ids []int) []*bean.UserRole
}

type MysqlUserRoleDao struct {
}

func (*MysqlUserRoleDao) AddUserDefaultRole(uId int64) (int64, error) {
	userRole := bean.UserRole{UId: uId, RoleId: 1}
	return db.Engine.InsertOne(userRole)
}
func (*MysqlUserRoleDao) AddUserRole(userRole *bean.UserRole) (int64, error) {
	return db.Engine.InsertOne(userRole)
}
func (*MysqlUserRoleDao) DeleteUserRoleById(id int) error {
	return nil
}
func (*MysqlUserRoleDao) FindUserRoleById(id int) *bean.UserRole {
	return nil
}
func (*MysqlUserRoleDao) FindUserRoleByName(name string) *bean.UserRole {
	return nil
}
func (*MysqlUserRoleDao) FindUserRoleByIds(ids []int) []*bean.UserRole {
	return nil
}
