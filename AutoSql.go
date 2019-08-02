package goSqlHelper

import (
	"database/sql"
)

const(
	SQL_SELECT="SELECT"
	SQL_UPDATE="UPDATE"
	SQL_DELETE="DELETE"
	SQL_INSERT="INSERT"
)

func NewAutoSql(helper *SqlHelper) *AutoSql{
	var orm=new(AutoSql)
	orm.joins=make([]string,0)
	orm.sqlHelper=helper
	return orm
}

type AutoSql struct{
	act string
	sqlHelper *SqlHelper
	fieldSql string
	tbname string
	where string
	groupBy string
	orderby string
	having string
	limit int
	joins []string
	set string
	setVals []interface{}
}
func (this *AutoSql)Select(fieldSql string) *AutoSql{
	this.act=SQL_SELECT
	this.fieldSql=fieldSql
	return this
}
func (this *AutoSql)Delete(tbname string) *AutoSql{
	this.act=SQL_DELETE
	this.tbname=tbname
	return this
}
func (this *AutoSql)Set(setSql string) *AutoSql{
	this.set=setSql
	return this
}
func (this *AutoSql)SetRow(row *HelperRow) *AutoSql{
	sql:=""
	vals:=make([]interface{},len(*row))
	i:=-1
	for key,val:=range *row {
		i++
		if i>0{
			sql+=","
		}
		sql+=key+"=?"
		vals[i]=val
	}
	this.set=sql
	this.setVals=vals
	return this
}
func (this *AutoSql)Update(tbname string) *AutoSql{
	this.act=SQL_UPDATE
	this.tbname=tbname
	return this
}
func (this *AutoSql)Insert(tbname string) *AutoSql{
	this.act=SQL_INSERT
	this.tbname=tbname
	return this
}
func (this *AutoSql)From(tbname string) *AutoSql{
	this.tbname=tbname
	return this
}
func (this *AutoSql)Where(where string) *AutoSql{
	this.where=where
	return this
}
func (this *AutoSql)Join(joinSql string) *AutoSql{
	this.joins = append(this.joins,joinSql)
	return this
}
func (this *AutoSql)Groupby(groupBySql string) *AutoSql{
	this.groupBy = groupBySql
	return this
}
func (this *AutoSql)Orderby(OrderbySql string) *AutoSql{
	this.orderby = OrderbySql
	return this
}
func (this *AutoSql)Having(having string) *AutoSql{
	this.having = having
	return this
}
func (this *AutoSql)Limit(limit int) *AutoSql{
	this.limit = limit
	return this
}
func (this *AutoSql) GenerateSql() string{
	switch this.act {
	case SQL_SELECT:return this.generateSelectSql()
	case SQL_INSERT:return this.generateInsertSql()
	case SQL_UPDATE:return this.generateUpdateSql()
	case SQL_DELETE:return this.generateDeleteSql()
	default:return this.generateSelectSql()
	}
	panic("no found act:"+this.act)
}

func (this *AutoSql)  QueryRows(args ...interface{})([]HelperRow, error) {
	sql:=this.GenerateSql()
	return this.sqlHelper.QueryRows(sql,args...)
}

func (this *AutoSql) QueryTable( args ...interface{})(*HelperTable, error) {
	sql:=this.GenerateSql()
	return this.sqlHelper.QueryTable(sql,args...)
}

func (this *AutoSql) Querying(args ...interface{})(*Querying,error){
	sql:=this.GenerateSql()
	return this.sqlHelper.Querying(sql,args...)
}

func (this *AutoSql) QueryRow( args ...interface{})(HelperRow, error) {
	sql:=this.GenerateSql()
	return this.sqlHelper.QueryRow(sql,args...)
}

func (this *AutoSql) QueryScalarInt(args ...interface{})(int, error) {
	sql:=this.GenerateSql()
	return this.sqlHelper.QueryScalarInt(sql,args...)
}

func (this *AutoSql) QueryOrm(orm IOrm, args ...interface{})(error) {
	sql:=this.GenerateSql()
	return this.sqlHelper.QueryOrm(orm,sql,args...)
}

func (this *AutoSql) Exec(args ...interface{})(sql.Result,error){
	sql:=this.GenerateSql()
	if this.setVals!=nil {
		args=append(this.setVals,args...)
	}
	return this.sqlHelper.Exec(sql,args...)
}

/*
execute insert sql
*/
func (this *AutoSql) ExecInsert(args ...interface{})(int64,error){
	sql:=this.GenerateSql()
	if this.setVals!=nil {

		args=append(this.setVals,args...)
	}
	return this.sqlHelper.ExecInsert(sql,args...)
}
/*
execute update or delete sql
*/
func (this *AutoSql) ExecUpdateOrDel(args ...interface{})(int64,error){
	sql:=this.GenerateSql()
	if this.setVals!=nil {
		args=append(this.setVals,args...)
	}
	return this.sqlHelper.ExecUpdateOrDel(sql,args...)
}