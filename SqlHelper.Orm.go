package goSqlHelper

import (
	"github.com/bobby96333/commonLib/stackError"
)

/**
  orm read data
*/
func (this *SqlHelper) QueryOrm(orm IEntity, sql string, args ...interface{})(stackError.StackError) {

	var err error
	rows,err := this.query(sql,args...)
	if err!=nil {
		return stackError.NewFromError( err,this.stckErrorPowerId)
	}
	defer rows.Close()
	if !rows.Next() {
		return  NoFoundError
	}
	cols,err:=rows.Columns()
	if err!=nil {
		return stackError.NewFromError(err,this.stckErrorPowerId)
	}
	points := orm.MapFields(cols)
	err = rows.Scan(points...)
	if err!=nil {
		return stackError.NewFromError(err,this.stckErrorPowerId)
	}
	return nil
}

/*
execute insert sql
*/
func (this *SqlHelper) OrmInsert(orm IEntity)(int64,stackError.StackError){
	sql:="INSERT INTO "+orm.TableName()+" SET "
	cols:=orm.MapColumn()
	i:=-1
	vals:=make([]interface{},len(cols))
	for key,val:=range cols{
		i++
		if i> 0{
			sql+=","
		}
		sql+=key+"=?"
		vals[i]=val
	}
	if i<0{
		return 0,stackError.New("no found insert data",this.stckErrorPowerId)
	}
	return this.ExecInsert(sql,vals...)
}

func (this *SqlHelper) OrmDelete(orm IEntity)(int64,error){
	sql:="DELETE FROM "+orm.TableName()+" WHERE "
	cols:=orm.MapColumn()
	keys:=orm.PrimaryKeys()
	if len(keys)<0{
		return 0,stackError.New("no found insert data",this.stckErrorPowerId)
	}

	vals:=make([]interface{},len(keys))
	for i,key:=range keys{
		if i> 0{
			sql+=" AND "
		}
		sql+=key+"=?"
		vals[i]=cols[key]
	}
	return this.ExecUpdateOrDel(sql,vals...)
}

func (this *SqlHelper) OrmUpdate(orm IEntity)(int64,stackError.StackError){
	sql:="UPDATE "+orm.TableName()+" SET "

	cols:=orm.MapColumn()
	keys:=orm.PrimaryKeys()
	if len(keys)<0{
		return 0,stackError.New("no found insert data",this.stckErrorPowerId)
	}
	vals:=make([]interface{},0,len(cols))

	i:=-1
	for key,val:=range cols{
		has:=false
		for _,primaryKey:=range keys{
			if primaryKey == key{
				has=true
				break
			}
		}
		if has{
			continue //priimarykey
		}
		i++
		if i> 0{
			sql+=" ,"
		}
		sql+=key+"=?"
		vals=append(vals,val)
	}
	sql+=" WHERE "
	for i,key:=range keys{
		if i> 0{
			sql+=" AND "
		}
		sql+=key+"=?"
		vals=append(vals,cols[key])
	}
	return this.ExecUpdateOrDel(sql,vals...)
}
