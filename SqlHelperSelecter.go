package goSqlHelper

import "database/sql"


func(this *SqlHelper) query(sql string,args ...interface{})(*sql.Rows,error){
	if this.tx != nil {
		if this.context==nil {
			return this.tx.QueryContext(this.context,sql,args)
		}else{
			return this.tx.Query(sql,args)
		}
	}else if this.context != nil {
		return this.Connection.QueryContext(this.context,sql,args)
	}else{
		return this.Connection.Query(sql,args...)
	}
}

func(this *SqlHelper) prepare(sql string) (*sql.Stmt, error){
	if this.tx != nil {
		if this.context==nil {
			return this.tx.PrepareContext(this.context,sql)
		}else{
			return this.tx.Prepare(sql)
		}
	}else if this.context != nil {
		return this.Connection.PrepareContext(this.context,sql)
	}else{
		return this.Connection.Prepare(sql)
	}
}