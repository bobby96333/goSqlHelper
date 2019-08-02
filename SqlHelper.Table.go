package goSqlHelper

/**
  read a record row
*/
func (this *SqlHelper) QueryRow(sql string, args ...interface{})(HelperRow, error) {

	query,err:= this.Querying(sql,args...)
	if err!=nil {
		return nil, err
	}
	defer query.Close()
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
  read a table rows
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
  query muliti rows
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
//
//func (this *SqlHelper) InsertRow(tbname string,row *HelperRow)(int64,error){
//	sql:="INSERT INTO "+tbname+" SET "
//	i:=-1
//	vals:=make([]interface{},len(*row))
//	for key,val:=range *row{
//		i++
//		if i>0{
//			sql+=","
//		}
//		sql+=key+"=?"
//		vals[i]=val
//	}
//	return this.ExecInsert(sql,vals...)
//}
//
//
//func (this *SqlHelper) UpdateRow(tbname string,setRow *HelperRow,whereRow *HelperRow)(int64,error){
//	sql:="UPDATE "+tbname+" SET "
//	i:=-1
//	vals:=make([]interface{},len(*setRow))
//	for key,val:=range *setRow{
//		i++
//		if i>0{
//			sql+=","
//		}
//		sql+=key+"=?"
//		vals[i]=val
//	}
//	sql+=" WHERE "
//	j:=-1
//	for key,val:=range *setRow{
//		j++
//		if j>0{
//			sql+=" AND "
//		}
//		sql+=key+"=?"
//		vals[i+j]=val
//	}
//	return this.ExecInsert(sql,vals...)
//}
//
//func (this *SqlHelper) UpdateRowSql(tbname string,setRow *HelperRow,whereSql string,whereArgs ...interface{})(int64,error){
//	sql:="UPDATE "+tbname+" SET "
//	i:=-1
//	vals:=make([]interface{},len(*setRow))
//	for key,val:=range *setRow{
//		i++
//		if i>0{
//			sql+=","
//		}
//		sql+=key+"=?"
//		vals[i]=val
//	}
//	if whereSql!=""{
//		sql+=" WHERE "+whereSql
//	}
//	vals=append(vals,whereArgs)
//	return this.ExecInsert(sql,vals...)
//}

