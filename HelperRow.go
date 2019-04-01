package goSqlHelper


import (
	"encoding/json"
	"fmt"
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
func (this *HelperRow) String(key string) string{
	var obj interface{}
	obj = (*this)[key]
	
	switch obj.(type) {
		case string:
			return obj.(string)
		case int:
			return strconv.Itoa(obj.(int))
		case int32:
			return  Int32ToStr(obj.(int32))
		case int64:
			return Int64ToStr(obj.(int64))
		case float64:
			return Float64ToStr(obj.(float64))
		case float32:
			return Float64ToStr(float64(obj.(float32)))
		
		
	}
	
	return fmt.Sprintf("%V",obj)
	
	
}


func (this *HelperRow) PInt(key string) int{
	val,err:=this.Int(key)
	if(err!=nil){
		panic(err)
	}
	return val
}
func (this *HelperRow) PInt64(key string) int64{
	val,err:=this.Int64(key)
	if(err!=nil){
		panic(err)
	}
	return val
}

/**
	int获取key
*/
func (this *HelperRow) Int(key string) (int,error){
	var obj interface{}
	obj = (*this)[key]
	
	switch obj.(type) {
		case string:
			return strconv.Atoi(obj.(string))
		case int:
			return obj.(int),nil
		case int32:
			return  int(obj.(int32)),nil
		case int64:
			return int(obj.(int64)),nil	
	}
	return 0,errors.New("convert to int error")
}


/**
	int64获取key
*/
func (this *HelperRow) Int64(key string) (int64,error){
	var obj interface{}
	obj = (*this)[key]
	
	switch obj.(type) {
		case string:
			return StrToInt64(obj.(string))
		case int:
			return int64(obj.(int)),nil
		case int32:
			return  int64(obj.(int32)),nil
		case int64:
			return obj.(int64),nil	
	}
	return 0,errors.New("convert to int64 error")
}



