package bean

import "time"

//用户类
type User struct {
	Id         int64     `json:"id"`
	Username   string    `json:"username" xorm:"username"`
	Password   string    `json:"password" xorm:"password"`
	Status     int8      `json:"status" xorm:"default 0"`
	CreateTime time.Time `json:"create_time" xorm:"created"`
	UpdateTime time.Time `json:"update_time" xorm:"created updated"`
	DeletedAt  time.Time `json:"deleted_at" xorm:"deleted"`
}

func (*User) TableName() string {
	return "user"
}

//用户信息
type UserInfo struct {
	Id           int64     `json:"id"`
	UId          int64     `json:"u_id" xorm:"u_id"`
	Name         string    `json:"name" xorm:"name"`
	Phone        string    `json:"phone" xorm:"phone"`
	Email        string    `json:"email"  xorm:"email"`
	Sex          int8      `json:"sex" xorm:"sex" `
	Avatar       string    `json:"avatar" xorm:"avatar" `
	Birthday     time.Time `json:"birthday" xorm:"birthday" `
	Signature    string    `json:"signature" xorm:"signature" `
	Introduction string    `json:"introduction" xorm:"introduction" `
}

func (*UserInfo) TableName() string {
	return "user_info"
}

type Role struct {
	Id         int64     `xorm:"id" `
	Name       string    `xorm:"name" `
	Status     int8      `xorm:"default 0"`
	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"created updated"`
	DeletedAt  time.Time `json:"deleted_at" xorm:"deleted"`
}

type Permission struct {
	Id         int64     `xorm:"id" `
	Title      string    `xorm:"title" `
	action     string    `xorm:"action" `
	Status     int8      `xorm:"default 0"`
	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"created updated"`
	DeletedAt  time.Time `json:"deleted_at" xorm:"deleted"`
}

type UserRole struct {
	Id         int64     `xorm:"id" `
	UId        int64     `xorm:"u_id" `
	RoleId     int64     `xorm:"role_id" `
	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"created updated"`
	DeletedAt  time.Time `json:"deleted_at" xorm:"deleted"`
}

type RolePermission struct {
	Id           int64
	RoleId       int64
	PermissionId int64
	CreateTime   time.Time `xorm:"created"`
	UpdateTime   time.Time `xorm:"created updated"`
	DeletedAt    time.Time `json:"deleted_at" xorm:"deleted"`
}

type Song struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name" xorm:"name" `
	SingerId   int64     `json:"singer_id" xorm:"singer_id" `
	SingerName string    `json:"singer_name" xorm:"singer_name" `
	Album      string    `json:"album" xorm:"album" `
	Url        string    `json:"url" xorm:"url" `
	Duration   string    `json:"duration" xorm:"duration" `
	UploadUId  int64     `json:"upload_u_id",xorm:"upload_u_id" `
	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"created updated"`
	DeletedAt  time.Time `json:"deleted_at" xorm:"deleted"`
}

type Singer struct {
	Id           int64
	Name         string    `json:"name" xorm:"name" `
	Introduction string    `json:"introduction" xorm:"-" `
	Image        string    `json:"image" xorm:"-" `
	CreateTime   time.Time `xorm:"created"`
	UpdateTime   time.Time `xorm:"created updated"`
	DeletedAt    time.Time `json:"deleted_at" xorm:"deleted"`
}

type UserLove struct {
	Id     int64 `json:"id"`
	UId    int64 `json:"u_id"`
	SongId int64 `json:"song_id"`
}

func (*UserLove) TableName() string {
	return "love"
}
