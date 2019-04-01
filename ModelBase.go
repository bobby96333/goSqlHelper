package goSqlHelper

import (
	"fmt"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

type Helper struct{
	Connection *sql.DB
}


/**
   初始化模块
*/
func (this *Helper) Open(connectionStr string) error{
	
	var err error
	
//	sql.Open
	this.Connection,err = sql.Open("mysql",connectionStr)
	if(err!=nil){
		return errors.New(fmt.Sprintf("数据库链接失败:%s",err.Error()))
	}
	
	return nil
}

/**
初始化
*/
func (this *Helper) SetDB (conn *sql.DB) {
		this.Connection=conn
}

/**
  读取多行
*/
func (this *Helper) QueryRows(sql string, args ...interface{})(*[]HelperRow, error) {
	
	rows,err := this.Connection.Query(sql,args...)
		
	if(err!=nil){
		return nil,err
	}
	defer rows.Close()
	var ret []HelperRow

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {

		err = rows.Scan(scanArgs...)
		record := make(HelperRow)
		for i, col := range values {
			if col == nil {
				continue
			}
			switch col.(type) {
				
				case []byte:
					record[columns[i]] = string(col.([]byte))
				default:
					record[columns[i]] = col
				
			}
		}
		ret = append(ret,record)
	}

	return &ret,nil
}


/**
  读取一行
*/
func (this *Helper) QueryRow(sql string, args ...interface{})(*HelperRow, error) {
	
	rows,err := this.Connection.Query(sql,args...)
		
	if(err!=nil){
		return nil,err
	}

	defer rows.Close()

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	if rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		record := make(HelperRow)
		for i, col := range values {
			if col == nil {
				continue
			}
			switch col.(type) {
				
				case []byte:
					record[columns[i]] = string(col.([]byte))
				default:
					record[columns[i]] = col
				
			}
			
		}
		return &record,nil
	}

	return nil,errors.New(fmt.Sprintf("no found row:%s",sql))
}
/**
  读取个值
*/
func (this *Helper) QueryScalarInt(sql string, args ...interface{})(int, error) {
	
	rows,err := this.Connection.Query(sql,args...)
	if(err!=nil){
		
		return 0, err
	}
	defer rows.Close()
	if rows.Next() {
		var val int
		err = rows.Scan(&val)
		return val,nil
	}

	return 0, errors.New("no found record.")
}


/*
执行sql
*/
func (this *Helper) Exec(sql string,args ...interface{})(sql.Result,error){
	stmt,err:=this.Connection.Prepare(sql)
	if(err!=nil){
		return nil,err
	}
	defer stmt.Close()
	result,exeerr := stmt.Exec(args...)
	if(exeerr!=nil){
		return nil,exeerr
	}
	return result,nil
}

/*
执行插入sql
*/
func (this *Helper) Insert(sql string, args ...interface{})(int64,error){
	result,err := this.Exec(sql,args...)
	if(err!=nil){
		return 0,err
	}
	
	id,err2 := result.LastInsertId()
	if(err2 != nil) {
		return 0,err2
	}
	return id , nil
}
/*
更新或删除sql
*/
func (this *Helper) UpdateOrDel(sql string, args ...interface{})(int64,error){
	result,err := this.Exec(sql,args...)
	if(err!=nil){
		return 0,err
	}
	
	cnt,err2 := result.RowsAffected()
	if(err2 != nil) {
		return 0,err2
	}
	return cnt , nil
}


/*
    关闭连接
*/
func (this *Helper) Close() error{
	err := this.Connection.Close()
	return err
}