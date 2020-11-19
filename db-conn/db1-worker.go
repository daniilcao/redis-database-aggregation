package db_conn

import (
	"context"
	"fmt"
	"reflect"
	"sync"
)

type DB1 interface {
	SetData(interface{})
	GetData(interface{}) interface{}
	Connect()
	GetAllParam(interface{}) interface{}
}


type Db1 struct {
	ConfigDb map[string]interface{}
	DbConn
}

func (db1 *Db1)SetData(interface{}) {

}

func (db1 *Db1) intrFconvetSlice(keys interface{}) []string {
	var keysStrConvert []string
	keysLen := reflect.ValueOf(keys)
	for i:=0; i< keysLen.Len(); i++ {
		re := keysLen.Index(i).String()
		keysStrConvert = append(keysStrConvert,re)
	}
	return keysStrConvert
}

func (db1 *Db1) GetAllParam(data interface{})  interface{}{
	name := db1.intrFconvetSlice(data)
	fmt.Println(name)

	var i interface{}
	return i
}

func (db1 *Db1)GetData(data interface{}) interface{} {
	name := db1.intrFconvetSlice(data)
	wg := sync.WaitGroup{}
	rwMut := sync.RWMutex{}
	var resData  []string
	for _,v := range name{
		wg.Add(1)
		go func(n string) {
			defer wg.Done()
			conn := db1.ConnectOpen
			var ctx = context.Background()
			keys,err := conn.LRange(ctx,n,0,0).Result()
			if err != nil{
				fmt.Println(err)
			}

			if len(keys) != 0 {
				rwMut.Lock()
				resData = append(resData, keys[0])
				rwMut.Unlock()
			}
		}(v)
	}
	wg.Wait()
	return resData
}

func (db1 *Db1)Connect() {
	db1.DbConn.DbHost = db1.ConfigDb["DbHost"].(string)
	db1.DbConn.DbPort = db1.ConfigDb["DbPort"].(string)
	db1.DbConn.DbNum  = 1
	db1.DbConn.DbPassword = db1.ConfigDb["DbPassword"].(string)
	db1.DbConn.PoolSiz = db1.ConfigDb["PoolSiz"].(int)
	db1.DbConn.Connect()
}