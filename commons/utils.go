package commons

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
)

func GetFieldByTag(data interface{}, tagName string, filterZeroVal bool) (map[string]interface{}, error) {
	fields := make(map[string]interface{})

	value := reflect.ValueOf(data).Elem()
	typ := reflect.TypeOf(data).Elem()

	//类型为struct
	kd := value.Kind()
	if reflect.Struct != kd {
		return nil, errors.New("the param should be struct")
	}

	filedNum := value.NumField()
	for i := 0; i < filedNum; i++ {
		val := value.Field(i).Interface()

		t := value.Field(i).Type()
		//判断字段是什么类型 过滤
		switch t.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			//整型 判断是否过滤0值
			if filterZeroVal {
				vStr := fmt.Sprintf("%v", val)
				if 0 == strings.Compare(vStr, "0") {
					//跳出本次循环不加入map
					continue
				}
			}
		case reflect.String:
			if filterZeroVal {
				vStr := val.(string)
				if "" == vStr {
					continue
				}
			}
		case reflect.Interface:
			if value.Field(i).IsNil() {
				continue
			}
		case reflect.Ptr:
			if value.Field(i).IsNil() {
				continue
			}
		default:
			//不支持float类型
			continue
		}
		key := typ.Field(i).Tag.Get(tagName)
		if "" != key {
			fields[key] = val
		}
	}
	return fields, nil
}
func GetMapsByTag(data interface{}, tagName string, filterZeroVal bool) ([]map[string]interface{}, error) {
	res := make([]map[string]interface{}, 0)
	value := reflect.ValueOf(data).Elem()

	//类型为struct
	kd := value.Kind()
	if reflect.Slice != kd {
		return nil, errors.New("the param should be struct")
	}

	for i := 0; i < value.Len(); i++ {
		v := value.Index(i).Interface()
		m, err := GetFieldByTag(v, tagName, filterZeroVal)
		if err != nil {
			return nil, err
		}
		res = append(res, m)
	}
	return res, nil
}

func MapToStruct(fields map[string]interface{}, data interface{}, tagName string) error {
	value := reflect.ValueOf(data).Elem()
	typ := reflect.TypeOf(data).Elem()
	kind := value.Kind()

	if reflect.Struct != kind {
		return errors.New("the data should be struct")
	}

	fieldsNum := value.NumField()
	for i := 0; i < fieldsNum; i++ {
		key := typ.Field(i).Tag.Get(tagName)
		if v, ok := fields[key]; ok {
			value.Field(i).Set(reflect.ValueOf(v))
		}
	}
	return nil
}

//必须都是指针 Ptr
func StructToStruct(from interface{}, to interface{}, tagName string) error {
	fields, err := GetFieldByTag(from, tagName, true)
	if err != nil {
		return err
	}
	return MapToStruct(fields, to, tagName)
}

func SliceToSlice(from interface{}, to interface{}, tagName string) error {
	if reflect.TypeOf(from).Kind() == reflect.Slice && reflect.TypeOf(to).Kind() == reflect.Slice {
		f := reflect.ValueOf(from)
		t := reflect.ValueOf(to)
		for i := 0; i < f.Len(); i++ {
			fv := f.Index(i).Interface()
			tv := t.Index(i).Interface()
			err := StructToStruct(fv, tv, tagName)
			if err != nil {
				return err
			}

		}
	}
	return nil
}

func VerifyCode(CId, code string) bool {
	return strings.EqualFold(Store.Get(CId, false), code)
}

func Md5(data, secret string) string {
	m := md5.New()
	io.WriteString(m, data+secret)
	return fmt.Sprintf("%x", m.Sum(nil))
}

const (
	PAGE_SIZE = 7
)

type Page struct {
	PageNo     int         `json:"page_no"`     //当前页码
	PageSize   int         `json:"page_size"`   //一页的大小
	TotalPage  int         `json:"total_page"`  //总页数
	TotalCount int         `json:"total_count"` //总记录数
	FirstPage  bool        `json:"first_page"`  //第一页
	LastPage   bool        `json:"last_page"`   //最后一页
	List       interface{} `json:"list"`        //数据data
}

//count 是总数
//pageNo 是当前页数
//pageSIze 是一页的大小
//list是数据
func PageUtil(count int, pageNo int, pageSize int, list interface{}) Page {
	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count/pageSize + 1
	}
	return Page{PageNo: pageNo, PageSize: pageSize, TotalPage: tp, TotalCount: count, FirstPage: pageNo == 1, LastPage: pageNo == tp, List: list}
}
