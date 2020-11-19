package db_conn

import (
	"context"
	"fmt"
	"reflect"
)

type DB2 interface {
	SetData(interface{})
	GetData(interface{}) interface{}
	Connect()
	GetAllParam(interface{}) interface{}
}

type Db2 struct {
	ConfigDb map[string]interface{}
	DbConn
}

func (db2 *Db2) intrFconvetSlice(keys interface{}) []string {
	var keysStrConvert []string
	keysLen := reflect.ValueOf(keys)
	for i:=0; i<keysLen.Len(); i++ {
		re := keysLen.Index(i).Elem().String()
		keysStrConvert = append(keysStrConvert,re)
	}
	return keysStrConvert
}

func (db2 *Db2) GetAllParam(data interface{})  interface{}{
	var i interface{}
	return i
}

func (db2 *Db2)SetData(data interface{}) {
	conn := db2.ConnectOpen
	var ctx = context.Background()
	err := conn.LPush(ctx,"misha",data).Err()
	if err != nil{
		fmt.Println(err)
	}

}

func (db2 *Db2) GetData(data interface{}) interface{} {
	conn := db2.ConnectOpen
	var ctx = context.Background()
	keys,err := conn.LRange(ctx,data.(string),0,-1).Result()
	if err != nil{fmt.Println(err)}
	return keys
}

func (db2 *Db2) Connect() {
	db2.DbConn.DbHost = db2.ConfigDb["DbHost"].(string)
	db2.DbConn.DbPort = db2.ConfigDb["DbPort"].(string)
	db2.DbConn.DbNum  = 2
	db2.DbConn.DbPassword = db2.ConfigDb["DbPassword"].(string)
	db2.DbConn.PoolSiz = db2.ConfigDb["PoolSiz"].(int)
	db2.DbConn.Connect()
}