package goSqlHelper


import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"errors"
)

type HelperRow map[string] interface{}

func (this *HelperRow) ToJson() string{
	bs,err := json.Marshal(&this)
	if err!=nil {
		panic(err)
	}
	return string(bs)
}
/**
字段串获取key
*/
func (this *HelperRow) String(key string) (*string,error){
	var obj interface{}
	obj = (*this)[key]
	if obj== nil {
		return nil,nil
	}
	var str string
	switch obj.(type) {
	case string:
		str= obj.(string)
	case int:
		str= strconv.Itoa(obj.(int))
	case int32:
		str = Int32ToStr(obj.(int32))
	case int64:
		str= Int64ToStr(obj.(int64))
	case float64:
		str = Float64ToStr(obj.(float64))
	case float32:
		str= Float64ToStr(float64(obj.(float32)))
	default:
		return nil,errors.New("don't konw type:"+reflect.TypeOf(obj).Name())
	}

	str = fmt.Sprintf("%V",obj)
	return &str,nil
}


func (this *HelperRow) PString(key string) *string{
	str,err:=this.String(key)
	if err!=nil {
		panic(err)
	}
	return str
}
func (this *HelperRow) CleverString(key string) string{
	str,err:=this.String(key)
	if err!=nil {
		panic(err)
	}
	if str==nil {
		return ""
	}
	return *str
}


func (this *HelperRow) PInt(key string) *int{
	val,err:=this.Int(key)
	if(err!=nil){
		panic(err)
	}
	return val
}
func (this *HelperRow) PInt64(key string) *int64{
	val,err:=this.Int64(key)
	if(err!=nil){
		panic(err)
	}
	return val
}

/**
	int获取key
*/
func (this *HelperRow) Int(key string) (val *int,err error){
	var obj interface{}
	obj = (*this)[key]

	if obj== nil {
		return nil,nil
	}
	var ret int


	switch obj.(type) {
		case string:
			ret,err = strconv.Atoi(obj.(string))
		case int:
			ret = obj.(int)
		case int32:
			ret =  int(obj.(int32))
		case int64:
			ret = int(obj.(int64))
		default: return nil,errors.New("convert to int error")
	}
	return &ret,err
}


/**
	int64获取key
*/
func (this *HelperRow) Int64(key string) (val *int64,err error){
	var obj interface{}
	obj = (*this)[key]

	if obj== nil {
		return nil,nil
	}

	var ret int64

	switch obj.(type) {
		case string:
			ret,err = StrToInt64(obj.(string))
		case int:
			ret = int64(obj.(int))
		case int32:
			ret =  int64(obj.(int32))
		case int64:
			ret= obj.(int64)
		default : err=errors.New("convert to int64 error")
	}
	return &ret,err
}




