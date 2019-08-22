package goSqlHelper

import (
	"database/sql"
	"github.com/bobby96333/commonLib/stackError"
)

type Querying struct {
	rows *sql.Rows
	_cols []string
	stackErrorId int
}


func (this *Querying) Close(){
	this.rows.Close()
	//this._cols=nil
}

func (this *Querying) Columns() ([]string, stackError.StackError){
	if this._cols==nil {
		var err error
		this._cols,err = this.rows.Columns()
		if err!=nil {
			return nil, stackError.NewFromError(err,this.stackErrorId)
		}
	}
	return this._cols,nil
}

func (this Querying) QueryRow() (HelperRow,stackError.StackError){

	cols,err:=this.Columns()
	if err!=nil {
		return nil,err
	}
	scanArgs := make([]interface{}, len(cols))
	values := make([]interface{}, len(cols))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	if this.rows.Next() {

		err := this.rows.Scan(scanArgs...)
		if err!=nil {
			return nil,stackError.NewFromError(err,this.stackErrorId)
		}
		record := make(HelperRow)
		for i, col := range values {
			if col == nil {
				continue
			}
			switch col.(type) {

			case []byte:
				record[cols[i]] = string(col.([]byte))
			default:
				record[cols[i]] = col

			}
		}

		return record,nil
	}
	return nil,NoFoundError
}

func (this Querying) Scan(vals ...interface{}) (stackError.StackError){

	if this.rows.Next() {
		err := this.rows.Scan(vals...)
		if err!=nil {
			return stackError.NewFromError(err,this.stackErrorId)
		}
		return nil
	}
	return NoFoundError
}




func NewQuerying(rows *sql.Rows,stackErrorId int) (*Querying){
	querying:=new(Querying)
	querying.rows=rows
	querying.stackErrorId=stackErrorId
	return querying
}