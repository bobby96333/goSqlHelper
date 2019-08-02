package goSqlHelper

import (
//	"fmt"
	"strconv"
)


/*
	string转float64
*/
func StrToFloat64(val string) (float64 , error) {
	return strconv.ParseFloat(val, 64)
}

/**
	float64 转 string
*/
func Float64ToStr(val float64) (string){
	return strconv.FormatFloat(val, 'f',-1,64)	
}


/*
	string转int64
*/
func StrToInt64(val string) (int64 , error) {
	return strconv.ParseInt(val, 10, 64)
}

/**
	int64 转 string
*/
func Int64ToStr(val int64) (string) {
	return strconv.FormatInt(val, 10)	
}
/*
	string转int32
*/
func StrToInt32(val string) (int32 , error) {
	int64,err:=strconv.ParseInt(val, 10, 64)
	if err!=nil {
		return 0,err
	}
	return int32(int64) ,nil
	
}

/**
	int64 转 string
*/
func Int32ToStr(val int32) (string) {
	return strconv.FormatInt(int64(val), 10)	
}

/*
	获取map的所有键值数组
*/
func map_keys(_map *map[interface{}] interface{}) []interface{} {
	
keys:= make([]interface{},len(*_map)) 
	i := 0
	for key, _ := range *_map {
		keys[i]=key
		i++
	}
	return keys
}

/*
	获取map的所有string键值数组
*/
func map_str_keys(_map *map[string] interface{}) []string {
	
	keys:= make([]string,len(*_map)) 
	i := 0
	for key, _ := range *_map {
		keys[i]=key
		i++
	}
	return keys

}