package goSqlHelper

import (
	"fmt"
	"testing"
	"time"
)

type Tt struct{

	Str string
	val int
	val2 int
	val3 int
	val1 int
	val21 int
	val131 int
	val12 int
	val13 int
	va1l1 int
	val121 int
	val31 int

}

func m1(val Tt){

}

func m2(val *Tt){

}

func TestTt(t *testing.T){



	tval1:=Tt{
		Str:"xxxxxxxxsdfsdfsfdfxxxxxxxxsdfsdfsfdfxxxxxxxxsdfsdfsfdfxxxxxxxxsdfsdfsfdf",
		val:88,

	}
	t1:= time.Now().UnixNano()

	for i:=0;i<10000;i++{
		m1(tval1)
	}
	t2:= time.Now().UnixNano()

	for i:=0;i<10000;i++{
		m2(&tval1)
	}
	t3:= time.Now().UnixNano()

	fmt.Println("step1:",t2-t1)
	fmt.Println("step2:",t3-t2)
	t.Skipped()

}