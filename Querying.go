package goSqlHelper

import "database/sql"

type Querying struct {
	rows *sql.Rows
	cols []string
}


func (this *Querying) Close(){
	this.rows.Close()
	this.cols=nil
}

func (this *Querying) QueryRow() (*HelperRow,error){

	scanArgs := make([]interface{}, len(this.cols))
	values := make([]interface{}, len(this.cols))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	if this.rows.Next() {

		err := this.rows.Scan(scanArgs...)
		if err!=nil {
			return nil,err
		}
		record := make(HelperRow)
		for i, col := range values {
			if col == nil {
				continue
			}
			switch col.(type) {

			case []byte:
				record[this.cols[i]] = string(col.([]byte))
			default:
				record[this.cols[i]] = col

			}
		}

		return &record,nil
	}

	return nil,nil

}


func NewQuerying(rows *sql.Rows) (*Querying,error){
	var err error
	querying:=new(Querying)
	querying.rows=rows

	querying.cols,err=rows.Columns()
	if(err!=nil){
		return nil,err
	}
	return querying,nil
}