package goSqlHelper

import "errors"

/**
  orm read data
*/
func (this *SqlHelper) QueryOrm(orm IOrm, sql string, args ...interface{})(error) {

	rows,err := this.query(sql,args...)
	if err!=nil {
		return err
	}
	defer rows.Close()
	if !rows.Next() {
		return  NoFoundError
	}
	cols,err:=rows.Columns()
	if err!=nil {
		return err
	}
	points := orm.MapFields(cols)
	err = rows.Scan(points...)
	if err!=nil {
		return err
	}
	return nil
}

/*
execute insert sql
*/
func (this *AutoSql) OrmInsert(orm IOrm)(int64,error){
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
		return 0,errors.New("no found insert data")
	}
	return this.sqlHelper.ExecInsert(sql,vals...)
}

func (this *AutoSql) OrmDelete(orm IOrm)(int64,error){
	sql:="DELETE FROM "+orm.TableName()+" WHERE "
	cols:=orm.MapColumn()
	keys:=orm.PrimaryKeys()
	if len(keys)<0{
		return 0,errors.New("no found insert data")
	}

	vals:=make([]interface{},len(keys))
	for i,key:=range keys{
		if i> 0{
			sql+=" AND "
		}
		sql+=key+"=?"
		vals[i]=cols[key]
	}
	return this.sqlHelper.ExecUpdateOrDel(sql,vals...)
}

func (this *AutoSql) OrmUpdate(orm IOrm)(int64,error){
	sql:="UPDATE "+orm.TableName()+" SET "

	cols:=orm.MapColumn()
	keys:=orm.PrimaryKeys()
	if len(keys)<0{
		return 0,errors.New("no found insert data")
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
	return this.sqlHelper.ExecUpdateOrDel(sql,vals...)
}
