package goSqlHelper

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type SqlHelper struct{
	Connection *sql.DB
}

const QUERY_BUFFER_SIZE=20


func MysqlOpen(connectionStr string) (*SqlHelper,error){

	sqlHelper :=new (SqlHelper)
	err:= sqlHelper.Open(connectionStr)
	if(err!=nil){
		return nil ,err
	}
	return sqlHelper,nil
}

func New(connectionStr string) (*SqlHelper,error){
	return MysqlOpen(connectionStr)
}

/**
@todo need add context,Transaction support
begin transaction
 */
func (this *SqlHelper) Begin()(*sql.Tx,error){
	val, err:= this.Connection.Begin()
	if err!=nil {
		return nil, err
	}
	return val,nil
}


/**
   初始化模块
*/
func (this *SqlHelper) Open(connectionStr string) error{
	var err error
	
//	sql.Open
	this.Connection,err = sql.Open("mysql",connectionStr)
	if(err!=nil){
		return errors.New(fmt.Sprintf("数据库链接失败:%s",err.Error()))
	}
	err=this.Connection.Ping();
	if err!=nil {
		return err
	}
	return nil
}

/**
初始化
*/
func (this *SqlHelper) SetDB (conn *sql.DB) {
		this.Connection=conn
}

/**
  读取多行
*/
func (this *SqlHelper) QueryRows(sql string, args ...interface{})([]HelperRow, error) {

	var rows =make([]HelperRow, 0, QUERY_BUFFER_SIZE)
	query,err:= this.Querying(sql,args...)
	if err!=nil {
		return nil, err
	}
	defer query.Close()
	for{
		row,err:= query.QueryRow();
		if err==nil {
			rows=append(rows,row)
			continue
		}
		if err== NoFoundError {
			break
		}
		return nil , err
	}

	return rows,nil
}

/**
  读取多行
*/
func (this *SqlHelper) QueryTable(sql string, args ...interface{})(*HelperTable, error) {

	var rows =make([]HelperRow,0,QUERY_BUFFER_SIZE)
	query,err:= this.Querying(sql,args...)
	if err!=nil {
		return nil, err
	}
	defer query.Close()
	cols,err:=query.Columns()
	if err!=nil {
		return nil ,err
	}
	for{
		row,err:= query.QueryRow();
		if err==nil {
			rows=append(rows,row)
			continue
		}
		if err== NoFoundError {
			break
		}
		return nil,err
	}
	return NewTable(rows,cols),nil
}


/**
get Querying handler
 */
func (this *SqlHelper) Querying(sql string,args ...interface{})(*Querying,error){

	rows,err := this.Connection.Query(sql,args...)
	if err!=nil {
		return nil, err
	}
	querying:= NewQuerying(rows)
	return querying,nil
}


/**
  读取一行
*/
func (this *SqlHelper) QueryRow(sql string, args ...interface{})(HelperRow, error) {

	query,err:= this.Querying(sql,args...)
	defer query.Close()
	if err!=nil {
		return nil, err
	}
	row,err:= query.QueryRow();
	if err!=nil {
		return nil ,err
	}
	if row == nil {
		return nil,nil
	}
	return row,nil
}
/**
  读取个值
*/
func (this *SqlHelper) QueryScalarInt(sql string, args ...interface{})(int, error) {
	
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
func (this *SqlHelper) Exec(sql string,args ...interface{})(sql.Result,error){
	stmt,err:=this.Connection.Prepare(sql)
	if err!=nil {
		return nil, err
	}
	defer stmt.Close()
	result,err := stmt.Exec(args...)
	if(err!=nil){
		return nil, err
	}
	return result,nil
}

/*
执行插入sql
*/
func (this *SqlHelper) Insert(sql string, args ...interface{})(int64,error){
	result,err := this.Exec(sql,args...)
	if err!=nil {
		return 0,err
	}
	
	id,err2 := result.LastInsertId()
	if err2 != nil {
		return 0, err2
	}
	return id , nil
}
/*
更新或删除sql
*/
func (this *SqlHelper) UpdateOrDel(sql string, args ...interface{})(int64,error){
	result,err := this.Exec(sql,args...)
	if err!=nil {
		return 0,err
	}
	
	cnt,err2 := result.RowsAffected()
	if(err2 != nil) {
		return 0, err2
	}
	return cnt , nil
}


/*
    关闭连接
*/
func (this *SqlHelper) Close() error{
	err := this.Connection.Close()
	return err
}