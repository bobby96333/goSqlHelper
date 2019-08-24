package goSqlHelper

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/bobby96333/commonLib/stackError"
	_ "github.com/go-sql-driver/mysql"
)

type SqlHelper struct{
	Connection *sql.DB
	context context.Context
	tx *sql.Tx
	debugMod bool
	stckErrorPowerId int
}

const QUERY_BUFFER_SIZE=20

/**
@todo no sql

	var obj=new(tb1)
	con.Insert(obj)
	obj.setup(conn)
	obj.Select("id,val").Where("id=2").QueryList()
	sqlHelper.Select("id,val").Where("id=2").QueryList()

*/
func MysqlOpen(connectionStr string) (*SqlHelper,stackError.StackError){

	sqlHelper :=new (SqlHelper)
	err:= sqlHelper.Init("mysql",connectionStr)
	if err!=nil {
		return nil ,stackError.NewFromError(err,-1)
	}
	return sqlHelper,nil
}

func New(connectionStr string) (*SqlHelper,stackError.StackError){
	return MysqlOpen(connectionStr)
}
/**
  open db
*/
func (this *SqlHelper) Init(driver,connectionStr string) stackError.StackError{
	if DefaultDebugModel {
		this.OpenDebug()
	}else{
		this.stckErrorPowerId = -1
	}

	var err error
	//	sql.Open
	this.Connection,err = sql.Open(driver,connectionStr)
	if err!=nil {
		return stackError.New(fmt.Sprintf("db connected failed:%s",err.Error()),this.stckErrorPowerId)
	}
	err=this.Connection.Ping();
	if err!=nil {
		return stackError.NewFromError( err,this.stckErrorPowerId)
	}
	return nil
}

/**
begin context
*/
func (this *SqlHelper) BeginContext(ctx context.Context) *SqlHelperRunner{
	runner :=new(SqlHelperRunner)
	runner.SetDB(this.Connection)
	runner.SetContext(ctx)
	return runner
}

/**
begin a trasnaction
*/
func (this *SqlHelper) Begin() *SqlHelperRunner{
	runner :=new(SqlHelperRunner)
	runner.SetDB(this.Connection)
	runner.Begin()
	return runner
}

/**
print sql and parameter at prepare exeucting
 */
func (this *SqlHelper) OpenDebug(){
	this.debugMod=true
	this.stckErrorPowerId=stackError.GetPowerKey()
	stackError.SetPower(true,this.stckErrorPowerId)

}

/**
begin a trasnaction
*/
func (this *SqlHelper) BeginTx(ctx context.Context, opts *sql.TxOptions) (*SqlHelperRunner,stackError.StackError) {
	runner :=new(SqlHelperRunner)
	runner.SetDB(this.Connection)
	err:= runner.BeginTx(ctx,opts)
	if err!=nil {
		return nil,stackError.NewFromError(err,this.stckErrorPowerId)
	}
	return runner,nil
}


/**
set db object
*/
func (this *SqlHelper) SetDB (conn *sql.DB) {
		this.Connection=conn
}


/**
get Querying handler
*/
func (this *SqlHelper) Querying(sql string,args ...interface{})(*Querying,stackError.StackError){
	if this.debugMod {
		fmt.Println(sql)
		fmt.Println(args)
	}
	var rows ,err = this.query(sql,args...)
	if err!=nil {
		return nil, stackError.NewFromError(err,this.stckErrorPowerId)
	}
	querying:= NewQuerying(rows,this.stckErrorPowerId)
	return querying,nil
}
/**
  read a int value
*/
func (this *SqlHelper) QueryScalar(val interface{} , sql string, args ...interface{}) stackError.StackError  {
	if this.debugMod {
		fmt.Println(sql)
		fmt.Println(args)
	}
	var err error
	rows ,err := this.query(sql,args...)
	if err!=nil {
		return stackError.NewFromError(err,this.stckErrorPowerId)
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(val)
		return stackError.NewFromError(err,this.stckErrorPowerId)
	}
	return NoFoundError
}
/**
  read a int value
*/
func (this *SqlHelper) QueryScalarInt(sql string, args ...interface{})(int,stackError.StackError) {
	var val int
	err :=this.QueryScalar(&val,sql,args...)
	return val, err
}
/**
  read a int value
*/
func (this *SqlHelper) QueryScalarString(sql string, args ...interface{})(string,stackError.StackError) {
	var val string
	err :=this.QueryScalar(&val,sql,args...)
	return val,err
}

/*
execute sql
*/
func (this *SqlHelper) Exec(sql string,args ...interface{})(sql.Result,stackError.StackError){
	if this.debugMod {
		fmt.Println(sql)
		fmt.Println(args)
	}
	var err error
	stmt,err:=this.prepare(sql)
	if err!=nil {
		return nil, stackError.NewFromError(err,this.stckErrorPowerId)
	}
	defer stmt.Close()
	result,err := stmt.Exec(args...)
	if err!=nil {
		return nil, stackError.NewFromError(err,this.stckErrorPowerId)
	}
	return result,nil
}

/*
execute insert sql
*/
func (this *SqlHelper) ExecInsert(sql string, args ...interface{})(int64,stackError.StackError){
	result,err := this.Exec(sql,args...)
	if err!=nil {
		return 0,err
	}
	
	id,err2 := result.LastInsertId()
	if err2 != nil {
		return 0, stackError.NewFromError(err2,this.stckErrorPowerId)
	}
	return id , nil
}
/*
execute update or delete sql
*/
func (this *SqlHelper) ExecUpdateOrDel(sql string, args ...interface{})(int64,stackError.StackError){
	result,err := this.Exec(sql,args...)
	if err!=nil {
		return 0,err
	}
	
	cnt,err2 := result.RowsAffected()
	if err2 != nil {
		return 0, stackError.NewFromError(err2,this.stckErrorPowerId)
	}
	return cnt , nil
}


/*
    close db pool
*/
func (this *SqlHelper) Close() stackError.StackError{
	err := this.Connection.Close()
	return stackError.NewFromError(err,this.stckErrorPowerId)
}

// get auto sql
func(this *SqlHelper) Auto() *AutoSql{
	return NewAutoSql(this)
}