# GoSqlHelper
GoSqlHelper is a go library to help you to execute Sql on mysql,you can easy to query a map[string]interface{} from this library.

## Usage

download source:
```shell
git clone git@github.com:bobby96333/GoSqlHelper.git
```
or
```shel
go get github.com/bobby96333/GoShellHelper
```

## Demo

Query a row data

```go

package main

import (
	"fmt"
	"github.com/bobby96333/goSqlHelper"
)

func main(){

	conn,err :=goSqlHelper.MysqlOpen("user:password@tcp(127.0.0.1:3306)/dbname")
	checkErr(err)
	row,err := conn.QueryRow("select * from table where col1 = ? and  col2 = ?","123","abc")
	errCheck(err)
	if *row==nil {
		fmt.Println("no found row")
	}else{
		fmt.Printf("query result：\n %s \n",row.ToJson())
		fmt.Println("get string:",row.String("col2"))
    
    //query a number
		fmt.Println("get Int:",row.PInt("col1"))
		//or
		if col1,err:=row.Int("col1");err!=nil {
			fmt.Println("query col 1 :",col1)
		}
    
	}
}

func checkErr(err error){
	if err!=nil {
		panic(err)
	}
}

```
Query multi-row data
```go
  rows,err := conn.QueryRows("select * from table where col1 = ? and  col2 = ?","123","abc")
	errCheck(err)
	for _,row :=range *rows {
		fmt.Println("row:",row.ToJson())
	}
```

Query a big data

```go
  querying,err := conn.Querying("select * from table where col1 = ? and  col2 = ?","123","abc")
	errCheck(err)
	for row,err:=querying.QueryRow();row!=nil&&err==nil;row,err=querying.QueryRow() {
		fmt.Println("row:",row.ToJson())
	}
```
update or delete sql
```go
  ret,err := conn.Exec("updatetable set col2 = ? where col1 = ? ","abc","123")
	errCheck(err)
	rowNum,err:= ret.RowsAffected()
	errCheck(err)
	fmt.Println("updated row:",rowNum)
```









