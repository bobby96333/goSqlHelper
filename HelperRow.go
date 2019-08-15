package goSqlHelper

import (
	"encoding/json"
	"github.com/bobby96333/commonLib/stackError"
	"reflect"
	"strconv"
)

type HelperRow map[string] interface{}

func (this HelperRow) ToJson() (string,error){
	bs,err := json.Marshal(this)
	if err!=nil {
		return "",err
	}
	return string(bs),nil
}
/**
字段串获取key
*/
func (this HelperRow) String(key string) (string,error){
	var obj interface{}
	obj = (this)[key]
	if obj== nil {
		return "", NoFoundError
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
		return "",stackError.New("don't konw type:"+reflect.TypeOf(obj).Name())
	}

	//str = fmt.Sprintf("%V",obj)
	return str,nil
}


func (this HelperRow) PString(key string) (string){
	str,err:=this.String(key)
	if err==nil {
		return str
	}
	if  err == NoFoundError {
		return ""
	}
	panic(err)
}
func (this HelperRow) CleverString(key string) string{
	str,err:=this.String(key)
	if err == nil {
		return str
	}
	if err == NoFoundError {
		return ""
	}
	return str
}


func (this HelperRow) PInt(key string) int{
	val,err:=this.Int(key)
	if err==nil {
		return val
	}
	if err== NoFoundError {
		return 0
	}
	panic(err)
}
func (this HelperRow) PInt64(key string) int64{
	val,err:=this.Int64(key)
	if err == nil {
		return val
	}
	if err == NoFoundError {
		return 0
	}
	panic(err)
}

/**
	int获取key
*/
func (this HelperRow) Int(key string) (int,error){
	var obj interface{}
	obj = this[key]

	if obj== nil {
		return 0, NoFoundError
	}
	var ret int
	var converr error
	switch obj.(type) {
		case string:
			ret,converr = strconv.Atoi(obj.(string))
		case int:
			ret = obj.(int)
		case int32:
			ret =  int(obj.(int32))
		case int64:
			ret = int(obj.(int64))
		default: return 0,stackError.New("convert to int error")
	}
	if converr!=nil {
		return 0,converr
	}
	return ret,nil
}


/**
	int64获取key
*/
func (this HelperRow) Int64(key string) (int64,error){
	var obj interface{}
	obj = (this)[key]

	if obj== nil {
		return 0, NoFoundError
	}

	var ret int64
	var converr error

	switch obj.(type) {
		case string:
			ret,converr = StrToInt64(obj.(string))
		case int:
			ret = int64(obj.(int))
		case int32:
			ret =  int64(obj.(int32))
		case int64:
			ret= obj.(int64)
		default : return 0,stackError.New("convert to int64 error")
	}
	if converr != nil {
		return 0,converr
	}
	return ret,nil
}




