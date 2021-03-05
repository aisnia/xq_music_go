package commons

import (
	"fmt"
	"testing"
	"xq_music_go/models/bean"
	"xq_music_go/models/vo"
)

func TestGetFieldByTag(t *testing.T) {
	user := &bean.User{Username: "xq", Password: "123456"}
	field, _ := GetFieldByTag(user, "json", true)
	fmt.Println(field)
}

func TestMapToStruct(t *testing.T) {
	field := make(map[string]interface{})
	field["username"] = "xq"
	field["password"] = "123456"
	user := &bean.User{}
	MapToStruct(field, user, "json")
	fmt.Println(user)
}

func TestStructToStruct(t *testing.T) {
	v := &vo.RegisterForm{Username: "xq",Password: "123456"}
	user := &bean.User{}
	StructToStruct(v,user,"json")
	fmt.Println(user)
}

