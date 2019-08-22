package goSqlHelper

import (
	"database/sql"
	"github.com/bobby96333/commonLib/stackError"
)


func(this *SqlHelper) query(sqlStr string,args ...interface{})(*sql.Rows,stackError.StackError){

	var err error
	var rows *sql.Rows
	if this.tx != nil {
		if this.context==nil {
			rows,err = this.tx.QueryContext(this.context,sqlStr,args)
		}else{
			rows,err = this.tx.Query(sqlStr,args)
		}
	}else if this.context != nil {
		rows,err = this.Connection.QueryContext(this.context,sqlStr,args)
	}else{
		rows,err = this.Connection.Query(sqlStr,args...)
	}
	return rows,stackError.NewFromError(err,this.stckErrorPowerId)
}

func(this *SqlHelper) prepare(sqlStr string) (*sql.Stmt, stackError.StackError){
	var smt *sql.Stmt
	var err error
	if this.tx != nil {
		if this.context==nil {
			smt,err = this.tx.PrepareContext(this.context,sqlStr)
		}else{
			smt,err = this.tx.Prepare(sqlStr)
		}
	}else if this.context != nil {
		smt,err = this.Connection.PrepareContext(this.context,sqlStr)
	}else{
		smt,err = this.Connection.Prepare(sqlStr)
	}
	return smt,stackError.NewFromError(err,this.stckErrorPowerId)
}