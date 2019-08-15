# GoSqlHelper
GoSqlHelper is a go library to help you to execute Sql on mysql,you can easy to query a map[string]interface{} from this library.

## Usage

download source:
```shell
git clone git@github.com:bobby96333/GoSqlHelper.git
```
or
```go.md

require github.com/bobby96333/goSqlHelper v0.0.3

```



Easy to use query result,HelperRow is a map[string]interface{} struct

## Demo code

open db
````go
func OpenDb() *goSqlHelper.SqlHelper{

	con,err:= goSqlHelper.MysqlOpen("root:123456@tcp(centos:3306)/db1")
	if err!=nil {
		panic(err)
	}
	return con
}
````

query row
```go
func main(){
	helper:= util.OpenDb()
	row,err:=helper.QueryRow("select * from tb_tb1 limit 1")
	if err!=nil {
		panic(err)
	}
	fmt.Printf("%+v\n",row)

	//equal to
	row,err=helper.Auto().Select("*").From("tb_tb1").Limit(1).QueryRow()
	if err!=nil {
		panic(err)
	}
	fmt.Printf("%+v\n",row)
	fmt.Println("done\n")

}

```
output:
```text
map[id:1 val:0WWWWWW]
map[id:1 val:0WWWWWW]
done

```

insert/update/delete data
```go


func main(){

	helper:= util.OpenDb()
	helper.OpenDebug()
	cnt,err:= helper.Exec("update tb_tb1 set val=concat(val,'W')")
	if err!=nil {
		panic(err)
	}
	fmt.Println("updates ",cnt," record")

	record:=goSqlHelper.HelperRow{
		"val":"8899",
	}
	id,err:=helper.Auto().Insert("tb_tb1").SetRow(&record).ExecInsert()
	if err!=nil {
		panic(err)
	}
	fmt.Printf("insert a record id:%d\n",id)



	updateCnt,err := helper.Auto().Update("tb_tb1").SetRow(&record).Where("id=?").ExecUpdateOrDel(5)
	if err!=nil {
		panic(err)
	}
	fmt.Printf("update record:%d\n",updateCnt)


	updateCnt,err = helper.Auto().Delete("tb_tb1").Where("id=?").ExecUpdateOrDel(id)
	if err!=nil {
		panic(err)
	}
	fmt.Printf("delete record:%d\n",updateCnt)

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
