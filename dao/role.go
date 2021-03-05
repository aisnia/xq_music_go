package dao

import (
	"xq_music_go/conf/db"
	"xq_music_go/models/bean"
)


type RoleDao interface {
	AddUser(Role *bean.Role) (int64, error)
	DeleteRoleById(id int) error
	FindRoleById(id int) *bean.Role
	FindRoleByName(name string) *bean.Role
	FindRoleByIds(ids []int) []*bean.Role
}

type MysqlRoleDao struct {
}

func (*MysqlRoleDao) AddRole(role *bean.Role) (int64, error) {
	return db.Engine.InsertOne(role)
}
func (*MysqlRoleDao) DeleteRoleById(id int) error {
	return nil
}
func (*MysqlRoleDao) FindRoleById(id int) *bean.Role {
	return nil
}
func (*MysqlRoleDao) FindRoleByName(name string) *bean.Role {
	return nil
}
func (*MysqlRoleDao) FindRoleByIds(ids []int) []*bean.Role {
	return nil
}

