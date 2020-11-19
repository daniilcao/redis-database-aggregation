package db_conn

import (
	"context"
	"fmt"
	"reflect"
	"strings"
)

//------------------------------------
type DB0 interface {
	SetData(interface{})
	GetData(interface{}) interface{}
	Connect()
	GetAllParam(interface{}) interface{}
}

//-------------------------------------

type Db0 struct {
	ConfigDb map[string]interface{}
	DbConn
}

func (db0 *Db0)SetData(interface{}) {

}

func (db0 *Db0) getAllUserWorkers(addr[]string,name string) []string {
	var resData []string
	for _,str := range addr{
		split := strings.Split(str, ".")
		if split[0] == name{
			resData = append(resData, str)
		}
	}

	return resData
}

func (db0 *Db0) GetAllParam( data interface{})  interface{}{
	var i interface{}
	return i
}

func (db0 *Db0)GetData(data interface{}) interface{} {
	re  := db0.getAllName()
	reS := db0.intrFconvetSlice(re)
	//reS = db0.getAllUserWorkers(reS,"347GVAZduaeZxXCfqzZhWfkbzRevyS4wV3")
	return reS
}

func (db0 Db0) intrFconvetSlice(keys interface{}) []string {
	var keysStrConvert []string
	keysLen := reflect.ValueOf(keys)
	for i:=0; i<keysLen.Len(); i++ {
		re := keysLen.Index(i).Elem().String()
		keysStrConvert = append(keysStrConvert,re)
	}
	return keysStrConvert
}

func (db0 *Db0) getAllName() interface{} {
	conn := db0.ConnectOpen
	var ctx = context.Background()
	keys,err := conn.Do(ctx,"KEYS", "*").Result()
	if err != nil{fmt.Println(err)}
	return keys
}

func (db0 *Db0)Connect() {
	db0.DbConn.DbHost = db0.ConfigDb["DbHost"].(string)
	db0.DbConn.DbPort = db0.ConfigDb["DbPort"].(string)
	db0.DbConn.DbNum  = 0
	db0.DbConn.DbPassword = db0.ConfigDb["DbPassword"].(string)
	db0.DbConn.PoolSiz = db0.ConfigDb["PoolSiz"].(int)
	db0.DbConn.Connect()
}